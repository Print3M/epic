// EPIC (Extensible Position Independent Code)
//
// Source: github.com/Print3M/epic
// Author: Print3M
//
#include <core/context.h>
#include <epic.h>
#include <win32/windows.h>

//
// * ======================================================================== *
// |																		  |
// |		    DO NOT TOUCH! The code below is required by EPIC.		      |
// |																		  |
// * ======================================================================== *
//

extern void main_pic();

const char __attribute__((section(".start_addr"))) __pic_start[0] = {};
const char __attribute__((section(".end_addr"))) __pic_end[0]     = {};

// This is the entry function of the payload
EPIC_START void __main_pic() {
    __asm__ volatile("push %rsi\n"
                     "mov %rsp, %rsi\n"
                     "and $0x0FFFFFFFFFFFFFFF0, %rsp\n"
                     "sub $0x20, %rsp\n");

    // Initializing CPU-based global context
    GlobalCtx ctx;
    SAVE_GLOBAL(ctx);

    // Getting context values
    ctx.pic_start = (void*)&__pic_start;
    ctx.pic_end   = (void*)&__pic_end;

    // Starting main execution...
    main_pic();

    __asm__ volatile("mov %rsi, %rsp\n"
                     "pop %rsi\n"
                     "ret\n");
}

#ifdef MONOLITH
void WINAPI WinMain() {
    __main_pic();
}
#endif