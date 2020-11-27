<?php

declare(strict_types=1);
error_reporting(E_ALL);

$secret_key = trim(file_get_contents('input'));
$number = 0;

$leading_zero_count = 5;
$leading_zero_str = str_repeat('0', $leading_zero_count);

while (true)
{
    $str = $secret_key . strval($number);
    $hash = md5($str);

    if (substr($hash, 0, $leading_zero_count) === $leading_zero_str)
    {
        break;
    }

    $number++;
}

print("Lowest number to mine AdventCoin: $number\n");