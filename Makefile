# md5sum binary
MD5SUM = md5sum

# tinygo binary for tests
TINYGO ?= tinygo

.PHONY: all clean fmt fmt-check test smoke-test

all: clean fmt-check test

test: smoke-test

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
		 build/examples_basic_gopher_badge.hex \
		 build/examples_basic_pybadge.hex \
		 build/examples_basic_wioterminal.hex \
		 build/examples_basic_badger2040.hex \
		 build/examples_colors.hex \
		 build/examples_httpclient.hex
.PHONY: smoketest $(TARGET)
smoke-test: clean $(TARGET)

build/examples_basic.hex:
	$(TINYGO) build -size short -o $@ -target=pyportal ./examples/basic
	@$(MD5SUM) $@

build/examples_basic_clue.hex:
	$(TINYGO) build -size short -o $@ -target=clue ./examples/basic
	@$(MD5SUM) $@

build/examples_basic_gopher_badge.hex:
	$(TINYGO) build -size short -o $@ -target=gopher-badge ./examples/basic
	@$(MD5SUM) $@

build/examples_basic_pybadge.hex:
	$(TINYGO) build -size short -o $@ -target=pybadge ./examples/basic
	@$(MD5SUM) $@

build/examples_basic_wioterminal.hex:
	$(TINYGO) build -size short -o $@ -target=wioterminal ./examples/basic
	@$(MD5SUM) $@

build/examples_basic_badger2040.hex:
	$(TINYGO) build -size short -o $@ -target=badger2040 ./examples/basic
	@$(MD5SUM) $@

build/examples_colors.hex:
	$(TINYGO) build -size short -o $@ -target=pyportal ./examples/colors
	@$(MD5SUM) $@

build/examples_httpclient.hex:
	$(TINYGO) build -size short -o $@ -target=pyportal -stack-size=4kb ./examples/httpclient
	@$(MD5SUM) $@

build/examples_pybadge.hex:
	$(TINYGO) build -size short -o $@ -target=pybadge ./examples/pybadge
	@$(MD5SUM) $@
