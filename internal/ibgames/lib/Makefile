ARFLAGS = rsv
CC = gcc
CFLAGS = -O2 -Wall -Wconversion -Wextra -Wshadow
CPPFLAGS = -I ../include

.PHONY: all
all: libgoodies.a librules.a

libgoodies.a: homeDir.o nonBlocking.o
	$(AR) $(ARFLAGS) $@ $^

homeDir.o: homeDir.c goodies.h
nonBlocking.o: nonBlocking.c goodies.h

librules.a: rules.o
	$(AR) $(ARFLAGS) $@ $^

rules.c: ../include/ibgames.h goodies.h rules.h

.PHONY: clean
clean:
	$(RM) *.a *.o

.PHONY: test
test:
