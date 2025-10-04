CC = gcc
CPPFLAGS = -I../../internal/ibgames/include -Iinclude
CXX = g++
CXXFLAGS = -O2 -Wall -Wconversion -Wextra -Wold-style-cast -Wshadow -std=c++98
LDFLAGS = -L../../internal/ibgames/lib
LDLIBS = -lgoodies -lrules

BINARY = ../../bin/$(shell go env GOOS)-$(shell go env GOARCH)/perivale
OBJECTS = perivale.o

.PHONY: all
all: $(BINARY)

$(BINARY): $(OBJECTS)
	$(CXX) $(CXXFLAGS) $(LDFLAGS) -o $@ $^ $(LDLIBS)

perivale.o: perivale.cc ../../internal/ibgames/include/ibgames/goodies.h \
	../../internal/ibgames/include/ibgames/rules.h \
	../../internal/ibgames/include/ibgames.h include/link.hh

.PHONY: clean
clean:
	$(RM) $(BINARY) *.o

.PHONY: test
test:
