<?php

declare(strict_types=1);
error_reporting(E_ALL);

$fp = fopen('input', 'r');

$paper_order = 0;

while (($line = fgets($fp)) !== false)
{
    // fgets includes the trailing line break, so remove it
    $line = trim($line);
    list($length, $width, $height) = explode('x', $line);

    $side_areas = [
        $length * $width,
        $width * $height,
        $height * $length
    ];

    $slack = min($side_areas);

    // Each side is repeated twice
    $paper_order += array_sum($side_areas) * 2;
    $paper_order += $slack;
}

fclose($fp);

print("The elves should order this many square feet of wrapping paper: $paper_order\n");