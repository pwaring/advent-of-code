#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>

#include <glib.h>

#include "functions.h"
#include "macros.h"

int main(void)
{
  uint16_t two_letter_count = 0;
  uint16_t three_letter_count = 0;

  const size_t maximum_line_length = 50;
  char *current_line = calloc(maximum_line_length, sizeof(char));

  char *letters = calloc(maximum_line_length, sizeof(char));
  uint8_t *frequencies = calloc(maximum_line_length, sizeof(uint8_t));

  while ((current_line = fgets(current_line, maximum_line_length, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';

    // Clear the letters and frequencies arrays
    for (uint8_t i = 0; i < maximum_line_length; i++)
    {
      letters[i] = '\0';
      frequencies[i] = 0;
    }

    // Iterate over each character in the current line
    for (uint8_t c = 0; current_line[c] != '\0'; c++)
    {
      char current_char = current_line[c];
      bool char_found = false;

      // Check if letter is already in array
      for (uint8_t i = 0; i < maximum_line_length && !char_found; i++)
      {
        // If letter is in array, increment by 1
        if (letters[i] == current_char)
        {
          char_found = true;
          frequencies[i] = frequencies[i] + 1;
        }

        // If letter is the NUL character, we have reached the end of the list,
        // so insert this letter with a frequency of 1
        if (letters[i] == '\0')
        {
          char_found = true;
          letters[i] = current_char;
          frequencies[i] = 1;
        }
      }
    }

    // We now have the frequencies for each character, so check if any have a
    // frequency of 2 or 3. We only want to match once for each frequency, so
    // if there are 2+ letters with a frequency of two that only counts once.
    bool two_found = false;
    bool three_found = false;

    for (uint8_t i = 0; i < maximum_line_length && (!two_found || !three_found); i++)
    {
      if (frequencies[i] == 2 && !two_found)
      {
        two_letter_count++;
        two_found = true;
      }
      else if (frequencies[i] == 3 && !three_found)
      {
        three_letter_count++;
        three_found = true;
      }
    }
  }

  free(frequencies);
  frequencies = NULL;

  free(letters);
  letters = NULL;

  free(current_line);
  current_line = NULL;

  uint32_t checksum = two_letter_count * three_letter_count;

  printf("Checksum: %" PRIu32 "\n", checksum);

  return EXIT_SUCCESS;
}
