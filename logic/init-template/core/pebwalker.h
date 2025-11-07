// EPIC (Extensible Position Independent Code)
//
// Source: github.com/Print3M/epic
// Author: Print3M
//
// Example implementation of PEB-walk mechanism for fast
// prototyping an EPIC project.
// 
#pragma once
#include <epic.h>
#include <win32/windows.h>

#define MAX_DLL_NAME_SIZE 256

HMODULE GetDllFromMemory(const wchar_t* name);
void* GetProcAddr(HMODULE dll, const char* funcName);
