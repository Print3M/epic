#pragma once
#include <win32/windows.h>
#include <epic.h>

HMODULE GetDllFromMemory(const wchar_t *name);

void *GetProcAddr(HMODULE dll, const char *funcName);
