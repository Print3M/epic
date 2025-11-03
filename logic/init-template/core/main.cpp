#include <epic.h>
#include <core/pebwalker.h>
#include <libc/stdbool.h>
#include <libc/stdint.h>
#include <modules/hello/hello.h>
#include <libc/stdlib.h>

typedef struct {
	const char *message;
	void *pic_start;	// 1st byte of the shellcode
	void *pic_end;      // 1st byte behind the shellcode
} GlobalCtx;

void print_hello() {
	GlobalCtx* ctx = (GlobalCtx *) GET_GLOBAL();
	
	hello::message(ctx->message);
}

SECOND_STAGE void main_pic() {
	GlobalCtx* ctx = (GlobalCtx *) GET_GLOBAL();

	ctx->message = "Hello EPIC!";

	print_hello();
}


//
// * ======================================================================== *
// |																		  |
// |		    DO NOT TOUCH! The code below is required by EPIC.		      |
// |																		  |
// * ======================================================================== *
//
const char __attribute__((section(".start_addr"))) __pic_start[0] = {};
const char __attribute__((section(".end_addr"))) __pic_end[0] = {};

// TODO: Test with C

FIRST_STAGE void __main_pic() {
	__asm__ volatile(
		"push %rsi\n"
		"mov %rsp, %rsi\n"
		"and $0x0FFFFFFFFFFFFFFF0, %rsp\n"
		"sub $0x20, %rsp\n"
	);

	// Initializing CPU-based global context
	GlobalCtx ctx;
	SAVE_GLOBAL(ctx);

	// Getting context values
	ctx.pic_start = (void*) &__pic_start;
	ctx.pic_end = (void*) &__pic_end;

	// Starting main execution...
	main_pic();

	__asm__ volatile(
		"mov %rsi, %rsp\n"
		"pop %rsi\n"
		"ret\n"
	);
}

#ifdef MONOLITH
void WINAPI WinMain() { __main_pic(); }
#endif