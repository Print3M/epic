#include <core/context.h>
#include <core/pebwalker.h>
#include <epic.h>
#include <libc/stdbool.h>
#include <libc/stdint.h>
#include <libc/stdlib.h>
#include <modules/hello/hello.h>

void print_hello() {
    auto ctx = GET_CONTEXT();

    hello::message(ctx->message);
}

extern "C" void main_pic() {
    // Start here...
    auto ctx = GET_CONTEXT();

    ctx->message = "Hello EPIC!";

    print_hello();
}