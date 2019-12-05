<?php

declare(strict_types=1);
error_reporting(E_ALL);

require_once __DIR__ . '/vendor/autoload.php';

use Phpml\Math\Distance\Manhattan;

// Grid starts as a two dimensional array with one element
$grid = [
  [
    []
  ],
];
$wires = [];

$fp = fopen('input', 'r');

while ($current_wire = fgets($fp))
{
  $wires[] = explode(',', $current_wire);
}

fclose($fp);

// print_r($grid[][]);
// print_r($wires);

for ($i = 0; $i < count($wires); $i++)
{
  $lines = $wires[$i];

  $x = 0;
  $y = 0;

  // Draw the wire, starting from the central point
  for ($j = 0; $j < count($lines); $j++)
  {
    $current_line = $lines[$j];
    $direction = substr($current_line, 0, 1);
    $length = intval(substr($current_line, 1));

    for ($k = 1; $k <= $length; $k++)
    {
      if (isset($grid[$x][$y]))
      {
        if (!in_array($i, $grid[$x][$y]))
        {
          $grid[$x][$y][] = $i;
        }
      }
      else
      {
        $grid[$x][$y][] = $i;
      }

      if ($direction === 'U') { $y++; }
      if ($direction === 'D') { $y--; }
      if ($direction === 'L') { $x--; }
      if ($direction === 'R') { $x++; }
    }
  }
}

$shortest_steps = PHP_INT_MAX;

foreach ($grid as $x => $grid_x)
{
  foreach ($grid as $y => $grid_y)
  {
    if (isset($grid[$x][$y]) && count($grid[$x][$y]) > 1 && !($x == 0 && $y == 0))
    {
      // Intersection found, follow each wire from the central point to the
      // intersection to find the number of steps
      $intersect_x = $x;
      $intersect_y = $y;

      print("Intersect found at: ($intersect_x,$intersect_y)\n");

      $steps = 0;

      for ($i = 0; $i < count($wires); $i++)
      {
        $lines = $wires[$i];

        $wire_x = 0;
        $wire_y = 0;
        $intersect_reached = false;

        // Draw the wire, starting from the central point
        for ($j = 0; $j < count($lines) && !$intersect_reached; $j++)
        {
          $current_line = $lines[$j];
          $direction = substr($current_line, 0, 1);
          $length = intval(substr($current_line, 1));

          // Intersection may occur part-way through a length of wire, so need
          // to check each time we move along one step
          for ($k = 1; $k <= $length && !$intersect_reached; $k++)
          {
            if ($direction === 'U') { $wire_y++; }
            if ($direction === 'D') { $wire_y--; }
            if ($direction === 'L') { $wire_x--; }
            if ($direction === 'R') { $wire_x++; }

            $steps++;

            if ($wire_x == $intersect_x && $wire_y == $intersect_y)
            {
              $intersect_reached = true;
            }
          }

          if ($wire_x == $intersect_x && $wire_y == $intersect_y)
          {
            $intersect_reached = true;
          }
        }
      }

      print("Steps: $steps\n");

      $shortest_steps = min($steps, $shortest_steps);
    }
  }
}

if ($shortest_steps < PHP_INT_MAX)
{
  print("$shortest_steps\n");
}
else
{
  die("No wires cross at any point other than 0,0\n");
}
