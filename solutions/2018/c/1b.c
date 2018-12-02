#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>

#include <glib.h>

#include "functions.h"

int main(void)
{
  int32_t current_frequency = 0;
  int32_t current_frequency_adjustment = 0;
  const size_t maximum_line_length = 20;
  char *current_line = calloc(maximum_line_length, sizeof(char));

  GSList *frequency_adjustments_list = NULL;

  GTree *adjusted_frequencies_list = g_tree_new(compare_int);
  g_tree_insert(adjusted_frequencies_list, GINT_TO_POINTER(current_frequency), GINT_TO_POINTER(1));

  gboolean current_frequency_found = false;

  // First build the list of frequency adjustments
  while ((current_line = fgets(current_line, maximum_line_length, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';
    current_frequency_adjustment = strtoimax(current_line, NULL, 10);
    frequency_adjustments_list = g_slist_append(frequency_adjustments_list, GINT_TO_POINTER(current_frequency_adjustment));
  }

  // Loop over the frequency adjustments list as often as necessary in order to find the first repetition
  for (GSList* current_frequency_adjustment_item = frequency_adjustments_list; !current_frequency_found; current_frequency_adjustment_item = current_frequency_adjustment_item->next)
  {
    if (current_frequency_adjustment_item == NULL)
    {
      // Loop round to the start of the list
      current_frequency_adjustment_item = frequency_adjustments_list;
    }

    current_frequency_adjustment = GPOINTER_TO_INT(current_frequency_adjustment_item->data);
    current_frequency += current_frequency_adjustment;

    current_frequency_found = g_tree_lookup_extended(adjusted_frequencies_list, GINT_TO_POINTER(current_frequency), NULL, NULL);

    if (current_frequency_found)
    {
      printf("First repeated frequency: %" PRIi32 "\n", current_frequency);
    }
    else
    {
      // Frequency has not been seen before, so add it to the list
      g_tree_insert(adjusted_frequencies_list, GINT_TO_POINTER(current_frequency), GINT_TO_POINTER(1));
    }
  }

  free(current_line);
  g_slist_free(frequency_adjustments_list);
  g_tree_destroy(adjusted_frequencies_list);

  if (!current_frequency_found)
  {
    //fprintf(stderr, "No repeated frequencies found\n");
    return EXIT_FAILURE;
  }

  return EXIT_SUCCESS;
}
