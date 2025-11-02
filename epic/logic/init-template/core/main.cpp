#include <epic.h>
#include <core/pebwalker.h>
#include <libc/stdbool.h>
#include <libc/stdint.h>
#include <modules/hello/hello.h>

typedef struct {
	const char *message;
	void *start;
	void *end;
} GlobalCtx;

void child_func() {
	GlobalCtx* ctx = (GlobalCtx *) GET_GLOBAL();
	
	hello::message(ctx->message);
}

// EPIC: Entry point
SECOND_STAGE void main_pic() {
	GlobalCtx ctx;
	SAVE_GLOBAL(ctx);
	
	ctx.message = "Hello EPIC!";
	
	child_func();
}

// TODO: Add banner "Don't touch"

// EPIC: Do not remove!
extern "C" const size_t __shellcode_size;
extern "C" const void * __shellcode_start;
extern "C" void * __shellcode_end;

FIRST_STAGE void __main_pic() {
	// TODO: Get start address of your payload
	__asm__ volatile(
		"push %rsi\n"
		"mov %rsp, %rsi\n"
		"and $0x0FFFFFFFFFFFFFFF0, %rsp\n"
		"sub $0x20, %rsp\n"

		"call main_pic\n"

		"mov %rsi, %rsp\n"
		"pop %rsi\n"
		"ret\n"
	);
}

// EPIC: Do not remove!
#ifdef MONOLITH
void WINAPI WinMain() { __main_pic(); }
#endif