<?php

declare(strict_types=1);
error_reporting(E_ALL);

$fp = fopen('input', 'r');

$nice_string_count = 0;

while (($line = fgets($fp)) !== false)
{
    // fgets includes the trailing line break, so remove it
    $line = trim($line);
    $nice_checks_passed = 0;

    // Assumption: Strings are composed of single-byte characters
    // A pair of two letters that appears twice, but not overlapping
    for ($s = 0; $s < strlen($line) - 1; $s++)
    {
        $pair = substr($line, $s, 2);

        if (strpos($line, $pair, $s + 2))
        {
            $nice_checks_passed++;
            break;
        }
    }

    // At least one letter repeated, with exactly one letter between them
    for ($s = 0; $s < strlen($line) - 2; $s++)
    {
        if ($line[$s] === $line[$s + 2])
        {
            $nice_checks_passed++;
            break;
        }
    }

    if ($nice_checks_passed === 2)
    {
        $nice_string_count++;
    }
}

fclose($fp);

print("Number of nice strings: $nice_string_count\n");