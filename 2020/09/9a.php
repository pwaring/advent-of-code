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

for ($check_index = $preamble_length; $check_index < $number_count; $check_index++)
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
        print("First number which is not a sum of 2 of the premium $preamble_length numbers: {$numbers[$check_index]}\n");
        exit;
    }
}