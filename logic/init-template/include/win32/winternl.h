// EPIC (Extensible Position Independent Code)
//
// Source: github.com/Print3M/epic
// Author: Print3M
//
#pragma once
#include "wintypes.h"

#define CONTAINING_RECORD(address, type, field) \
    ((type*)((char*)(address) - (ULONG_PTR)(&((type*)0)->field)))

// https://www.vergiliusproject.com/kernels/x64/windows-11/25h2/_LIST_ENTRY
typedef struct _LIST_ENTRY {
    struct _LIST_ENTRY* Flink;  // 0x0
    struct _LIST_ENTRY* Blink;  // 0x8
} LIST_ENTRY, *PLIST_ENTRY;

// https://www.vergiliusproject.com/kernels/x64/windows-11/25h2/_PEB_LDR_DATA
typedef struct _PEB_LDR_DATA {
    ULONG              Length;                           // 0x0
    UCHAR              Initialized;                      // 0x4
    VOID*              SsHandle;                         // 0x8
    struct _LIST_ENTRY InLoadOrderModuleList;            // 0x10
    struct _LIST_ENTRY InMemoryOrderModuleList;          // 0x20
    struct _LIST_ENTRY InInitializationOrderModuleList;  // 0x30
    VOID*              EntryInProgress;                  // 0x40
    UCHAR              ShutdownInProgress;               // 0x48
    VOID*              ShutdownThreadId;                 // 0x50
} PEB_LDR_DATA, *PPEB_LDR_DATA;

// https://www.vergiliusproject.com/kernels/x64/windows-11/25h2/_LDR_DATA_TABLE_ENTRY
typedef struct _LDR_DATA_TABLE_ENTRY {
    struct _LIST_ENTRY     InLoadOrderLinks;            // 0x0
    struct _LIST_ENTRY     InMemoryOrderLinks;          // 0x10
    struct _LIST_ENTRY     InInitializationOrderLinks;  // 0x20
    VOID*                  DllBase;                     // 0x30
    VOID*                  EntryPoint;                  // 0x38
    ULONG                  SizeOfImage;                 // 0x40
    struct _UNICODE_STRING FullDllName;                 // 0x48
    struct _UNICODE_STRING BaseDllName;                 // 0x58
    ULONG                  Flags;                       // 0x68

    //...
} LDR_DATA_TABLE_ENTRY, *PLDR_DATA_TABLE_ENTRY;

// https://www.vergiliusproject.com/kernels/x64/windows-11/25h2/_PEB
typedef struct _PEB {
    UCHAR                                InheritedAddressSpace;     // 0x0
    UCHAR                                ReadImageFileExecOptions;  // 0x1
    UCHAR                                BeingDebugged;             // 0x2
    UCHAR                                BitField;                  // 0x3
    UCHAR                                Padding0[4];               // 0x4
    VOID*                                Mutant;                    // 0x8
    VOID*                                ImageBaseAddress;          // 0x10
    struct _PEB_LDR_DATA*                Ldr;                       // 0x18
    struct _RTL_USER_PROCESS_PARAMETERS* ProcessParameters;         // 0x20
    VOID*                                SubSystemData;             // 0x28
    VOID*                                ProcessHeap;               // 0x30

    // ...
} PEB, *PPEB;

// https://www.vergiliusproject.com/kernels/x64/windows-11/25h2/_IMAGE_DOS_HEADER
typedef struct _IMAGE_DOS_HEADER {
    USHORT e_magic;     // 0x0
    USHORT e_cblp;      // 0x2
    USHORT e_cp;        // 0x4
    USHORT e_crlc;      // 0x6
    USHORT e_cparhdr;   // 0x8
    USHORT e_minalloc;  // 0xa
    USHORT e_maxalloc;  // 0xc
    USHORT e_ss;        // 0xe
    USHORT e_sp;        // 0x10
    USHORT e_csum;      // 0x12
    USHORT e_ip;        // 0x14
    USHORT e_cs;        // 0x16
    USHORT e_lfarlc;    // 0x18
    USHORT e_ovno;      // 0x1a
    USHORT e_res[4];    // 0x1c
    USHORT e_oemid;     // 0x24
    USHORT e_oeminfo;   // 0x26
    USHORT e_res2[10];  // 0x28
    LONG   e_lfanew;    // 0x3c
} IMAGE_DOS_HEADER, *PIMAGE_DOS_HEADER;

// https://pinvoke.net/default.aspx/Structures.IMAGE_EXPORT_DIRECTORY
typedef struct _IMAGE_EXPORT_DIRECTORY {
    DWORD Characteristics;
    DWORD TimeDateStamp;
    WORD  MajorVersion;
    WORD  MinorVersion;
    DWORD Name;
    DWORD Base;
    DWORD NumberOfFunctions;
    DWORD NumberOfNames;
    DWORD AddressOfFunctions;     // RVA from base of image
    DWORD AddressOfNames;         // RVA from base of image
    DWORD AddressOfNameOrdinals;  // RVA from base of image
} IMAGE_EXPORT_DIRECTORY, *PIMAGE_EXPORT_DIRECTORY;
