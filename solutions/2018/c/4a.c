#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>

#include "glib_indirect.h"

enum { EVENT_BEGINS_SHIFT, EVENT_FALLS_ASLEEP, EVENT_WAKES_UP };

static const size_t MAXIMUM_LINE_LENGTH = 50;

struct guard_event {
  char str[MAXIMUM_LINE_LENGTH];
  uint16_t guard_id;
  uint8_t state;
  uint8_t minute;
};

struct guard {
  uint16_t id;
  uint32_t total_minutes_asleep;
};

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

  while ((current_line = fgets(current_line, MAXIMUM_LINE_LENGTH, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';
    current_event = calloc(1, sizeof(struct guard_event));
    strcpy(current_event->str, current_line);
    event_list = g_slist_insert_sorted(event_list, current_event, compare_events);
  }

  // Run through the sorted events and set the other fields
  uint16_t current_guard_id = 0;
  const char *generic_event_str_format = "[%*4d-%*2d-%*2d %*2d:%2d] %5s";
  const char *shift_start_event_str_format = "[%*4d-%*2d-%*2d %*2d:%2d] %5s #%4d";
  const uint8_t STATE_STR_LENGTH = 6; // 5 chars + NUL terminator

  for (GSList *event_list_item = event_list; event_list_item != NULL; event_list_item = event_list_item->next)
  {
    char state_str[STATE_STR_LENGTH];
    current_event = (struct guard_event *) event_list_item->data;

    sscanf(current_event->str, generic_event_str_format, current_event->minute, state_str);
  }

  return EXIT_SUCCESS;
}
