CFLAGS=-Werror -Wall -Wextra -Wno-padded -Wno-gnu-folding-constant
CC=clang

%: %.c
	$(CC) $(CFLAGS) -g -o $@ $< $(shell pkg-config --cflags --libs glib-2.0)
