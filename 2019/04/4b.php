<?php

declare(strict_types=1);
error_reporting(E_ALL);

define('RANGE_START', 168630);
define('RANGE_END', 718098);

$password_matches = 0;

for ($password = RANGE_START; $password <= RANGE_END; $password++)
{
  $password_chars = str_split(strval($password));
  $adjacent_digits = false;
  $digits_increase = true;

  for ($c = 0; $c < (count($password_chars) - 1); $c++)
  {
    if ($password_chars[$c] == $password_chars[$c + 1])
    {
      $previous_digit_match = true;
      $next_digit_match = true;

      if ($c == 0)
      {
        $previous_digit_match = false;
      }
      else
      {
        $previous_digit_match = ($password_chars[$c - 1] == $password_chars[$c]);
      }

      // Check the next digit but one if we are not at the penultimate digit
      if ($c == count($password_chars) - 2)
      {
        $next_digit_match = false;
      }
      else
      {
        $next_digit_match = ($password_chars[$c] == $password_chars[$c + 2]);
      }

      if (!$previous_digit_match && !$next_digit_match)
      {
        $adjacent_digits = true;
      }
    }

    if ($password_chars[$c] > $password_chars[$c + 1])
    {
      $digits_increase = false;
    }
  }

  if ($adjacent_digits && $digits_increase)
  {
    $password_matches++;
  }
}

print("$password_matches\n");
