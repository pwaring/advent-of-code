<?php

declare(strict_types=1);
error_reporting(E_ALL);

$preamble_length = intval($argv[1]);
$input_file = $argv[2];

$input_data = trim(file_get_contents($input_file));
$numbers = preg_split('/\R/', $input_data);

// Convert all numbers to integers - by default they will be strings
$numbers = array_map('intval', $numbers);
$number_count = count($numbers);

$target_number = 0;

// First we need to find the target number
for ($check_index = $preamble_length; $check_index < $number_count && $target_number === 0; $check_index++)
{
    $match_found = false;

    for ($first_index = $check_index - $preamble_length; $first_index < $check_index && !$match_found; $first_index++)
    {
        for ($second_index = $first_index + 1; $second_index < $check_index && !$match_found; $second_index++)
        {
            if ($numbers[$first_index] + $numbers[$second_index] === $numbers[$check_index])
            {
                $match_found = true;
            }
        }
    }

    if (!$match_found)
    {
        $target_number = $numbers[$check_index];
    }
}

print("Target number: $target_number\n");

// Now we have the target number, find a continuous series which sums to it
for ($start_index = 0; $start_index < $number_count; $start_index++)
{
    $current_sum = $numbers[$start_index];

    for ($end_index = $start_index + 1; $end_index < $number_count && $current_sum < $target_number; $end_index++)
    {
        $current_sum += $numbers[$end_index];

        if ($current_sum === $target_number)
        {
            // Get the part of the array from the start to end index
            $slice_length = $end_index - $start_index;
            $slice = array_slice($numbers, $start_index, $slice_length);
            $min_max_sum = min($slice) + max($slice);

            print("Sum of the min and max of series: $min_max_sum\n");
        }
    }
}