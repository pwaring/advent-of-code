# Advent of Code

Advent of Code solutions, primarily in C. I am solving these problems for fun
and to (re)learn languages that I don't use commercially, but would like to.

These are not necessarily the most efficient, elegant or robust solutions. For
example, I have not always freed memory allocated on the heap, partly because
these solutions run for a very short period of time - seconds in most cases -
and the kernel will reclaim any memory allocated to a process when the process
terminates anyway.

As the input is 'known' there is minimal error checking, real programs would
have significantly more validation and unit tests.

For C solutions, C99 is assumed as a minimum and solutions from 2018 day 4 onwards
have been compiled with:

```
clang -Wall -Wextra -Werror -Weverything -pedantic -Wno-padded -Wno-gnu-folding-constant -g -o ${EX_NUM} ${EX_NUM}.c $(shell pkg-config --cflags --libs glib-2.0)
```

`-Wno-padded` is used because padding a struct is a matter for the compiler. If
you rely on a particular type of padding, be it on your own head.

`-Wno-gnu-folding-constant` is used because folding constants is a compiler
optimisation and generally not something that should raise a warning, at least
in the limited scope of these solutions.

`-Weverything` blows up on GLib so the `include` statement is directed through
`glib_indirect.h` which turns off `-Weverything` for the third party library.
