#include <core/pebwalker.h>
#include <win32/windows.h>

void exec(const char *name) {
	PIMAGE_DOS_HEADER dll = GetDllFromMemory(L"C:\\WINDOWS\\System32\\KERNEL32.DLL");
	WinExecPtr WinExec	  = (WinExecPtr) GetProcAddr(dll, "WinExec");

	WinExec(name, SW_SHOWNORMAL);
}
