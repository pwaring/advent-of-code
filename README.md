# Advent of Code

Advent of Code solutions


## Day 1

As the string of digits is circular, the easiest way to handle this is to append
the first digit to the end of the string. We then compare each digit to the
previous digit and add it to the sum if both digits are equal.
