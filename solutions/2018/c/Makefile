CFLAGS=-Wall -Wextra -Werror
CC=clang

%: %.c
	$(CC) $(CFLAGS) -g -o $@ $< $(shell pkg-config --cflags --libs glib-2.0)
