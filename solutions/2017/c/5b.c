#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>
#include <stdbool.h>

#include <glib.h>

int main(void)
{
  // Reserve some elements to prevent frequent re-allocations
  const guint reserved_size = 256;
  const size_t maximum_line_length = 12;

  uint32_t instruction_count = 0;

  GArray *instructions = g_array_sized_new(FALSE, TRUE, sizeof(int32_t), reserved_size);
  char *current_line = calloc(maximum_line_length, sizeof(char));

  // Read all instructions from standard input
  while ((current_line = fgets(current_line, maximum_line_length, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';

    int32_t current_number = strtoimax(current_line, NULL, 10);
    g_array_append_val(instructions, current_number);
    instruction_count++;
  }

  free(current_line);

  // Execute all instructions
  uint32_t executed_instructions = 0;
  uint32_t current_instruction_index = 0;

  while (current_instruction_index < instructions->len)
  {
    // Fetch current instruction
    int32_t *current_instruction = &g_array_index(instructions, int32_t, current_instruction_index);

    // Jump based on current instruction
    current_instruction_index += *current_instruction;

    // Update old instruction (which we have just jumped from)
    // Note we must do this after jumping because we are using a pointer to both
    // fetch and set the instruction value (as g_array_index returns a pointer).
    if (*current_instruction >= 3)
    {
      *current_instruction -= 1;
    }
    else
    {
      *current_instruction += 1;
    }

    executed_instructions++;
  }

  g_array_free(instructions, FALSE);

  printf("Executed instructions before exiting list: %" PRIu16 "\n", executed_instructions);

  return EXIT_SUCCESS;
}
