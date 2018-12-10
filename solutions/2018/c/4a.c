#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>

#include "glib_indirect.h"

enum { EVENT_BEGINS_SHIFT, EVENT_FALLS_ASLEEP, EVENT_WAKES_UP };

static const size_t MAXIMUM_LINE_LENGTH = 50;
static const uint8_t MINUTES_IN_HOUR = 60;

struct guard_event {
  char str[MAXIMUM_LINE_LENGTH];
  uint16_t guard_id;
  uint8_t state;
  uint8_t minute;
};

struct guard {
  uint16_t id;
  uint32_t total_minutes_asleep;
  GSList *events;
};

#if 0
static gint compare_guards(gconstpointer a, gconstpointer b)
{
  const struct guard *a_comp = (const struct guard *) a;
  const struct guard *b_comp = (const struct guard *) b;

  if (a_comp->id == b_comp->id)
  {
    return 0;
  }
  else
  {
    return 1;
  }
}
#endif

static gint compare_events(gconstpointer a, gconstpointer b)
{
  const struct guard_event *a_comp = (const struct guard_event *) a;
  const struct guard_event *b_comp = (const struct guard_event *) b;

  return strcmp(a_comp->str, b_comp->str);
}

int main(void)
{
  char *current_line = calloc(MAXIMUM_LINE_LENGTH, sizeof(char));
  GSList *event_list = NULL;
  struct guard_event *current_event = NULL;
  struct guard_event *new_event = NULL;

  while ((current_line = fgets(current_line, MAXIMUM_LINE_LENGTH, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';
    new_event = calloc(1, sizeof(struct guard_event));
    strcpy(new_event->str, current_line);
    event_list = g_slist_insert_sorted(event_list, new_event, compare_events);
  }

  // Run through the sorted events and set the other fields, plus consolidate
  // into guard list
  uint16_t current_guard_id = 0;
  const char generic_event_str_format[] = "[%*4d-%*2d-%*2d %*2d:%2" SCNu8 "] %5s";
  const char shift_start_event_str_format[] = "[%*4d-%*2d-%*2d %*2d:%*2d] %*5s #%4" SCNu16;
  const uint8_t STATE_STR_LENGTH = 6; // 5 chars + NUL terminator

  GSList *guard_list = NULL;
  uint16_t guard_count = 0;
  bool guard_found = false;
  struct guard *current_guard = NULL;
  struct guard *new_guard = NULL;

  for (GSList *event_list_item = event_list; event_list_item != NULL; event_list_item = event_list_item->next)
  {
    char state_str[STATE_STR_LENGTH];
    current_event = (struct guard_event *) event_list_item->data;

    printf("%s\n", current_event->str);

    sscanf(current_event->str, generic_event_str_format, &current_event->minute, state_str);

    if (strcmp(state_str, "Guard") == 0)
    {
      current_event->state = EVENT_BEGINS_SHIFT;
      sscanf(current_event->str, shift_start_event_str_format, &current_guard_id);
    }
    else if (strcmp(state_str, "wakes") == 0)
    {
      current_event->state = EVENT_WAKES_UP;
    }
    else if (strcmp(state_str, "falls") == 0)
    {
      current_event->state = EVENT_FALLS_ASLEEP;
    }
    else
    {
      fprintf(stderr, "Invalid event state: %s\n", state_str);
      return EXIT_FAILURE;
    }

    current_event->guard_id = current_guard_id;
    guard_found = false;

    // Assign event to a guard if one exists, otherwise create a new guard
    for (GSList *guard_list_item = guard_list; guard_list_item != NULL && !guard_found; guard_list_item = guard_list_item->next)
    {
      current_guard = (struct guard *) guard_list_item->data;

      if (current_guard->id == current_guard_id)
      {
        current_guard->events = g_slist_append(current_guard->events, current_event);
        guard_found = true;
      }
    }

    if (!guard_found)
    {
      new_guard = calloc(1, sizeof(struct guard *));
      new_guard->id = current_guard_id;
      new_guard->total_minutes_asleep = 0;
      new_guard->events = NULL;
      new_guard->events = g_slist_append(new_guard->events, current_event);

      guard_list = g_slist_append(guard_list, new_guard);
      guard_count++;
    }
  }

  printf("Guard count: %" PRIu16 "\n", guard_count);

  // Run through all the guards and calculate their total minutes asleep, keeping
  // track of which one spends the most time asleep
  struct guard *guard_most_time_asleep = NULL;
  uint8_t current_asleep_minute = 0;
  uint8_t current_wake_minute = 0;

  for (GSList *guard_list_item = guard_list; guard_list_item != NULL; guard_list_item = guard_list_item->next)
  {
    current_guard = (struct guard *) guard_list_item->data;

    if (guard_most_time_asleep == NULL)
    {
      guard_most_time_asleep = current_guard;
    }

    for (GSList *event_list_item = current_guard->events; event_list_item != NULL; event_list_item = event_list_item->next)
    {
      current_event = (struct guard_event *) event_list_item->data;

      // Assumption: 'falls asleep' is always followed by 'wakes up'
      if (current_event->state == EVENT_FALLS_ASLEEP)
      {
        current_asleep_minute = current_event->minute;
      }
      else if (current_event->state == EVENT_WAKES_UP)
      {
        current_wake_minute = current_event->minute;
        current_guard->total_minutes_asleep += (current_wake_minute - current_asleep_minute);
      }
    }

    // Update the sleepiest guard
    if (current_guard->total_minutes_asleep > guard_most_time_asleep->total_minutes_asleep)
    {
      guard_most_time_asleep = current_guard;
    }
  }

  printf("Guard who spent most time asleep has ID: %" PRIu16 "\n", guard_most_time_asleep->id);

  // Make a record of how often the sleepiest guard is asleep each minute
  uint8_t *asleep_minutes = calloc(MINUTES_IN_HOUR, sizeof(uint8_t));
  for (uint8_t i = 0; i < MINUTES_IN_HOUR; i++)
  {
    asleep_minutes[i] = 0;
  }

  for (GSList *event_list_item = guard_most_time_asleep->events; event_list_item != NULL; event_list_item = event_list_item->next)
  {
    current_event = (struct guard_event *) event_list_item->data;
    printf("Checking sleepiest guard event. ID: %" PRIu16 ", State: %" PRIu8 ", Minute: %" PRIu8 "\n", current_event->guard_id, current_event->state, current_event->minute);

    if (current_event->state == EVENT_FALLS_ASLEEP)
    {
      current_asleep_minute = current_event->minute;
    }
    else if (current_event->state == EVENT_WAKES_UP)
    {
      current_wake_minute = current_event->minute;

      for (uint8_t j = current_asleep_minute; j < current_wake_minute; j++)
      {
        asleep_minutes[j]++;
      }
    }
  }

  // Find which 'minute' this guard was asleep the most
  uint8_t sleepiest_minute = 0;
  uint8_t highest_asleep_count = 0;

  for (uint8_t i = 0; i < MINUTES_IN_HOUR; i++)
  {
    printf("In minute %" PRIu8 ", guard was asleep %" PRIu8 " times\n", i, asleep_minutes[i]);
    if (asleep_minutes[i] > highest_asleep_count)
    {
      sleepiest_minute = i;
      highest_asleep_count = asleep_minutes[i];
    }
  }

  printf("Sleepiest minute: %" PRIu8 "\n", sleepiest_minute);

  uint32_t strategy_one = guard_most_time_asleep->id * sleepiest_minute;

  printf("Strategy one: %" PRIu32 "\n", strategy_one);

  return EXIT_SUCCESS;
}
