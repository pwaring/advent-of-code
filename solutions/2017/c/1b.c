#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>

int main(void)
{
  const size_t maximum_str_length = 4096;
  char *digits_str = calloc(maximum_str_length, sizeof(char));

  digits_str = fgets(digits_str, maximum_str_length, stdin);

  if (digits_str == NULL)
  {
    fprintf(stderr, "Could not read digits from stdin\n");
    exit(EXIT_FAILURE);
  }

  // Remove any trailing line feeds or carriage returns
  digits_str[strcspn(digits_str, "\r\n")] = '\0';

  const uint16_t string_length = strlen(digits_str);

  if (string_length % 2 != 0)
  {
    fprintf(stderr, "String length is not even: %" PRIu16 "\n", string_length);
    exit(EXIT_FAILURE);
  }

  uint16_t digits_sum = 0;
  uint16_t step = string_length / 2;

  for (uint16_t i = 0; i < string_length; i++)
  {
    uint16_t next_char_position = i + step;

    // If we go beyond the end of the string, wrap around
    if (next_char_position >= string_length)
    {
      next_char_position -= string_length;
    }

    char current_char = digits_str[i];
    char next_char = digits_str[next_char_position];

    if (next_char == current_char && current_char >= '0' && current_char <= '9')
    {
      // Hacky solution, assumes 0-9 are sequential in character set
      digits_sum += (current_char - '0');
    }
  }

  printf("Sum of repeated digits: %" PRIu16 "\n", digits_sum);

  free(digits_str);

  return EXIT_SUCCESS;
}
