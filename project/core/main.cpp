#include "pebwalker.h"
#include <epic.h>
#include <libc/stdbool.h>
#include <libc/stdint.h>
#include <modules/hello/hello.h>

typedef struct {
	const char *message;
	void *start;
	void *end;
} Context;

void child_func() {
	auto ctx = (Context *) GET_GLOBAL();
	
	hello::message(ctx->message);
}

SECOND_STAGE void main_pic() {
	Context ctx;
	SAVE_GLOBAL(ctx);

	ctx.message = "Hello EPIC!";

	child_func();
}

FIRST_STAGE void __main_pic() {
	__asm__ volatile("push %rsi\n"
					 "mov %rsp, %rsi\n"
					 "and $0x0FFFFFFFFFFFFFFF0, %rsp\n"
					 "sub $0x20, %rsp\n"

					 "call main_pic\n"

					 "mov %rsi, %rsp\n"
					 "pop %rsi\n"
					 "ret\n");
}