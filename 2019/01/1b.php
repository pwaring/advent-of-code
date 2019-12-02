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

  for ($fuel_mass = $module_fuel, $fuel_required = true; $fuel_required; $fuel_mass = $fuel_fuel)
  {
    $fuel_fuel = floor($fuel_mass / 3) - 2;

    if ($fuel_fuel > 0)
    {
      $total_fuel += $fuel_fuel;
    }
    else
    {
      $fuel_required = false;
    }
  }
}

fclose($fp);

print("$total_fuel\n");
