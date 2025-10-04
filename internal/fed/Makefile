.PHONY: all
all: lib

.PHONY: lib
lib:
	$(MAKE) -C $@ all

.PHONY: clean
clean: clean-lib

.PHONY: clean-lib
clean-lib:
	$(MAKE) -C lib clean

.PHONY: test
test:
