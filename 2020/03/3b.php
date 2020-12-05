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

$trees_hit = [];

$paths = [
    [
        'x_offset' => 1,
        'y_offset' => 1,
    ],
    [
        'x_offset' => 3,
        'y_offset' => 1,
    ],
    [
        'x_offset' => 5,
        'y_offset' => 1,
    ],
    [
        'x_offset' => 7,
        'y_offset' => 1,
    ],
    [
        'x_offset' => 1,
        'y_offset' => 2,
    ]
];

for ($p = 0; $p < count($paths); $p++)
{
    $current_path = $paths[$p];
    $trees_hit[$p] = 0;

    $x = 0;
    $y = 0;

    while ($y < ($map_edge_y -1))
    {
        $x += $current_path['x_offset'];

        // Wrap around if we have gone over the edge because the
        // map repeats itself indefinitely to the right
        if ($x > $map_edge_x)
        {
            $x = $x - ($map_edge_x + 1);
        }

        $y += $current_path['y_offset'];

        if ($map[$x][$y] === TREE)
        {
            $trees_hit[$p]++;
        }
    }
}

$trees_hit_product = array_product($trees_hit);

print("Product of trees hit: $trees_hit_product\n");