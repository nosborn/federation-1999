CC = gcc
CPPFLAGS = -DDATADIR=\"data\" -DNDEBUG -I../include -I../../ibgames/include
CXXFLAGS = -O2 -Wall -Wconversion -Wextra -Wshadow
LDFLAGS = -shared

.PHONY: all
all: libworkbench.so

libworkbench.so: workbench.o
	$(CC) $(CFLAGS) $^ -o $@ $(LDFLAGS)

workbench.o: workbench.cc workbench.h

.PHONY: clean
clean:
	$(RM) *.o *.so

.PHONY: test
test:
