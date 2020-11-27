<?php

declare(strict_types=1);
error_reporting(E_ALL);

$secret_key = trim(file_get_contents('input'));
$number = 0;

while (true)
{
    $str = $secret_key . strval($number);
    $hash = md5($str);

    if (substr($hash, 0, 5) === '00000')
    {
        break;
    }

    $number++;
}

print("Lowest number to mine AdventCoin: $number\n");