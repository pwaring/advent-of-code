<?php

declare(strict_types=1);
error_reporting(E_ALL);

$fp = fopen('input', 'r');

$ribbon_order = 0;

while (($line = fgets($fp)) !== false)
{
    // fgets includes the trailing line break, so remove it
    $line = trim($line);
    list($length, $width, $height) = explode('x', $line);

    $side_perimeters = array_map(function($x) { return $x * 2; }, [
        $length + $width,
        $width + $height,
        $height + $length
    ]);

    $perimeter_ribbon = min($side_perimeters);
    $volume_ribbon = $length * $width * $height;

    $ribbon_order += $perimeter_ribbon;
    $ribbon_order += $volume_ribbon;
}

fclose($fp);

print("The elves should order this many feet of ribbon: $ribbon_order\n");