<?php

declare(strict_types=1);
error_reporting(E_ALL);

$floor = 0;
$position = 0;

$fp = fopen('input', 'r');

while (($char = fgetc($fp)) !== false)
{
    $position++;

    if ($char === '(')
    {
        $floor++;
    }
    elseif ($char === ')')
    {
        $floor--;
    }

    if ($floor < 0)
    {
        print("Position of basement: $position\n");
        break;
    }
}

fclose($fp);
