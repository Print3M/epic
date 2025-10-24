export CC := x86_64-w64-mingw32-gcc

export ROOT_OUTPUT_DIR := $(CURDIR)/output
export ROOT_ASSETS_DIR := $(CURDIR)/assets
export ROOT_OBJ_DIR := $(ROOT_OUTPUT_DIR)/objects
export ROOT_ARCHIVE_DIR := $(ROOT_OUTPUT_DIR)/archives


OUTPUT_BIN := $(ROOT_OUTPUT_DIR)/payload.bin
CORE_OBJECTS := $(shell find . -name '*.o')

.PHONY: all clean core modules link loader standalone

all: core modules link loader standalone

core:
	@echo "\n[*] Building src/core/"
	$(MAKE) -C src/core


modules:
	@echo "\n[*] Building src/modules/"
	@for dir in src/modules/*/; do \
		if [ -f "$$dir/Makefile" ]; then \
			$(MAKE) -C "$$dir" || exit 1; \
		fi; \
	done


link:
	@echo "\n[*] Linking PIC payload"
	@mkdir -p $(ROOT_OUTPUT_DIR)
	ld -T $(ROOT_ASSETS_DIR)/linker.ld -o $(OUTPUT_BIN) $(CORE_OBJECTS) $(shell find $(ROOT_ARCHIVE_DIR) -name "*.a")


loader: 
	@echo "\n[*] Compiling loader EXE"
	$(MAKE) -C $(ROOT_ASSETS_DIR)/loader

standalone:
	@echo "\n[*] Compiling standalone EXE"
	@mkdir -p $(ROOT_OUTPUT_DIR)
	$(CC) $(CFLAGS) -o $(ROOT_OUTPUT_DIR)/standalone.exe \
		$(shell find src -name "*.c") \
		-w -Os \
		-Wl,--subsystem,console -static -s
	@echo "Executable created: $@"


clean:
	rm -rf $(ROOT_OUTPUT_DIR)
