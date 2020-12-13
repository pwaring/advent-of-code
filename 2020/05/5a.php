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

$seats_data = file_get_contents('input');
$seats = preg_split('/\R/', $seats_data);
$highest_seat_id = 0;

foreach ($seats as $seat)
{
    if ($seat)
    {
        $row_str = substr($seat, 0, 7);
        $column_str = substr($seat, 7);

        // Rewrite row data to use left/right
        $row_str = str_replace('F', 'L', $row_str);
        $row_str = str_replace('B', 'R', $row_str);

        $row = binary_space_partition($row_str);
        $column = binary_space_partition($column_str);

        $seat_id = ($row * 8) + $column;
        $highest_seat_id = max($highest_seat_id, $seat_id);
    }
}

print("Highest seat ID: $highest_seat_id\n");