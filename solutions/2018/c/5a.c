#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>
#include <ctype.h>

#include "glib_indirect.h"

static GString *react_polymer(GString *polymer)
{
  bool units_destroyed = false;
  char current_char;
  char next_char;

  do
  {
    units_destroyed = false;

    for (uint32_t i = 0; i < (polymer->len -1) && !units_destroyed; i++)
    {
      current_char = polymer->str[i];
      next_char = polymer->str[i+1];

      // Destroy the characters if they have different case and their lowercase
      // values are equal
      if (((islower(current_char) && isupper(next_char)) || (isupper(current_char) && islower(next_char))) &&
        (tolower(current_char) == tolower(next_char)))
      {
        polymer = g_string_erase(polymer, i, 2);
        units_destroyed = true;
      }
    }
  } while (units_destroyed);

  return polymer;
}

int main(void)
{
  GString *original_input = g_string_new(NULL);
  char current_char;

  while ((current_char = (char) getchar()) != EOF)
  {
    original_input = g_string_append_c(original_input, current_char);
  }

  printf("Units before reacting: %" G_GSIZE_FORMAT "\n", original_input->len);

  GString *initial_reaction_input = g_string_new(original_input->str);

  initial_reaction_input = react_polymer(initial_reaction_input);

  printf("Units remaining after reacting: %" G_GSIZE_FORMAT "\n", initial_reaction_input->len - 1);

  uint64_t shortest_polymer_length = UINT64_MAX;

  // Assumption: characters a-z are sequential (true in ASCII character set)
  // Assumption: char can be incremented like an int (required by C standard?)
  for (char candidate_char = 'a'; candidate_char <= 'z'; candidate_char++)
  {
    GString *current_polymer = g_string_new(original_input->str);

    // Remove all lowercase and uppercase instances of the candidate character
    uint32_t i = 0;
    while (i < current_polymer->len)
    {
      current_char = current_polymer->str[i];

      if (tolower(current_char) == candidate_char)
      {
        current_polymer = g_string_erase(current_polymer, i, 1);
      }
      else
      {
        // Only increment i if we did not erase a character, otherwise we will
        // move along and skip a character each time we erase, since the string
        // will have changed.
        i++;
      }
    }

    // React the polymer and check its length
    current_polymer = react_polymer(current_polymer);

    if (current_polymer->len < shortest_polymer_length)
    {
      shortest_polymer_length = current_polymer->len;
    }
  }

  // Reduce the length by 1 as we have a NUL terminating character
  shortest_polymer_length--;

  printf("Shortest polymer length after removing a character and reacting: %" PRIu64 "\n", shortest_polymer_length);

  return EXIT_SUCCESS;
}
