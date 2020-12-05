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
        list($full_text, $min_occurs, $max_occurs, $char, $password) = $matches;
        $char_count = substr_count($password, $char);

        if ($char_count >= $min_occurs && $char_count <= $max_occurs)
        {
            $valid_password_count++;
        }
    }
}

fclose($fp);

print("Valid passwords: $valid_password_count\n");