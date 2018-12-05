#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>
#include <time.h>
#include <assert.h>

#include "glib_indirect.h"

enum { EVENT_BEGINS_SHIFT, EVENT_FALLS_ASLEEP, EVENT_WAKES_UP };

struct guard_event {
  struct tm *timestamp;
  uint16_t guard_id;
  uint8_t state;
};

struct guard {
  uint16_t id;
  uint32_t total_minutes_asleep;
  GSList *events;
};

static gint compare_events(gconstpointer a, gconstpointer b)
{
  const struct guard_event *a_event = (const struct guard_event *) a;
  const struct guard_event *b_event = (const struct guard_event *) b;

  time_t a_time = mktime(a_event->timestamp);
  time_t b_time = mktime(b_event->timestamp);

  double diff_time = difftime(a_time, b_time);

  if (diff_time < 0)
  {
    return -1;
  }
  else if (diff_time > 0)
  {
    return 1;
  }
  else
  {
    return 0;
  }
}

static gint compare_guards(gconstpointer a, gconstpointer b)
{
  const struct guard *a_guard = (const struct guard *) a;
  const struct guard *b_guard = (const struct guard *) b;

  if (a_guard->id == b_guard->id)
  {
    return 0;
  }
  else
  {
    return 1;
  }
}

int main(void)
{
  const size_t maximum_line_length = 50;
  char *current_line = calloc(maximum_line_length, sizeof(char));
  GSList *event_list = NULL;
  struct guard_event *current_event = NULL;
  struct tm *event_time = NULL;

  while ((current_line = fgets(current_line, maximum_line_length, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';
    current_event = calloc(1, sizeof(struct guard_event));

    const size_t digit_buffer_size = 3; // 2 digits plus one for the NUL character
    char digit_buffer[digit_buffer_size];

    const size_t state_buffer_size = 6; // 5 chars plus one for NUL
    char state_buffer[state_buffer_size];

    int output_length;

    event_time = calloc(1, sizeof(struct tm));

    // Year is always the same, so just set it to 2016 (last leap year). Using
    // a pre-1970 year will mean that time comparison functions will not work.
    event_time->tm_year = 116; // Number of years since 1900

    // Read in month
    output_length = snprintf(digit_buffer, digit_buffer_size, "%s", &current_line[6]);
    //printf("Output length: %d\n", output_length);
    //printf("Month: %s\n", buffer);
    event_time->tm_mon = (int) strtoimax(digit_buffer, NULL, 10);
    // Subtract 1 from month as time.h assumes zero-indexed months
    event_time->tm_mon--;

    // Read in day
    output_length = snprintf(digit_buffer, digit_buffer_size, "%s", &current_line[9]);
    event_time->tm_mday = (int) strtoimax(digit_buffer, NULL, 10);

    // Read in hour
    output_length = snprintf(digit_buffer, digit_buffer_size, "%s", &current_line[12]);
    event_time->tm_hour = (int) strtoimax(digit_buffer, NULL, 10);

    // Read in minute
    output_length = snprintf(digit_buffer, digit_buffer_size, "%s", &current_line[15]);
    event_time->tm_min = (int) strtoimax(digit_buffer, NULL, 10);

    current_event->timestamp = event_time;

    // Read in first word of action - this is sufficient to work out what state
    // is represented by the event. We can 'cheat' by just reading 5 characters
    // since every state can be identified by that.
    output_length = snprintf(state_buffer, state_buffer_size, "%s", &current_line[19]);

    if (strcmp(state_buffer, "Guard"))
    {
      current_event->state = EVENT_BEGINS_SHIFT;

      // Guard ID starts at position 26. Read in until we run out of digits
      char current_char;

      for (uint8_t c = 26; c < maximum_line_length && current_line[c] >= '0' && current_line[c] <= '9'; c++)
      {
        current_char = current_line[c];
        current_event->guard_id *= 10;
        current_event->guard_id += strtoumax(&current_char, NULL, 10);
      }
    }
    else if (strcmp(state_buffer, "falls"))
    {
      current_event->state = EVENT_FALLS_ASLEEP;
    }
    else if (strcmp(state_buffer, "wakes"))
    {
      current_event->state = EVENT_WAKES_UP;
    }

    // Prepend elements as this is quicker and we are going to sort the list anyway
    event_list = g_slist_prepend(event_list, current_event);
  }

  event_list = g_slist_sort(event_list, compare_events);

  // Build the events into a series of guards
  GSList *guard_list = NULL;
  GSList *guard_list_item = NULL;
  struct guard *current_guard = calloc(1, sizeof(struct guard));
  struct guard *new_guard = NULL;

  for (GSList *list_item = event_list; list_item != NULL; list_item = list_item->next)
  {
    current_event = (struct guard_event *) list_item->data;

    if (current_event->state == EVENT_BEGINS_SHIFT)
    {
      // New shift started, so update current_guard ID
      current_guard->id = current_event->guard_id;
    }

    // We have the guard ID, so check if it already exists in the list
    guard_list_item = g_slist_find_custom(guard_list, current_guard, compare_guards);

    if (guard_list_item != NULL)
    {
      // Guard exists, so append event to its list
      current_guard = (struct guard *) guard_list_item->data;
      current_guard->events = g_slist_append(current_guard->events, current_event);
    }
    else
    {
      // Guard does not exist, so create it and add it to the list
      new_guard = calloc(1, sizeof(struct guard));
      new_guard->id = current_guard->id;
      new_guard->events = g_slist_append(new_guard->events, current_event);

      // Add guard to list
      guard_list = g_slist_append(guard_list, new_guard);
    }
  }

  // Now we have all the guards, we can iterate over their events and calculate the time spent asleep
  time_t time_fall_asleep = 0;
  time_t time_wake_up = 0;
  double time_spent_asleep;

  for (GSList *guard_list_item = guard_list; guard_list_item != NULL; guard_list_item = guard_list_item->next)
  {
    current_guard = (struct guard *) guard_list_item->data;

    for (GSList *event_list_item = current_guard->events; event_list_item != NULL; event_list_item = event_list_item->next)
    {
      struct guard_event *current_event = (struct guard_event *) event_list_item->data;

      if (current_event->state == EVENT_WAKES_UP)
      {
        time_fall_asleep = mktime(current_event->timestamp);
      }
      else if (current_event->state == EVENT_FALLS_ASLEEP)
      {
        time_wake_up = mktime(current_event->timestamp);
        time_spent_asleep = difftime(time_fall_asleep, time_wake_up);

        // Increase the total time spent asleep for this guard
        current_guard = (struct guard *) guard_list_item->data;
        current_guard->total_minutes_asleep += (time_spent_asleep / 60);
      }
    }
  }

  // Find the 'sleepiest' guard
  struct guard *sleepiest_guard = NULL;

  for (GSList *guard_list_item = guard_list; guard_list_item != NULL; guard_list_item = guard_list_item->next)
  {
    current_guard = (struct guard *) guard_list_item->data;

    if (sleepiest_guard == NULL)
    {
      // We haven't found a guard yet, so use the current one
      sleepiest_guard = current_guard;
    }
    else if (current_guard->total_minutes_asleep > sleepiest_guard->total_minutes_asleep)
    {
      // Current guard is sleepier
      sleepiest_guard = current_guard;
    }
  }

  // Find out which minute the sleepiest guard was asleep the most
  uint8_t *minutes_asleep = calloc(60, sizeof(uint8_t));
  int falls_asleep_minute = -1;
  int wakes_up_minute = -1;

  for (GSList *event_list_item = sleepiest_guard->events; event_list_item != NULL; event_list_item = event_list_item->next)
  {
    current_event = (struct guard_event *) event_list_item->data;

    if (current_event->state == EVENT_FALLS_ASLEEP)
    {
      falls_asleep_minute = current_event->timestamp->tm_min;
    }
    else if (current_event->state == EVENT_WAKES_UP)
    {
      wakes_up_minute = current_event->timestamp->tm_min;

      for (int event_minute = falls_asleep_minute; event_minute < wakes_up_minute; event_minute++)
      {
        minutes_asleep[event_minute]++;
      }
    }
  }

  int sleepiest_minute = -1;
  uint8_t highest_sleep_frequency = 0;

  for (uint8_t i = 0; i < 60; i++)
  {
    if (minutes_asleep[i] > highest_sleep_frequency)
    {
      highest_sleep_frequency = minutes_asleep[i];
      sleepiest_minute = i;
    }
  }

  int strategy_one = sleepiest_guard->id * sleepiest_minute;

  printf("Strategy 1 sleepiest guard ID: %" PRIu16 "\n", sleepiest_guard->id);
  printf("Strategy 1 sleepiest minute: %" PRIu8 "\n", sleepiest_minute);

  printf("Strategy 1 answer: %d\n", strategy_one);

  return EXIT_SUCCESS;
}
