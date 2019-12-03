<?php

declare(strict_types=1);
error_reporting(E_ALL);

define('TARGET_OUTPUT', 19690720);

$fp = fopen('input', 'r');
$input = fgets($fp);
fclose($fp);

$original_intcodes = explode(',', $input);
$intcodes = [];

while (true)
{
  for ($noun = 0; $noun <= 99; $noun++)
  {
    for ($verb = 0; $verb <= 99; $verb++)
    {
      // Reset program
      $intcodes = $original_intcodes;
      $intcodes[1] = $noun;
      $intcodes[2] = $verb;

      for ($p = 0, $halt = false; $p < count($intcodes) && !$halt; $p += 4)
      {
        $opcode = $intcodes[$p];

        if ($opcode == 1)
        {
          $intcodes[$intcodes[$p + 3]] = $intcodes[$intcodes[$p + 1]] + $intcodes[$intcodes[$p + 2]];
        }
        elseif ($opcode == 2)
        {
          $intcodes[$intcodes[$p + 3]] = $intcodes[$intcodes[$p + 1]] * $intcodes[$intcodes[$p + 2]];
        }
        elseif ($opcode == 99)
        {
          $halt = true;
        }
        else
        {
          // Unknown opcode
          print_r($intcodes);
          print("noun: $noun\n");
          print("verb: $verb\n");
          die("Unknown opcode " . $intcodes[$p] . " at position $p\n");
        }
      }

      if ($intcodes[0] == TARGET_OUTPUT)
      {
        $answer = (100 * $noun) + $verb;
        print("$answer\n");
        exit();
      }
    }
  }
}
