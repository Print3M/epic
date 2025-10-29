#include <epic.h>
#include <win32/windows.h>
#include <libc/stdbool.h>
#include <libc/stdint.h>
#include <libc/string.h>
#include <libc/wchar.h>

PPEB GetPEB(void) {
	uint64_t value = 0;

	// Inline assembly to read from the GS segment
	__asm__ volatile("movq %%gs:%1, %0"
					 : "=r"(value)			   // output
					 : "m"(*(uint64_t *) 0x60) // input
					 :						   // no clobbered registers
	);

	return (PPEB) value;
}

PIMAGE_DOS_HEADER GetDllFromMemory(const wchar_t *name) {
	PPEB peb = GetPEB();

	wchar_t tmp_name[1024];
	wchar_t tmp_dll[1024];

	// wcs_to_lowercase(name, tmp_name);

	// TODO: Split by \\, take the last part.
	// TODO: Test this on Windows 10

	PLIST_ENTRY item		  = peb->Ldr->InMemoryOrderModuleList.Blink;
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

void *GetProcAddr(PIMAGE_DOS_HEADER dll, const char *funcName) {
	// Get address of PE headers
	PBYTE pe_hdrs = (PBYTE) ((PBYTE) dll + dll->e_lfanew);

	// Get Export Address Table RVA
	DWORD eat_rva = *(PDWORD) (pe_hdrs + 0x88);

	// Get address of Export Address Table
	PIMAGE_EXPORT_DIRECTORY eat = (PIMAGE_EXPORT_DIRECTORY) ((PBYTE) dll + eat_rva);

	// Get address of function names table
	PDWORD name_rva = (PDWORD) ((PBYTE) dll + eat->AddressOfNames);

	// Get function name
	uint64_t i = 0;

	do {
		char *tmp = (char *) ((PBYTE) dll + name_rva[i]);

		if (strcmp(tmp, funcName) == 0) {
			break;
		}
		i++;
	} while (true);

	// Get function ordinal
	PWORD ordinals = (PWORD) ((PBYTE) dll + eat->AddressOfNameOrdinals);
	WORD ordinal   = ordinals[i];

	// Get function pointer
	PDWORD func_rvas = (PDWORD) ((PBYTE) dll + eat->AddressOfFunctions);
	DWORD func_rva	 = func_rvas[ordinal];

    void *func_addr = (void *) ((PBYTE) dll + func_rva);

	return func_addr;
}