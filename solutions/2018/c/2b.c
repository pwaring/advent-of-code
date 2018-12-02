#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>

#include <glib.h>

void destroy_box_id(gpointer data)
{
  // box_id (data) is created via g_strdup, so must be freed using g_free
  g_free(data);
}

int main(void)
{
  const size_t maximum_line_length = 50;
  char *current_line = calloc(maximum_line_length, sizeof(char));

  GSList *box_id_list = NULL;

  while ((current_line = fgets(current_line, maximum_line_length, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';

    // Prepend elements as this is quicker and we don't actually care about
    // the order in this case.
    box_id_list = g_slist_prepend(box_id_list, g_strdup(current_line));
  }

  char *current_box_id_str = calloc(maximum_line_length, sizeof(char));
  char *next_box_id_str = calloc(maximum_line_length, sizeof(char));

  for (GSList *current_box_id = box_id_list; current_box_id != NULL; current_box_id = current_box_id->next)
  {
    for (GSList *next_box_id = current_box_id->next; next_box_id != NULL; next_box_id = next_box_id->next)
    {
      current_box_id_str = (char *)(current_box_id->data);
      next_box_id_str = (char *)(next_box_id->data);

      uint8_t different_chars_count = 0;
      int8_t different_char_position = -1;

      // Count the number of different characters between the two box IDs
      for (uint16_t c = 0; current_box_id_str[c] != '\0' && next_box_id_str[c] != '\0'; c++)
      {
        char current_box_char = current_box_id_str[c];
        char next_box_char = next_box_id_str[c];

        if (current_box_char != next_box_char)
        {
          different_chars_count++;
          different_char_position = c;
        }
      }

      if (different_chars_count == 1)
      {
        printf("Common letters between two correct box IDs: ");

        // Print the 'string', minus the different char
        for (uint16_t c = 0; current_box_id_str[c] != '\0'; c++)
        {
          if (c != different_char_position)
          {
            printf("%c", current_box_id_str[c]);
          }
        }

        printf("\n");
      }
    }
  }

  // Free memory in reverse order of allocation
  free(next_box_id_str);
  next_box_id_str = NULL;

  free(current_box_id_str);
  current_box_id_str = NULL;

  g_slist_free_full(box_id_list, destroy_box_id);
  box_id_list = NULL;

  free(current_line);
  current_line = NULL;

  return EXIT_SUCCESS;
}
