<?php

declare(strict_types=1);
error_reporting(E_ALL);

const TREE = '#';
const OPEN_SPACE = '.';

$map = [[]];
$x = 0;
$y = 0;
$map_edge_x = 0;
$map_edge_y = 0;

$fp = fopen('input', 'r');

// Build the map by converting the file into an array
while (($line = fgets($fp)) !== false)
{
    $line = trim($line);

    for ($c = 0, $x = 0; $c < strlen($line); $c++, $x++)
    {
        $map[$x][$y] = $line[$c];
    }

    $map_edge_x = max($map_edge_x, strlen($line) - 1);
    $y++;
}

$map_edge_y = $y;

fclose($fp);

$trees_hit = 0;
$x = 0;
$y = 0;

while ($y < ($map_edge_y -1))
{
    $x += 3;

    // Wrap around if we have gone over the edge because the
    // map repeats itself indefinitely to the right
    if ($x > $map_edge_x)
    {
        $x = $x - ($map_edge_x + 1);
    }

    $y += 1;

    if ($map[$x][$y] === TREE)
    {
        $trees_hit++;
    }
}

print("Number of trees hit: $trees_hit\n");