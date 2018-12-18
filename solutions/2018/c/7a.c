#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>

#include "glib_indirect.h"
#include "macros.h"

#define DEBUG 1

struct step
{
  char letter;
  bool visited;
  bool active_worker;
  int16_t seconds_remaining;
  GSList *parents;
  GSList *children;
};

static const size_t MAXIMUM_LINE_LENGTH = 50;
static const char step_format[] = "Step %c must be finished before step %c can begin.";
static const int8_t STEP_MINIMUM_SECONDS = 60;
static const uint8_t WORKER_COUNT = 5;

static gint compare_steps(gconstpointer a, gconstpointer b)
{
  const struct step *step_a = (const struct step *) a;
  const struct step *step_b = (const struct step *) b;

  if (step_a->letter == step_b->letter)
  {
    return 0;
  }
  else
  {
    return 1;
  }
}

static bool step_parents_visited(const struct step *current_step)
{
  // Short-circuit - if there are no parents then by definition they have been visited
  if (current_step->parents == NULL)
  {
    return true;
  }

  for (GSList *parent_list_item = current_step->parents; parent_list_item != NULL; parent_list_item = parent_list_item->next)
  {
    struct step *parent_step = (struct step *) parent_list_item->data;

    if (!parent_step->visited)
    {
      // Short-circuit - one unvisited parent is enough to stop
      return false;
    }
  }

  // If we get this far, we have found no unvisited parents, so all must have been visited
  return true;
}

static bool step_parents_complete(const struct step *current_step)
{
  // Short-circuit - if there are no parents then by definition they have been completed
  if (current_step->parents == NULL)
  {
    return true;
  }

  for (GSList *parent_list_item = current_step->parents; parent_list_item != NULL; parent_list_item = parent_list_item->next)
  {
    struct step *parent_step = (struct step *) parent_list_item->data;

    if (parent_step->seconds_remaining > 0)
    {
      // Short-circuit - one parent with time remaining is enough to stop
      return false;
    }
  }

  // If we get this far, we have found no parents with time remaining, so all must have been completed
  return true;
}

