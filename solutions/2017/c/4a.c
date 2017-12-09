#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>
#include <stdbool.h>

#include <glib.h>

int main(void)
{
  const size_t maximum_line_length = 4096;
  const gchar *delimiter = " ";

  char *current_line = calloc(maximum_line_length, sizeof(gchar));
  uint32_t valid_passphrase_count = 0;

  while ((current_line = fgets(current_line, maximum_line_length, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';

    // Assume passphrase is valid unless we find a duplicate word
    bool valid_passphrase = true;

    gchar **line_words = g_strsplit(current_line, delimiter, -1);
    uint32_t word_count = 0;

    // Count the words as this makes later array manipulation easier than using pointers
    for (gchar **current_word = line_words; *current_word != NULL && valid_passphrase; current_word++)
    {
      word_count++;
    }

    for (uint32_t i = 0; i < word_count && valid_passphrase; i++)
    {
      for (uint32_t j = i + 1; j < word_count && valid_passphrase; j++)
      {
        if (strcmp(line_words[i], line_words[j]) == 0)
        {
          valid_passphrase = false;
        }
      }
    }

    g_strfreev(line_words);

    if (valid_passphrase)
    {
      valid_passphrase_count++;
    }
  }

  free(current_line);

  printf("Valid passphrases: %" PRIu32 "\n", valid_passphrase_count);

  return EXIT_SUCCESS;
}
