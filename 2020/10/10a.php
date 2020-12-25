<?php

declare(strict_types=1);
error_reporting(E_ALL);

$input_file = $argv[1];

$input_data = trim(file_get_contents($input_file));
$adapters = preg_split('/\R/', $input_data);
$adapters = array_map('intval', $adapters);

// Add the outlet as an adapter
$adapters[] = 0;

// Add the built-in adapter
$adapters[] = max($adapters) + 3;

// Sort the adapters, because we need the smallest link between them
sort($adapters, SORT_NUMERIC);

$adapter_count = count($adapters);

$chains = [
    0 => 0,
    1 => 0,
    2 => 0,
    3 => 0
];

// From the second (element 1) adapter, find the chain between it and the previous adapter
for ($a = 1; $a < $adapter_count; $a++)
{
    $chains[$adapters[$a] - $adapters[$a - 1]]++;
}

for ($c = 0; $c <= 3; $c++)
{
    print("$c-jolt differences: {$chains[$c]}\n");
}

$diff_product = $chains[1] * $chains[3];

print("1-jolt differences x 3-jolt differences: $diff_product\n");