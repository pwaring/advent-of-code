<?php

declare(strict_types=1);
error_reporting(E_ALL);

$accumulator = 0;
$input = trim(file_get_contents('input'));
$instructions = preg_split('/\R/', $input);

// Parse instructions into structure
$instruction_count = count($instructions);
for ($ins = 0; $ins < $instruction_count; $ins++)
{
    list($operation, $argument) = preg_split('/\s+/', $instructions[$ins]);

    $instructions[$ins] = [
        'operation' => $operation,
        'argument' => intval($argument),
        'execution_count' => 0
    ];
}

// Execute the instructions
$sp = 0;

while (true)
{
    if ($instructions[$sp]['execution_count'] === 1)
    {
        break;
    }
    else
    {
        $instructions[$sp] ['execution_count'] = 1;
    }

    if ($instructions[$sp]['operation'] === 'nop')
    {
        $sp++;
    }
    elseif ($instructions[$sp]['operation'] === 'jmp')
    {
        $sp += $instructions[$sp]['argument'];
    }
    elseif ($instructions[$sp]['operation'] === 'acc')
    {
        $accumulator += $instructions[$sp]['argument'];
        $sp++;
    }
}

print("Final value of accumulator: $accumulator\n");