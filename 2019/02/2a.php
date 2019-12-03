<?php

declare(strict_types=1);
error_reporting(E_ALL);

$fp = fopen('input', 'r');
$input = fgets($fp);
fclose($fp);

$intcodes = explode(',', $input);

// Override intcodes
$intcodes[1] = 12;
$intcodes[2] = 2;

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
    die("Unknown opcode " . $intcodes[$p] . " at position $p\n");
  }
}

print($intcodes[0] . "\n");
