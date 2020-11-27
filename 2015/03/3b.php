<?php

declare(strict_types=1);
error_reporting(E_ALL);

$santas = [
    1 => [
        'name' => 'Human-Santa',
        'x' => 0,
        'y' => 0,
    ],
    2 => [
        'name' => 'Robo-Santa',
        'x' => 0,
        'y' => 0
    ]
];

// Drop off presents from each Santa at the starting point
$presents[0][0] = count($santas);
$houses_visited = 1;

// Human-Santa goes first
$current_santa = 1;

$fp = fopen('input', 'r');

while (($move = fgetc($fp)) !== false)
{
    // Move to new location
    if ($move === '^')
    {
        $santas[$current_santa]['y']++;
    }
    elseif ($move === 'v')
    {
        $santas[$current_santa]['y']--;
    }
    elseif ($move === '<')
    {
        $santas[$current_santa]['x']--;
    }
    elseif ($move === '>')
    {
        $santas[$current_santa]['x']++;
    }
    else
    {
        die("Invalid move: $move\n");
    }

    // Drop off a present at the new location
    if (isset($presents[$santas[$current_santa]['x']][$santas[$current_santa]['y']]))
    {
        $presents[$santas[$current_santa]['x']][$santas[$current_santa]['y']]++;
    }
    else
    {
        $presents[$santas[$current_santa]['x']][$santas[$current_santa]['y']] = 1;
        $houses_visited++;
    }

    // Switch to the next Santa
    if ($current_santa < count($santas))
    {
        $current_santa++;
    }
    else
    {
        $current_santa = 1;
    }
}

fclose($fp);

print("Houses visited: $houses_visited\n");