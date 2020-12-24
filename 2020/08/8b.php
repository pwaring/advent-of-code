<?php

declare(strict_types=1);
error_reporting(E_ALL);

/**
 * Execute a program to find the accumulator.
 * @param array $instructions The instructions to execute.
 * @return int|null The value of the accumulator, or null if the program is terminated (e.g. infinite loop detected).
 */
function execute_program($instructions) : ?int
{
    $accumulator = 0;
    $sp = 0;
    $instruction_count = count($instructions);

    // Execute the instructions
    $sp = 0;

    while (true)
    {
        if ($instructions[$sp]['execution_count'] === 1)
        {
            // Infinite loop detected
            return null;
        }
        elseif ($sp > $instruction_count)
        {
            // Stack pointer has gone beyond the end of the program
            return null;
        }
        else
        {
            $instructions[$sp]['execution_count'] = 1;
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

        // If the next instruction is one beyond the end, the program
        // has completed successfully and we can return the accumulator
        if ($sp === $instruction_count)
        {
            return $accumulator;
        }
    }
}

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

$accumulator = null;
$next_nop_jmp = 0;

while ($accumulator === null)
{
    // Always start with a fresh copy of the instructions, since we must only change one instruction
    $instructions_copy = $instructions;

    for ($ins = $next_nop_jmp; $ins < $instruction_count; $ins++)
    {
        if ($instructions_copy[$ins]['operation'] === 'nop' || $instructions_copy[$ins]['operation'] === 'jmp')
        {
            $next_nop_jmp = $ins;
            break;
        }
    }

    if ($instructions_copy[$next_nop_jmp]['operation'] === 'nop')
    {
        $instructions_copy[$next_nop_jmp]['operation'] = 'jmp';
    }
    elseif ($instructions_copy[$next_nop_jmp]['operation'] === 'jmp')
    {
        $instructions_copy[$next_nop_jmp]['operation'] = 'nop';
    }

    $accumulator = execute_program($instructions_copy);

    $next_nop_jmp++;
}

print("Accumulator value when program runs successfully: $accumulator\n");