<?php

declare(strict_types=1);
error_reporting(E_ALL);

function binary_space_partition(string $pattern): ?int
{
    $lower_bound = 0;
    $upper_bound = (2 ** strlen($pattern)) - 1;

    for ($c = 0; $c < strlen($pattern); $c++)
    {
        $char = $pattern[$c];
        $mid_point = intval(floor(($lower_bound + $upper_bound) / 2));

        if ($char === 'L')
        {
            $upper_bound = $mid_point;
        }
        elseif ($char === 'R')
        {
            $lower_bound = $mid_point + 1;
        }
    }

    return ($lower_bound === $upper_bound) ? $lower_bound : null;
}

const ROW_COUNT = 128;
const COLUMN_COUNT = 8;

$passes_data = file_get_contents('input');
$passes = preg_split('/\R/', $passes_data);
$seats = [];

// Populate seats
for ($r = 0; $r < ROW_COUNT; $r++)
{
    for ($c = 0; $c < COLUMN_COUNT; $c++)
    {
        $seats[$r][$c] = false;
    }
}

// Occupy seats based on boarding passes
foreach ($passes as $pass)
{
    if ($pass)
    {
        $row_str = substr($pass, 0, 7);
        $column_str = substr($pass, 7);

        // Rewrite row data to use left/right
        $row_str = str_replace('F', 'L', $row_str);
        $row_str = str_replace('B', 'R', $row_str);

        $row = binary_space_partition($row_str);
        $column = binary_space_partition($column_str);

        $seats[$row][$column] = true;
    }
}

// Find the one unoccupied seats
for ($r = 0; $r < ROW_COUNT; $r++)
{
    // Only look between columns 1 and COUNT - 1, because the seats
    // either side of our seat must exist and be occupied
    for ($c = 1; $c < COLUMN_COUNT -1; $c++)
    {
        if ($seats[$r][$c] === false && $seats[$r][$c - 1] && $seats[$r][$c + 1])
        {
            $seat_id = ($r * 8) + $c;
            print("My seat ID: $seat_id\n");
            exit;
        }
    }
}
