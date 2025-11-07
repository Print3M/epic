// EPIC (Extensible Position Independent Code)
//
// Source: github.com/Print3M/epic
// Author: Print3M
//
// Utility functions to be used during MONOLITH compilation.
#pragma once

#ifdef MONOLITH

#include <libc/stddef.h>

extern "C" int printf(const char* __restrict __format, ...);

extern "C" int wprintf(const wchar_t* __restrict __format, ...);

#endif