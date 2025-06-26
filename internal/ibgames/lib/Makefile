CC = gcc
CFLAGS = -fPIC -O2 -Wall -Wconversion -Wextra -Wshadow
CPPFLAGS = -I../include
LDFLAGS = -shared

.PHONY: all
all: libgoodies.so liblog.so librules.so

libgoodies.so: close.o fileSize.o homeDir.o nonBlocking.o read.o trimString.o
	$(CC) $(CFLAGS) $^ -o $@ $(LDFLAGS)

close.o: close.c goodies.h
fileSize.o: fileSize.c goodies.h
homeDir.o: homeDir.c goodies.h
nonBlocking.o: nonBlocking.c goodies.h
read.o: read.c goodies.h
trimString.o: trimString.c goodies.h

liblog.so: log.o logOptions.o
	$(CC) $(CFLAGS) $^ -o $@ $(LDFLAGS)

log.o: log.c goodies.h log.h
logOptions.o: logOptions.c log.h

librules.so: rules.o
	$(CC) $(CFLAGS) $^ -o $@ $(LDFLAGS)

rules.o: rules.c goodies.h rules.h ../include/ibgames.h

.PHONY: clean
clean:
	$(RM) *.o *.so

.PHONY: test
test:
