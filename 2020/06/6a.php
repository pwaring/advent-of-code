<?php

declare(strict_types=1);
error_reporting(E_ALL);

$customs_data = file_get_contents('input');
$groups = preg_split('/\R\R/', $customs_data);

$counts_sum = 0;

foreach ($groups as $group)
{
    if ($group)
    {
        foreach (range('a', 'z') as $char)
        {
            // If the current character exists anywhere in the string, it counts as one
            // This is sufficient as we don't care how many people answered, just that one or more did
            if (strpos($group, $char) !== false)
            {
                $counts_sum++;
            }
        }
    }
}

print("Sum of counts: $counts_sum\n");