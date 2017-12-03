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

  uint16_t string_length = strlen(digits_str);

  // Take first character and append it to the end of the string
  digits_str[string_length] = digits_str[0];
  digits_str[string_length + 1] = '\0';

  string_length = strlen(digits_str);

  uint16_t digits_sum = 0;

  for (uint16_t i = 1; i < string_length; i++)
  {
    char current_char = digits_str[i];
    char previous_char = digits_str[i - 1];

    if (previous_char == current_char && current_char >= '0' && current_char <= '9')
    {
      // Hacky solution, assumes 0-9 are sequential in character set
      digits_sum += (current_char - '0');
    }
  }

  printf("Sum of repeated digits: %" PRIu16 "\n", digits_sum);

  return EXIT_SUCCESS;
}
