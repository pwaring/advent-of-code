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
  uint16_t fabric_width;
  uint16_t fabric_height;
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

    current_claim->fabric_width = current_claim->left_edge + current_claim->width;
    current_claim->fabric_height = current_claim->top_edge + current_claim->height;

    // Prepend elements as this is quicker and the order doesn't matter when
    // we are comparing elements
    claim_list = g_slist_prepend(claim_list, current_claim);
  }

  // Find the smallest piece of fabric which can fit every claim. The width will
  // be max(left_edge + width) and the height will be max(top_edge + height).
  uint16_t fabric_width = 0;
  uint16_t fabric_height = 0;

  for (GSList *list_item = claim_list; list_item != NULL; list_item = list_item->next)
  {
    current_claim = (struct claim *) list_item->data;

    if (current_claim->fabric_width > fabric_width)
    {
      fabric_width = current_claim->fabric_width;
    }

    if (current_claim->fabric_height > fabric_height)
    {
      fabric_height = current_claim->fabric_height;
    }
  }

  // Now that we have the fabric width and height, we can build an array which
  // maintains a count of the number of times a square inch is covered.
  uint16_t **fabric = calloc(fabric_width, sizeof(uint16_t *));
  for (uint16_t i = 0; i < fabric_height; i++)
  {
    fabric[i] = calloc(fabric_height, sizeof(uint16_t));
  }

  for (GSList *list_item = claim_list; list_item != NULL; list_item = list_item->next)
  {
    current_claim = (struct claim *) list_item->data;

    for (uint16_t x = current_claim->left_edge; x < current_claim->fabric_width; x++)
    {
      for (uint16_t y = current_claim->top_edge; y < current_claim->fabric_height; y++)
      {
        fabric[x][y]++;
      }
    }
  }

  // Find all the inches (i.e. unique x,y co-ordinates in fabric) with two or
  // more claims
  uint32_t double_claimed_inches = 0;

  for (uint16_t x = 0; x < fabric_width; x++)
  {
    for (uint16_t y = 0; y < fabric_width; y++)
    {
      if (fabric[x][y] >= 2)
      {
        double_claimed_inches++;
      }
    }
  }

  printf("Double claimed inches: %" PRIu32 "\n", double_claimed_inches);

  return EXIT_SUCCESS;
}
