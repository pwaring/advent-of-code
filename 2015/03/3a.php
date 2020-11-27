<?php

declare(strict_types=1);
error_reporting(E_ALL);

// Drop off a present at the starting point
$x = 0;
$y = 0;

$presents[$x][$y] = 1;
$houses_visited = 1;

$fp = fopen('input', 'r');

while (($move = fgetc($fp)) !== false)
{
    // Move to new location
    if ($move === '^')
    {
        $y++;
    }
    elseif ($move === 'v')
    {
        $y--;
    }
    elseif ($move === '<')
    {
        $x--;
    }
    elseif ($move === '>')
    {
        $x++;
    }
    else
    {
        die("Invalid move: $move\n");
    }

    // Drop off a present at the new location
    if (isset($presents[$x][$y]))
    {
        $presents[$x][$y]++;
    }
    else
    {
        $presents[$x][$y] = 1;
        $houses_visited++;
    }
}

fclose($fp);

print("Houses visited: $houses_visited\n");