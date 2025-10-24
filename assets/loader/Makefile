.PHONY: all build

CC := x86_64-w64-mingw32-gcc

OUTPUT_ASSETS_DIR := $(ROOT_OUTPUT_DIR)/assets/
OUTPUT_EXE  	  := $(ROOT_OUTPUT_DIR)/loader.exe
PAYLOAD_BIN 	  := $(ROOT_OUTPUT_DIR)/payload.bin

all: build

build:
	@echo "Creating loader executable..."
	@mkdir -p $(OUTPUT_ASSETS_DIR)
	@# Convert binary to \xHH format and inject into loader.c
	@sed "s/:PAYLOAD:/$$(xxd -p $(PAYLOAD_BIN) | sed 's/../\\\\x&/g' | tr -d '\n')/g" \
		loader.c > $(OUTPUT_ASSETS_DIR)/loader.c
	@# Compile to executable
	@$(CC) -o $(OUTPUT_EXE) $(OUTPUT_ASSETS_DIR)/loader.c -static -s
	@echo "Loader created: $(OUTPUT_EXE)"

clean:
	rm $(OUTPUT_EXE)
