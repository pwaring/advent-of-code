CFLAGS=-Wall -Wextra -Werror
CC=clang

4a: 4a.c
	$(CC) $(CFLAGS) -o $@ $< $(shell pkg-config --cflags --libs glib-2.0)

%: %.c
	$(CC) $(CFLAGS) -o $@ $<
