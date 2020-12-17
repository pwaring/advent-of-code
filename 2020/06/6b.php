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
        $people = preg_split('/\R/', trim($group));
        $people_count = count($people);

        foreach (range('a', 'z') as $char)
        {
            $answer_count = 0;

            foreach ($people as $person)
            {
                if (strpos($person, $char) !== false)
                {
                    $answer_count++;
                }
            }

            if ($answer_count === $people_count)
            {
                $counts_sum++;
            }
        }
    }
}

print("Sum of counts: $counts_sum\n");