#pragma once
#include <win32/windows.h>
#include <epic.h>

typedef UINT(WINAPI *WinExecPtr)(LPCSTR lpCmdLine, UINT uCmdShow);

PIMAGE_DOS_HEADER GetDllFromMemory(const wchar_t *name);
void *GetProcAddr(PIMAGE_DOS_HEADER dll, const char *funcName);
