ARFLAGS = rsv
CC = gcc
CPPFLAGS = -DDATADIR=\"data\" -DNDEBUG -I../include -I../../ibgames/include
CXXFLAGS = -O2 -Wall -Wconversion -Wextra -Wshadow

.PHONY: all
all: libworkbench.a

libworkbench.a: workbench.o
	$(AR) $(ARFLAGS) $@ $^

workbench.o: workbench.cc workbench.h

.PHONY: clean
clean:
	$(RM) *.a *.o

.PHONY: test
test:
