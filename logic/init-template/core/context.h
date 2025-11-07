// EPIC (Extensible Position Independent Code)
//
// Source: github.com/Print3M/epic
// Author: Print3M
//
#pragma once
#include <epic.h>


typedef struct {
    const char* message;
    void*       pic_start;  // 1st byte of the shellcode
    void*       pic_end;    // 1st byte behind the shellcode
} GlobalCtx;


#define GET_CONTEXT() (GlobalCtx*)GET_GLOBAL();
