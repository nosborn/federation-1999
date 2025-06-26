CPPFLAGS = -I../../internal/fed/include -I../../internal/ibgames/include
CXX = g++
CXXFLAGS = -O2 -Wall -Wconversion -Wextra -Wold-style-cast -Wshadow
LDFLAGS = -L../../internal/ibgames/lib -s
LDLIBS = -lgoodies -lrules -Wl,-rpath=/opt/ibgames/lib

BINARY = ../../bin/perivale
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
