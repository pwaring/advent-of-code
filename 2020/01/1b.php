<?php

declare(strict_types=1);
error_reporting(E_ALL);

$expenses = [];
$target_sum = 2020;
$product = 0;

$fp = fopen('input', 'r');

while (($line = fgets($fp)) !== false)
{
    $expenses[] = intval(trim($line));
}

fclose($fp);

$expense_count = count($expenses);

for ($first = 0, $match_found = false; $first < $expense_count - 2 && !$match_found; $first++)
{
    for ($second = $first + 1; $second < $expense_count -1 && !$match_found; $second++)
    {
        for ($third = $second + 1; $third < $expense_count && !$match_found; $third++)
        {
            if ($expenses[$first] + $expenses[$second] + $expenses[$third] === $target_sum)
            {
                $product = $expenses[$first] * $expenses[$second] * $expenses[$third];
                $match_found = true;
            }
        }
    }
}

print("Product of entries that sum to $target_sum: $product\n");