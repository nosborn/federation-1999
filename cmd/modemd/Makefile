CC = gcc
CFLAGS = -O2 -Wall -Wmissing-prototypes -Wstrict-prototypes -pthread
LDFLAGS = -s

BINARY = ../../bin/modemd
OBJECTS = modem.o modemd.o
HEADERS = modem.h

.PHONY: all
all: $(BINARY)

$(BINARY): $(OBJECTS)
	$(CC) $(CFLAGS) $(LDFLAGS) -o $@ $^ $(LDLIBS)

$(OBJECTS): $(HEADERS)

.PHONY: clean
clean:
	$(RM) $(BINARY) $(OBJECTS)

.PHONY: test
test:
