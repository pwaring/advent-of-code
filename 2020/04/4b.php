<?php

declare(strict_types=1);
error_reporting(E_ALL);

const VALID_HAIR_COLOUR = '/^#([0-9a-f]{6})$/';
const VALID_PASSPORT_ID = '/^([0-9]{9})$/';

$required_field_names = [
    'byr',
    'iyr',
    'eyr',
    'hgt',
    'hcl',
    'ecl',
    'pid'
];

$required_field_count = count($required_field_names);

$valid_eye_colours = [
    'amb',
    'blu',
    'brn',
    'gry',
    'grn',
    'hzl',
    'oth'
];

$passports_data = file_get_contents('input');
$passports = preg_split('/\R\R/', $passports_data);

$valid_passport_count = 0;

foreach ($passports as $passport)
{
    $passport_fields = preg_split('/\s+/', $passport);
    $valid_fields = 0;

    for ($pf = 0; $pf < count($passport_fields); $pf++)
    {
        if (!empty($passport_fields[$pf]))
        {
            list($name, $data) = explode(':', $passport_fields[$pf]);

            if ($name === 'byr')
            {
                if ($data >= 1920 && $data <= 2002)
                {
                    $valid_fields++;
                }
            }
            elseif ($name === 'iyr')
            {
                if ($data >= 2010 && $data <= 2020)
                {
                    $valid_fields++;
                }
            }
            elseif ($name === 'eyr')
            {
                if ($data >= 2020 && $data <= 2030)
                {
                    $valid_fields++;
                }
            }
            elseif ($name === 'hgt')
            {
                $unit = substr($data, -2);
                $height = substr($data, 0, strlen($data) - 2);

                if ($unit === 'cm' && $height >= 150 && $height <= 193)
                {
                    $valid_fields++;
                }
                elseif ($unit === 'in' && $height >= 59 && $height <= 76)
                {
                    $valid_fields++;
                }
            }
            elseif ($name === 'hcl')
            {
                if (preg_match(VALID_HAIR_COLOUR, $data) === 1)
                {
                    $valid_fields++;
                }
            }
            elseif ($name === 'ecl')
            {
                if (in_array($data, $valid_eye_colours))
                {
                    $valid_fields++;
                }
            }
            elseif ($name === 'pid')
            {
                if (preg_match(VALID_PASSPORT_ID, $data) === 1)
                {
                    $valid_fields++;
                }
            }
        }
    }

    if ($valid_fields === $required_field_count)
    {
        $valid_passport_count++;
    }
}

print("Number of valid passports: $valid_passport_count\n");