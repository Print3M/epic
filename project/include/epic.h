#pragma once

// TODO: Add documentation here

#define ENTRY	  __attribute__((section(".entry"))) __attribute__((naked))
#define KEEP	  __attribute__((used))
#define MODULE	  __attribute__((weak))
#define EXISTS(x) ((x) != NULL)

// CPU-based global variable mechanism. Memory address is stored in a fixed CPU register.
// Usage of this CPU register must be disabled at the compilation level so that our
// "global" pointer is not overwritten.
#define SAVE_GLOBAL(var) __asm__ volatile("mov %0, %%rbx" :: "r"(&var))
#define GET_GLOBAL() ({ void* __ret; __asm__ volatile("mov %%rbx, %0" : "=r"(__ret)); __ret; })