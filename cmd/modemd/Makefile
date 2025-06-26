BINARY = ../../bin/$(shell go env GOOS)-$(shell go env GOARCH)/modemd
CC = gcc
CFLAGS = -O2 -Wall -Wmissing-prototypes -Wstrict-prototypes -pthread

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
