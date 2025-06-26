.PHONY: all
all: lib install

.PHONY: lib
lib:
	$(MAKE) -C $@ all

.PHONY: install
install: lib
	mkdir -p include/ibgames/
	install -p lib/goodies.h include/ibgames/goodies.h
	install -p lib/log.h include/ibgames/log.h
	install -p lib/rules.h include/ibgames/rules.h

.PHONY: clean
clean: clean-lib

.PHONY: clean-lib
clean-lib:
	$(MAKE) -C lib clean
	$(RM) -r include/ibgames/

.PHONY: test
test:
