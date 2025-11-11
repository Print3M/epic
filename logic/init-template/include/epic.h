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

#define EPIC_START __attribute__((section(".entry"))) __attribute__((naked))

// CPU-based global variable mechanism. Memory address is stored in a fixed CPU register.
// Usage of this CPU register must be disabled at the compilation level so that our
// "global" pointer is not overwritten.
static inline void SAVE_GLOBAL(void *var) {
	__asm__ volatile("mov %0, %%rbx" ::"r"(var));
}

static inline void *GET_GLOBAL() {
	void *__ret;                                    
	__asm__ volatile("mov %%rbx, %0" : "=r"(__ret)); 
	
	return __ret;  
}