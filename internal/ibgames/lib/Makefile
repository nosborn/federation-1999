ARFLAGS = rsv
CC = gcc
CFLAGS = -O2 -Wall -Wconversion -Wextra -Wshadow
CPPFLAGS = -I../include

.PHONY: all
all: libgoodies.a liblog.a librules.a

libgoodies.a: close.o fileSize.o homeDir.o nonBlocking.o read.o trimString.o
	$(AR) $(ARFLAGS) $@ $^

close.o: close.c goodies.h
fileSize.o: fileSize.c goodies.h
homeDir.o: homeDir.c goodies.h
nonBlocking.o: nonBlocking.c goodies.h
read.o: read.c goodies.h
trimString.o: trimString.c goodies.h

liblog.a: log.o logOptions.o
	$(AR) $(ARFLAGS) $@ $^

log.o: log.c goodies.h log.h
logOptions.o: logOptions.c log.h

librules.a: rules.o
	$(AR) $(ARFLAGS) $@ $^

rules.o: rules.c goodies.h rules.h ../include/ibgames.h

.PHONY: clean
clean:
	$(RM) *.a *.o

.PHONY: test
test:
