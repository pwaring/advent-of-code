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

$shortest_manhatten_distance = PHP_INT_MAX;

foreach ($grid as $x => $grid_x)
{
  foreach ($grid as $y => $grid_y)
  {
    if (isset($grid[$x][$y]) && count($grid[$x][$y]) > 1)
    {
      $manhattan = new Manhattan();
      $distance = $manhattan->distance([0, 0], [$x, $y]);

      // Only use the distance if it is non-zero, as we exclude the central core
      // due to all wires crossing there
      if ($distance)
      {
        $shortest_manhatten_distance = min($distance, $shortest_manhatten_distance);
      }
    }
  }
}

if ($shortest_manhatten_distance < PHP_INT_MAX)
{
  print("$shortest_manhatten_distance\n");
}
else
{
  die("No wires cross at any point other than 0,0\n");
}
