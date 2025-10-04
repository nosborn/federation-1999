CC = gcc
CPPFLAGS = -I../../internal/fed/include -I../../internal/ibgames/include
CXX = g++
CXXFLAGS = -O2 -Wall -Wconversion -Wextra -Wold-style-cast -Wshadow
LDFLAGS = -L../../internal/ibgames/lib
LDLIBS = -lgoodies -lrules

BINARY = ../../bin/$(shell go env GOOS)-$(shell go env GOARCH)/perivale
OBJECTS = perivale.o

.PHONY: all
all: $(BINARY)

$(BINARY): $(OBJECTS)
	$(CXX) $(CXXFLAGS) $(LDFLAGS) -o $@ $^ $(LDLIBS)

perivale.o: perivale.cc

.PHONY: clean
clean:
	$(RM) $(BINARY) $(OBJECTS)

.PHONY: test
test:
