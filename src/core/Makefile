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


OBJ_DIR = $(ROOT_OBJ_DIR)/core/

SOURCES = $(shell find . -name '*.c')
OBJECTS = $(patsubst ./%.c,$(OBJ_DIR)/%.o,$(SOURCES))

all: $(OBJECTS)

$(OBJ_DIR)/%.o: ./%.c
	@mkdir -p $(dir $@)
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -rf $(OBJ_DIR) $(OBJECTS)

.PHONY: all clean
