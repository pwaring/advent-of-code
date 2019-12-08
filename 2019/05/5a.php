<?php

declare(strict_types=1);
error_reporting(E_ALL);

define('DEBUG', false);

$fp = fopen('input', 'r');
$input = fgets($fp);
fclose($fp);

$intcodes = explode(',', $input);

// Remove any trailing spaces from instructions
$intcodes = array_map('trim', $intcodes);

$halt = false;
$ip = 0;

while (!$halt)
{
  if (DEBUG) { print_r($intcodes); }
  $instruction = strval($intcodes[$ip]);
  if (DEBUG) { print("Instruction: $instruction\n"); }
  $opcode = trim(substr($instruction, -2));
  if (DEBUG) { print("Opcode: $opcode\n"); }

  // Default to position mode on all parameters
  $parameter_modes = [0, 0, 0];

  if (strlen($instruction) === 5)
  {
    $parameter_modes[0] = substr($instruction, 2, 1);
    $parameter_modes[1] = substr($instruction, 1, 1);
    $parameter_modes[2] = substr($instruction, 0, 1);
  }
  elseif (strlen($instruction) === 4)
  {
    $parameter_modes[0] = substr($instruction, 1, 1);
    $parameter_modes[1] = substr($instruction, 0, 1);
  }
  elseif (strlen($instruction) === 3)
  {
    $parameter_modes[0] = substr($instruction, 0, 1);
  }

  if ($opcode == 1)
  {
    $a = $parameter_modes[0] ? $intcodes[$ip + 1] : $intcodes[$intcodes[$ip + 1]];
    $b = $parameter_modes[1] ? $intcodes[$ip + 2] : $intcodes[$intcodes[$ip + 2]];
    $intcodes[$intcodes[$ip + 3]] = $a + $b;

    $ip += 4;
  }
  elseif ($opcode == 2)
  {
    $a = $parameter_modes[0] ? $intcodes[$ip + 1] : $intcodes[$intcodes[$ip + 1]];
    $b = $parameter_modes[1] ? $intcodes[$ip + 2] : $intcodes[$intcodes[$ip + 2]];
    $intcodes[$intcodes[$ip + 3]] = $a * $b;

    $ip += 4;
  }
  elseif ($opcode == 3)
  {
    //$input = readline("Input: ");
    $input = 1;
    $address = $intcodes[$ip + 1];
    $intcodes[$address] = $input;

    $ip += 2;
  }
  elseif ($opcode == 4)
  {
    $output = $parameter_modes[0] ? $intcodes[$ip + 1] : $intcodes[$intcodes[$ip + 1]];
    print("$output\n");

    $ip += 2;
  }
  elseif ($opcode == 99)
  {
    $halt = true;
  }
  else
  {
    die("Invalid opcode: $opcode\n");
  }
}
