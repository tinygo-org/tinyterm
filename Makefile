# md5sum binary
MD5SUM = md5sum

# tinygo binary for tests
TINYGO ?= tinygo

.PHONY: all clean fmt fmt-check

all: clean fmt-check smoketest

clean:
	@mkdir -p build
	@rm -f $(TARGET)

FMT_PATHS = ./*.go examples
fmt:
	@gofmt -l -w $(FMT_PATHS)
fmt-check:
	@unformatted=$$(gofmt -l $(FMT_PATHS)); [ -z "$$unformatted" ] && exit 0; echo "Unformatted:"; for fn in $$unformatted; do echo "  $$fn"; done; exit 1

TARGET = build/examples_basic.hex \
		 build/examples_basic_clue.hex \
		 build/examples_colors.hex \
		 build/examples_httpclient.hex \
		 build/examples_pybadge.hex
.PHONY: smoketest $(TARGET)
smoke-test: clean $(TARGET)

build/examples_basic.hex:
	$(TINYGO) build -size short -o $@ -target=pyportal ./examples/basic
	@$(MD5SUM) $@

build/examples_basic_clue.hex:
	$(TINYGO) build -size short -o $@ -target=clue ./examples/basic-clue
	@$(MD5SUM) $@

build/examples_colors.hex:
	$(TINYGO) build -size short -o $@ -target=pyportal ./examples/colors
	@$(MD5SUM) $@

build/examples_httpclient.hex:
	$(TINYGO) build -size short -o $@ -target=pyportal ./examples/httpclient
	@$(MD5SUM) $@

build/examples_pybadge.hex:
	$(TINYGO) build -size short -o $@ -target=pybadge ./examples/pybadge
	@$(MD5SUM) $@
