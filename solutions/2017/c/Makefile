CFLAGS=-Wall -Wextra -Werror
CC=clang

4a: 4a.c
	$(CC) $(CFLAGS) -o $@ $< $(shell pkg-config --cflags --libs glib-2.0)

4b: 4b.c
	$(CC) $(CFLAGS) -o $@ $< $(shell pkg-config --cflags --libs glib-2.0)

5a: 5a.c
	$(CC) $(CFLAGS) -o $@ $< $(shell pkg-config --cflags --libs glib-2.0)

5b: 5b.c
	$(CC) $(CFLAGS) -o $@ $< $(shell pkg-config --cflags --libs glib-2.0)

%: %.c
	$(CC) $(CFLAGS) -o $@ $<
