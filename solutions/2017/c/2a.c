#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>

int main(void)
{
  const size_t maximum_line_length = 4096;
  const size_t maximum_number_length = 32;
  char *current_line = calloc(maximum_line_length, sizeof(char));
  uint32_t checksum = 0;

  while ((current_line = fgets(current_line, maximum_line_length, stdin)) != NULL)
  {
    uint32_t smallest_number = 0;
    uint32_t largest_number = 0;

    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';

    uint16_t line_length = strlen(current_line);
    uint16_t line_position = 0;

    while (line_position < line_length)
    {
      char current_number_str[maximum_number_length] = "";
      uint8_t current_number_pos = 0;

      // Greedily match all numbers
      while (current_line[line_position] >= '0' && current_line[line_position] <= '9' && line_position < line_length)
      {
        current_number_str[current_number_pos] = current_line[line_position];
        current_number_pos++;
        line_position++;
      }

      // Convert the string we have extracted to a number
      uint32_t current_number = strtoumax(current_number_str, NULL, 10);

      // Set smallest/largest numbers to current number if appropriate
      if (current_number < smallest_number || smallest_number == 0)
      {
        smallest_number = current_number;
      }

      if (current_number > largest_number)
      {
        largest_number = current_number;
      }

      line_position++;
    }

    // Line has been processed, so update the checksum
    checksum += (largest_number - smallest_number);
  }

  printf("Checksum: %" PRIu32 "\n", checksum);

  free(current_line);

  return EXIT_SUCCESS;
}
