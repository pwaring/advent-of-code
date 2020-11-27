<?php

declare(strict_types=1);
error_reporting(E_ALL);

$floor = 0;

$fp = fopen('input', 'r');

while (($char = fgetc($fp)) !== false)
{
    if ($char === '(')
    {
        $floor++;
    }
    elseif ($char === ')')
    {
        $floor--;
    }
}

fclose($fp);

print("Final floor: $floor\n");