CFLAGS = -nostdlib \
         -fPIC \
         -nostartfiles \
         -Os \
         -fno-asynchronous-unwind-tables \
         -ffreestanding \
         -fno-builtin \
         -ffunction-sections \
         -fno-ident \
         -falign-jumps=1 \
         -mno-sse \
         -mno-mmx \
         -mgeneral-regs-only \
         -mno-stack-arg-probe \
         -mno-red-zone \
         -fdiagnostics-color=always \
         -std=c17

AR = x86_64-w64-mingw32-ar

CURRENT_DIR = $(notdir $(CURDIR))
OBJ_DIR = $(ROOT_OBJ_DIR)/modules/$(CURRENT_DIR)
ARCHIVE_DIR = $(ROOT_ARCHIVE_DIR)/modules
ARCHIVE = $(ARCHIVE_DIR)/$(CURRENT_DIR).a

SOURCES = $(shell find . -name '*.c')
OBJECTS = $(patsubst ./%.c,$(OBJ_DIR)/%.o,$(SOURCES))

all: $(ARCHIVE)

$(ARCHIVE): $(OBJECTS)
	@mkdir -p $(ARCHIVE_DIR)
	$(AR) rcs $@ $^

$(OBJ_DIR)/%.o: ./%.c
	@mkdir -p $(dir $@)
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -rf $(OBJ_DIR) $(ARCHIVE)

.PHONY: all clean
