<?php

declare(strict_types=1);
error_reporting(E_ALL);

$pattern = '/^([0-9]+)-([0-9]+)\s+([a-z]+):\s+([a-z]+)$/';
$valid_password_count = 0;

$fp = fopen('input', 'r');

while (($line = fgets($fp)) !== false)
{
    $line = trim($line);
    $matches = [];
    
    if (preg_match($pattern, $line, $matches) !== false)
    {
        list($full_text, $first_position, $second_position, $char, $password) = $matches;

        // Subtract one from each position because password policy is 1-indexed
        // whereas strings are 0-indexed
        $first_position--;
        $second_position--;

        if (($password[$first_position] === $char && $password[$second_position] !== $char) || ($password[$first_position] !== $char && $password[$second_position] === $char))
        {
            $valid_password_count++;
        }
    }
}

fclose($fp);

print("Valid passwords: $valid_password_count\n");