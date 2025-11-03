// EPIC (Extensible Position Independent Code)
//
// Source: github.com/Print3M/epic
// Author: Print3M

#pragma once

//
// * ======================================================================== *
// |																		  |
// |		    DO NOT TOUCH! The code below is required by EPIC.		      |
// |																		  |
// * ======================================================================== *
//

// Force compiler and linker not to eliminate the function
#define KEEP	  __attribute__((used))

// Check if is linked at runtime
#define EXISTS(x) ((x) != NULL)

// Functions exported from modules should be marked as weak references
#define MODULE	  __attribute__((weak))

#define __FIRST_STAGE __attribute__((section(".entry"))) __attribute__((naked))

#ifdef __cplusplus

#define FIRST_STAGE extern "C" __FIRST_STAGE 
#define SECOND_STAGE extern "C" KEEP

#else

#define FIRST_STAGE __FIRST_STAGE
#define SECOND_STAGE KEEP

#endif

// CPU-based global variable mechanism. Memory address is stored in a fixed CPU register.
// Usage of this CPU register must be disabled at the compilation level so that our
// "global" pointer is not overwritten.
#define SAVE_GLOBAL(var) __asm__ volatile("mov %0, %%rbx" ::"r"(&var))
#define GET_GLOBAL()                                     \
	({                                                   \
		void *__ret;                                     \
		__asm__ volatile("mov %%rbx, %0" : "=r"(__ret)); \
		__ret;                                           \
	})

