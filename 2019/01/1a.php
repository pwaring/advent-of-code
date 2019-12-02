<?php

declare(strict_types=1);
error_reporting(E_ALL);

$total_fuel = 0;

$fp = fopen('input', 'r');

while ($module_mass = fgets($fp))
{
  $module_mass = intval($module_mass);
  $module_fuel = floor($module_mass / 3) - 2;
  $total_fuel += $module_fuel;
}

fclose($fp);

print("$total_fuel\n");
