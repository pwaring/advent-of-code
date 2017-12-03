#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>
#include <stdbool.h>

int main(void)
{
  const size_t maximum_line_length = 4096;
  const size_t maximum_number_length = 32;
  const uint32_t maximum_numbers_found = 32;

  char *current_line = calloc(maximum_line_length, sizeof(char));
  uint32_t checksum = 0;

  while ((current_line = fgets(current_line, maximum_line_length, stdin)) != NULL)
  {
    uint32_t numbers_found[maximum_numbers_found] = {0};
    uint32_t numbers_found_count = 0;

    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';

    uint32_t line_length = strlen(current_line);
    uint32_t line_position = 0;

    while (line_position < line_length)
    {
      char current_number_str[maximum_number_length] = "";
      uint32_t current_number_pos = 0;

      // Greedily match all numbers
      while (current_line[line_position] >= '0' && current_line[line_position] <= '9' && line_position < line_length)
      {
        current_number_str[current_number_pos] = current_line[line_position];
        current_number_pos++;
        line_position++;
      }

      // Convert the string we have extracted to a number
      uint32_t current_number = strtoumax(current_number_str, NULL, 10);
      numbers_found[numbers_found_count] = current_number;
      numbers_found_count++;

      line_position++;
    }

    // Line has been processed, so update the checksum
    bool divisors_found = false;
    uint32_t divisors_result = 0;

    for (uint32_t i = 0; i < numbers_found_count && !divisors_found; i++)
    {
      for (uint32_t j = i + 1; j < numbers_found_count && !divisors_found; j++)
      {
        uint32_t min_number = numbers_found[i] < numbers_found[j] ? numbers_found[i] : numbers_found[j];
        uint32_t max_number = numbers_found[i] > numbers_found[j] ? numbers_found[i] : numbers_found[j];

        if (max_number % min_number == 0)
        {
          divisors_found = true;
          divisors_result = max_number / min_number;
        }
      }
    }

    if (!divisors_found)
    {
      fprintf(stderr, "No divisors found\n");
      exit(EXIT_FAILURE);
    }

    checksum += divisors_result;
  }

  printf("Checksum: %" PRIu32 "\n", checksum);

  free(current_line);

  return EXIT_SUCCESS;
}
