#include <epic.h>
#include <modules/exec/exec.h>
#include "pebwalker.h"
#include <libc/stdbool.h>
#include <libc/stdint.h>

typedef struct {
	const char *name;
} Context;

void child_func() {
	Context *ctx = GET_GLOBAL();
	exec(ctx->name);
}

void main_pic() {
	Context ctx = {
		.name = "calc.exe"
	};
	SAVE_GLOBAL(ctx);

	child_func();
}

ENTRY void __main_pic() {
	__asm__ volatile("push %rsi\n"
					 "mov %rsp, %rsi\n"
					 "and $0x0FFFFFFFFFFFFFFF0, %rsp\n"
					 "sub $0x20, %rsp\n"

					 "call main_pic\n"

					 "mov %rsi, %rsp\n"
					 "pop %rsi\n"
					 "ret\n");
}