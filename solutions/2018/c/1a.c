#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>

int main(void)
{
  int32_t current_frequency = 0;
  int32_t frequency_adjustment = 0;
  const size_t maximum_line_length = 20;
  char *current_line = calloc(maximum_line_length, sizeof(char));

  while ((current_line = fgets(current_line, maximum_line_length, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';
    frequency_adjustment = strtoimax(current_line, NULL, 10);
    current_frequency += frequency_adjustment;
  }

  free(current_line);

  printf("Final frequency: %" PRIi32 "\n", current_frequency);

  return EXIT_SUCCESS;
}
