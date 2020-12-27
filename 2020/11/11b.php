<?php

declare(strict_types=1);
error_reporting(E_ALL);

/**
 * Return all the seats visible from a given coordinate. Will only return seats
 * so the number of results can be 0 to 8.
 * @param array $map
 * @param int $row_index
 * @param int $column_index
 * @return array Array of seats
 */
function get_visible_seats(array $map, int $row_index, int $column_index): array
{
    $visible_seats = [];
    $adjacent_adjustments = [
        [
            // Top left
            'row' => -1,
            'col' => -1,
        ],
        [
            // Top middle
            'row' => -1,
            'col' => +0,
        ],
        [
            // Top right
            'row' => -1,
            'col' => +1,
        ],
        [
            // Left
            'row' => +0,
            'col' => -1,
        ],
        [
            // Right
            'row' => +0,
            'col' => +1,
        ],
        [
            // Bottom left
            'row' => +1,
            'col' => -1,
        ],
        [
            // Bottom middle
            'row' => +1,
            'col' => +0,
        ],
        [
            // Bottom right
            'row' => +1,
            'col' => +1,
        ]
    ];

    foreach ($adjacent_adjustments as $adjustment)
    {
        $seat_found = false;
        $outside_map = false;

        $row_adjustment = 0;
        $column_adjustment = 0;

        // Keep applying the adjustment until we find a seat or move outside the map
        while (!$seat_found && !$outside_map)
        {
            $row_adjustment += $adjustment['row'];
            $column_adjustment += $adjustment['col'];

            if (isset($map[$row_index + $row_adjustment][$column_index + $column_adjustment]))
            {
                $potential_seat = $map[$row_index + $row_adjustment][$column_index + $column_adjustment];

                if ($potential_seat === SEAT_EMPTY || $potential_seat === SEAT_OCCUPIED)
                {
                    $visible_seats[] = $potential_seat;
                    $seat_found = true;
                }
            }
            else
            {
                // If coordinates do not exist, we have moved outside the map
                $outside_map = true;
            }
        }
    }

    return $visible_seats;
}

const SEAT_EMPTY = 'L';
const SEAT_OCCUPIED = '#';
const FLOOR = '.';

$input_file = $argv[1];

$input_data = trim(file_get_contents($input_file));

// First get all the rows as an array, one per line of the file
$map = preg_split('/\R/', $input_data);

// Convert the text description of each row into columns
$row_count = count($map);

for ($row_index = 0; $row_index < $row_count; $row_index++)
{
    $map[$row_index] = str_split($map[$row_index]);
}

$map_change_count = 0;

do
{
    // All changes are made simultaneously and do not affect each other,
    // therefore we need a copy of the map to work on
    $map_copy = $map;

    $cell_change_count = 0;
    $row_count = count($map_copy);
    $column_count = count($map_copy[0]);

    for ($row_index = 0; $row_index < $row_count; $row_index++)
    {
        for ($column_index = 0; $column_index < $column_count; $column_index++)
        {
            // Remember: We use the current map as the reference but update the copy
            $current_cell = $map[$row_index][$column_index];

            if ($current_cell === SEAT_EMPTY || $current_cell === SEAT_OCCUPIED)
            {
                $visible_seats = get_visible_seats($map, $row_index, $column_index);
                $occupied_visible_seat_count = count(array_filter($visible_seats, function($val) {
                    return $val === SEAT_OCCUPIED;
                }));

                if ($current_cell === SEAT_EMPTY && $occupied_visible_seat_count === 0)
                {
                    $map_copy[$row_index][$column_index] = SEAT_OCCUPIED;
                    $cell_change_count++;
                }
                elseif ($current_cell === SEAT_OCCUPIED && $occupied_visible_seat_count >= 5)
                {
                    $map_copy[$row_index][$column_index] = SEAT_EMPTY;
                    $cell_change_count++;
                }
            }
        }
    }

    // If one or more cells have changed, the whole map has changed
    // and the copy becomes the new map
    if ($cell_change_count > 0)
    {
        $map_change_count++;
        $map = $map_copy;
    }
}
while ($cell_change_count > 0);

print("Number of times map changed before stabilising: $map_change_count\n");

$occupied_seat_count = 0;
for ($row_index = 0; $row_index < $row_count; $row_index++)
{
    for ($column_index = 0; $column_index < $column_count; $column_index++)
    {
        if ($map[$row_index][$column_index] === SEAT_OCCUPIED)
        {
            $occupied_seat_count++;
        }
    }
}

print("Number of occupied seats: $occupied_seat_count\n");