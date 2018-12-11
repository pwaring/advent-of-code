#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>
#include <ctype.h>

#include "glib_indirect.h"

int main(void)
{
  GString *input = g_string_new(NULL);
  char current_char;
  char next_char;

  while ((current_char = (char) getchar()) != EOF)
  {
    input = g_string_append_c(input, current_char);
  }

  printf("Units before reacting: %" G_GSIZE_FORMAT "\n", input->len);

  bool units_destroyed;

  do
  {
    units_destroyed = false;

    for (uint32_t i = 0; i < (input->len -1) && !units_destroyed; i++)
    {
      current_char = input->str[i];
      next_char = input->str[i+1];

      // Destroy the characters if they have different case and their lowercase
      // values are equal
      if (((islower(current_char) && isupper(next_char)) || (isupper(current_char) && islower(next_char))) &&
        (tolower(current_char) == tolower(next_char)))
      {
        //printf("Erasing %c%c at position %" PRIu32 "\n", current_char, next_char, i);
        input = g_string_erase(input, i, 2);
        units_destroyed = true;
      }
    }
  } while (units_destroyed);

  printf("Units remaining after reacting: %" G_GSIZE_FORMAT "\n", input->len - 1);

  return EXIT_SUCCESS;
}
