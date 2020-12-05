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

for ($current_expense = 0, $match_found = false; $current_expense < $expense_count - 1 && !$match_found; $current_expense++)
{
    for ($next_expense = $current_expense + 1; $next_expense < $expense_count && !$match_found; $next_expense++)
    {
        if ($expenses[$current_expense] + $expenses[$next_expense] === $target_sum)
        {
            $product = $expenses[$current_expense] * $expenses[$next_expense];
            $match_found = true;
        }
    }
}

print("Product of entries that sum to $target_sum: $product\n");