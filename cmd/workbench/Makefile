BINARY = ../../bin/$(shell go env GOOS)-$(shell go env GOARCH)/workbench
CC = gcc
CPPFLAGS = -DDATADIR=\"data\" -DNDEBUG -I../../internal/fed/include -I../../internal/ibgames/include
CXX = g++
CXXFLAGS = -O2 -Wall -Wconversion -Wextra -Wold-style-cast -Wshadow
LDFLAGS = -L../../internal/fed/lib -L../../internal/ibgames/lib
LDLIBS = -lworkbench -llog -lgoodies

HEADERS = fedwb.hh messages.hh
OBJECTS = checker.o drvio.o editor.o events.o genesis.o locations.o messages.o objects.o workbench.o

.PHONY: all
all: $(BINARY)

$(BINARY): $(OBJECTS)
	$(CXX) $(CXXFLAGS) $(LDFLAGS) -o $@ $^ $(LDLIBS)

$(OBJECTS): $(HEADERS)

messages.cc: messages.txt
	perl -- ../../tools/txt2cc.pl <$< >$@

messages.hh: messages.txt
	awk -f ../../tools/txt2hh.awk <$< >$@

.PHONY: clean
clean:
	$(RM) $(BINARY) $(OBJECTS)
	$(RM) messages.cc message.hh

.PHONY: test
test:
