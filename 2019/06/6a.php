<?php

declare(strict_types=1);
error_reporting(E_ALL);

// Read in the orbits as key => value pairs
$direct_orbits = [];

$fp = fopen('input', 'r');

while ($current_orbit = fgets($fp))
{
  $current_orbit = trim($current_orbit);
  list($orbiting, $object) = explode(')', $current_orbit);
  $direct_orbits[$object] = $orbiting;
}

fclose($fp);

$orbit_count = 0;

foreach ($direct_orbits as $object => $orbiting)
{
  // The direct orbit counts as one
  $orbit_count++;

  // Check all the indirect orbits
  while (isset($direct_orbits[$orbiting]))
  {
    $orbit_count++;
    $orbiting = $direct_orbits[$orbiting];
  }
}

print("$orbit_count\n");
