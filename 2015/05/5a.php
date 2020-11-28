<?php

declare(strict_types=1);
error_reporting(E_ALL);

$fp = fopen('input', 'r');

$nice_string_count = 0;

$vowels = ['a', 'e', 'i', 'o', 'u'];
$forbidden_strings = ['ab', 'cd', 'pq', 'xy'];

while (($line = fgets($fp)) !== false)
{
    // fgets includes the trailing line break, so remove it
    $line = trim($line);
    $nice_checks_passed = 0;
    $nice_checks_failed = 0;
    $vowels_found = 0;

    // Assumption: Strings are composed of single-byte characters
    for ($s = 0; $s < strlen($line); $s++)
    {
        if (in_array($line[$s], $vowels))
        {
            $vowels_found++;
        }
    }

    if ($vowels_found >= 3)
    {
        $nice_checks_passed++;
    }

    for ($s = 1; $s < strlen($line); $s++)
    {
        if ($line[$s] === $line[$s - 1])
        {
            $nice_checks_passed++;
            break;
        }
    }

    for ($fs = 0; $fs < count($forbidden_strings); $fs++)
    {
        if (strpos($line, $forbidden_strings[$fs]) !== false)
        {
            $nice_checks_failed = 1;
            break;
        }
    }

    if ($nice_checks_passed === 2 && $nice_checks_failed === 0)
    {
        $nice_string_count++;
    }
}

fclose($fp);

print("Number of nice strings: $nice_string_count\n");