int main(void)
{
  GSList *step_list = NULL;
  GSList *source_item = NULL;
  GSList *target_item = NULL;

  struct step *source_step = NULL;
  struct step *target_step = NULL;

  char *current_line = calloc(MAXIMUM_LINE_LENGTH, sizeof(char));

  while ((current_line = fgets(current_line, MAXIMUM_LINE_LENGTH, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';

    // Inefficient - we may not use every source/target step so this wastes memory
    source_step = calloc(1, sizeof(struct step));
    target_step = calloc(1, sizeof(struct step));

    // Populate the current source and target steps
    sscanf(current_line, step_format, &source_step->letter, &target_step->letter);

    source_step->visited = false;
    source_step->active_worker = false;
    source_step->seconds_remaining = STEP_MINIMUM_SECONDS + (source_step->letter - 'A') + 1; // Add 1 because 'A' - 'A' is 0
    source_step->parents = NULL;
    source_step->children = NULL;

    target_step->visited = false;
    target_step->active_worker = false;
    target_step->seconds_remaining = STEP_MINIMUM_SECONDS + (target_step->letter - 'A') + 1;
    target_step->parents = NULL;
    target_step->parents = NULL;

    // Find existing list items for source and target
    source_item = g_slist_find_custom(step_list, source_step, compare_steps);
    target_item = g_slist_find_custom(step_list, target_step, compare_steps);

    if (source_item == NULL && target_item == NULL)
    {
      // Neither source nor target exist, so add source and target to list and link to each other
      printf("Source (%c) and target (%c) do not exist, creating both\n", source_step->letter, target_step->letter);

      source_step->children = g_slist_append(source_step->children, target_step);
      target_step->parents = g_slist_append(target_step->parents, source_step);

      step_list = g_slist_append(step_list, source_step);
      step_list = g_slist_append(step_list, target_step);
    }
    else if (source_item != NULL && target_item == NULL)
    {
      // Source exists but target does not, so link target to source
      printf("Source (%c) exists but target (%c) does not exist, adding target to source\n", source_step->letter, target_step->letter);

      source_step = (struct step *) source_item->data;
      source_step->children = g_slist_append(source_step->children, target_step);
      target_step->parents = g_slist_append(target_step->parents, source_step);

      step_list = g_slist_append(step_list, target_step);
    }
    else if (source_item == NULL && target_item != NULL)
    {
      // Source does not exist but target does, so link source to root and link
      // target to source
      printf("Source (%c) does not exist and target (%c) does exist, adding source to root and target to source\n", source_step->letter, target_step->letter);

      target_step = (struct step *) target_item->data;
      source_step->children = g_slist_append(source_step->children, target_step);
      target_step->parents = g_slist_append(target_step->parents, source_step);

      step_list = g_slist_append(step_list, source_step);
    }
    else if (source_item != NULL && target_item != NULL)
    {
      // Source and target exist, so link target to source
      printf("Source (%c) and target (%c) exist, adding target to source\n", source_step->letter, target_step->letter);

      source_step = (struct step *) source_item->data;
      target_step = (struct step *) target_item->data;
      source_step->children = g_slist_append(source_step->children, target_step);
      target_step->parents = g_slist_append(target_step->parents, source_step);
    }
    else
    {
      // This should never happen
      fprintf(stderr, "Failed to add source (%c) and target (%c) nodes\n", source_step->letter, target_step->letter);
      return EXIT_FAILURE;
    }
  }

  #if DEBUG
  for (GSList *step_list_item = step_list; step_list_item != NULL; step_list_item = step_list_item->next)
  {
    struct step *current_step = (struct step *) step_list_item->data;

    printf("Step: %c\n", current_step->letter);
    printf("Children: ");

    for (GSList *child_list_item = current_step->children; child_list_item != NULL; child_list_item = child_list_item->next)
    {
      struct step *child_step = (struct step *) child_list_item->data;
      printf("%c", child_step->letter);
    }

    printf("\n");
  }
  #endif

  guint step_count = g_slist_length(step_list);
  guint steps_visited = 0;

  printf("Part A: ");

  while (steps_visited < step_count)
  {
    struct step *next_step = NULL;

    for (GSList *step_list_item = step_list; step_list_item != NULL; step_list_item = step_list_item->next)
    {
      struct step *current_step = (struct step *) step_list_item->data;

      if (!current_step->visited && step_parents_visited(current_step))
      {
        // See if this step's letter is before the next step
        if (next_step == NULL || current_step->letter < next_step->letter)
        {
          next_step = current_step;
        }
      }
    }

    next_step->visited = true;

    printf("%c", next_step->letter);

    steps_visited++;
  }

  printf("\n");

  printf("Part B: ");

  uint16_t seconds_counter = 0;
  struct step *workers[WORKER_COUNT];
  guint steps_remaining = g_slist_length(step_list);

  // Initialise all workers to have no steps
  for (uint8_t w = 0; w < WORKER_COUNT; w++)
  {
    workers[w] = NULL;
  }

  while (steps_remaining > 0)
  {
    // Check if any workers are idle and we have available steps - if so assign them
    for (uint8_t w = 0; w < WORKER_COUNT; w++)
    {
      // Find the next step which has time remaining, all parents are complete and is not being actively worked on
      for (GSList *step_list_item = step_list; workers[w] == NULL && step_list_item != NULL; step_list_item = step_list_item->next)
      {
        struct step *current_step = (struct step *) step_list_item->data;

        if (current_step->seconds_remaining > 0 && !current_step->active_worker && step_parents_complete(current_step))
        {
          workers[w] = current_step;
          workers[w]->active_worker = true;
        }
      }
    }

    // Time passes
    seconds_counter++;

    // Count down on all worker tasks
    for (uint8_t w = 0; w < WORKER_COUNT; w++)
    {
      if (workers[w] != NULL)
      {
        workers[w]->seconds_remaining--;

        if (workers[w]->seconds_remaining == 0)
        {
          // Step complete, free up worker for another step
          workers[w]->active_worker = false;
          workers[w] = NULL;
          steps_remaining--;
        }
      }
    }
  }

  printf("%" PRIu16 "\n", seconds_counter);

  return EXIT_SUCCESS;
}
