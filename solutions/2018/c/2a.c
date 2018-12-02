#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>

#include <glib.h>

#include "functions.h"

int main(void)
{
  uint16_t two_letter_count = 0;
  uint16_t three_letter_count = 0;

  const size_t maximum_line_length = 50;
  char *current_line = calloc(maximum_line_length, sizeof(char));

  while ((current_line = fgets(current_line, maximum_line_length, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';

    GTree *letter_frequencies = g_tree_new(compare_char);
    gpointer old_letter_count = NULL;
    uint8_t new_letter_count = 0;

    // Build the tree of letter frequencies
    for (uint8_t i = 0; current_line[i] != '\0'; i++)
    {
      old_letter_count = g_tree_lookup(letter_frequencies, GCHAR_TO_POINTER(current_line[i]));

      if (old_letter_count == NULL)
      {
        // Letter has not been seen before, so add it to the tree
        new_letter_count = 1;
        g_tree_insert(letter_frequencies, GCHAR_TO_POINTER(current_line[i]), GUINT_TO_POINTER(new_letter_count));
      }
      else
      {
        // Letter has been seen before, so increment count by one
        new_letter_count = GPOINTER_TO_UINT(old_letter_count);
        new_letter_count++;
        g_tree_replace(letter_frequencies, GCHAR_TO_POINTER(current_line[i]), GUINT_TO_POINTER(new_letter_count));
      }
    }

    // Traverse the tree to find 2 and 3 letter frequencies - but only count them once
    bool two_letter_found = false;
    bool three_letter_found = false;

    

    g_tree_destroy(letter_frequencies);
    letter_frequencies = NULL;
  }

  free(current_line);
  current_line = NULL;
}
