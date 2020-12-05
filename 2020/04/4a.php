<?php

declare(strict_types=1);
error_reporting(E_ALL);

$required_field_names = [
    'byr',
    'iyr',
    'eyr',
    'hgt',
    'hcl',
    'ecl',
    'pid'
];

$passports_data = file_get_contents('input');
$passports = preg_split('/\R\R/', $passports_data);

$valid_passport_count = 0;

foreach ($passports as $passport)
{
    $passport_fields = preg_split('/\s+/', $passport);

    if (count($passport_fields) >= 1)
    {
        $passport_field_names = [];

        foreach ($passport_fields as $passport_field)
        {
            if (!empty($passport_field))
            {
                list($name, $data) = explode(':', $passport_field);
                
                if ($name !== 'cid')
                {
                    $passport_field_names[] = $name;
                }
            }
        }

        if (count(array_diff($required_field_names, $passport_field_names)) === 0)
        {
            $valid_passport_count++;
        }
    }
}

print("Number of valid passports: $valid_passport_count\n");