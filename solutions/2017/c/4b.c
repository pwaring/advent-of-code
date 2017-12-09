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
  const gchar skip_char = 'X';

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
        // Words must be equal in length in order to be anagrams
        if (strlen(line_words[i]) == strlen(line_words[j]))
        {
          // Take copies of current and next word so we can manipulate them
          gchar *current_word_copy = calloc(strlen(line_words[i]) + 1, sizeof(gchar));
          gchar *next_word_copy = calloc(strlen(line_words[j]) + 1, sizeof(gchar));

          strcpy(current_word_copy, line_words[i]);
          strcpy(next_word_copy, line_words[j]);

          uint32_t characters_matched = 0;

          // Compare every character of current word with next word
          // Eliminate matching characters as we go along by replacing them
          // with skip_char
          for (uint32_t k = 0; k < strlen(current_word_copy); k++)
          {
            bool character_match = false;

            for (uint32_t m = 0; m < strlen(next_word_copy) && !character_match; m++)
            {
              if (current_word_copy[k] == next_word_copy[m])
              {
                next_word_copy[m] = skip_char;
                characters_matched++;
                character_match = true;
              }
            }
          }

          free(current_word_copy);
          free(next_word_copy);

          if (characters_matched == strlen(line_words[i]))
          {
            // Current word is anagram of next word, so this passphrase is invalid
            valid_passphrase = false;
          }
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
