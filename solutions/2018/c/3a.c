#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>

#include <glib.h>

struct claim {
  uint16_t id;
  uint16_t left_edge;
  uint16_t top_edge;
  uint16_t width;
  uint16_t height;
};

int main(void)
{
  const size_t maximum_line_length = 25;
  char *current_line = calloc(maximum_line_length, sizeof(char));
  char current_char;

  GSList *claim_list = NULL;
  struct claim *current_claim = NULL;

  while ((current_line = fgets(current_line, maximum_line_length, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';
    current_claim = calloc(1, sizeof(struct claim));
    uint8_t c = 1;

    // Extract each component from the line
    // Start off getting the claim ID, which is all the characters from position
    // 1 until we hit a non-numeric value
    while (c < maximum_line_length && current_line[c] >= '0' && current_line[c] <= '9')
    {
      current_char = current_line[c];

      // Shift the current number along by multiplying by ten, then add the
      // new digit. For example, if the first digit is 1 and the second is 3
      // then the calculation is 1 * 10 = 10 + 3 = 13
      current_claim->id *= 10;
      current_claim->id += strtoumax(&current_char, NULL, 10);

      c++;
    }

    // Skip all characters until we get to a digit again
    while (c < maximum_line_length && !(current_line[c] >= '0' && current_line[c] <= '9')) { c++; }

    // Read in all digits as distance from left edge
    while (c < maximum_line_length && current_line[c] >= '0' && current_line[c] <= '9')
    {
      current_char = current_line[c];
      current_claim->left_edge *= 10;
      current_claim->left_edge += strtoumax(&current_char, NULL, 10);
      c++;
    }

    // Skip all characters until we get to a digit again
    while (c < maximum_line_length && !(current_line[c] >= '0' && current_line[c] <= '9')) { c++; }

    // Read in all digits as distance from top edge
    while (c < maximum_line_length && current_line[c] >= '0' && current_line[c] <= '9')
    {
      current_char = current_line[c];
      current_claim->top_edge *= 10;
      current_claim->top_edge += strtoumax(&current_char, NULL, 10);
      c++;
    }

    // Skip all characters until we get to a digit again
    while (c < maximum_line_length && !(current_line[c] >= '0' && current_line[c] <= '9')) { c++; }

    // Read in all digits as width
    while (c < maximum_line_length && current_line[c] >= '0' && current_line[c] <= '9')
    {
      current_char = current_line[c];
      current_claim->width *= 10;
      current_claim->width += strtoumax(&current_char, NULL, 10);
      c++;
    }

    // Skip all characters until we get to a digit again
    while (c < maximum_line_length && !(current_line[c] >= '0' && current_line[c] <= '9')) { c++; }

    // Read in all digits as height
    while (c < maximum_line_length && current_line[c] >= '0' && current_line[c] <= '9')
    {
      current_char = current_line[c];
      current_claim->height *= 10;
      current_claim->height += strtoumax(&current_char, NULL, 10);
      c++;
    }

    printf("#%" PRIu16 " @ %" PRIu16 ",%" PRIu16 ": %" PRIu16 "x%" PRIu16 "\n", current_claim->id, current_claim->left_edge, current_claim->top_edge, current_claim->width, current_claim->height);

    // Prepend elements as this is quicker and the order doesn't matter when
    // we are comparing elements
    claim_list = g_slist_prepend(claim_list, current_claim);
  }



  return EXIT_SUCCESS;
}
