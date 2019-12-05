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
      $adjacent_digits = true;
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
