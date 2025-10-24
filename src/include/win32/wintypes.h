/*
 * wintypes.h
 *
 * Windows Data Types

 * Copyright 2020 (c) Samantaz Fox
 *
 * This file is in the public domain.
 * Feel free to copy, modify, redistribute it!
 *
*/

#pragma once
#include <stddef.h>
#include <stdint.h>

/*
 *
 * https://stackoverflow.com/questions/384502/what-is-the-bit-size-of-long-on-64-bit-windows
 *
 *
 *   Type                        | S/U | x86    | x64
 *  ----------------------------+-----+--------+-------
 *  BYTE, BOOLEAN               | U   | 8 bit  | 8 bit
 *  ----------------------------+-----+--------+-------
 *  SHORT                       | S   | 16 bit | 16 bit
 *  USHORT, WORD                | U   | 16 bit | 16 bit
 *  ----------------------------+-----+--------+-------
 *  INT, LONG                   | S   | 32 bit | 32 bit
 *  UINT, ULONG, DWORD          | U   | 32 bit | 32 bit
 *  ----------------------------+-----+--------+-------
 *  INT_PTR, LONG_PTR, LPARAM   | S   | 32 bit | 64 bit
 *  UINT_PTR, ULONG_PTR, WPARAM | U   | 32 bit | 64 bit
 *  ----------------------------+-----+--------+-------
 *  LONGLONG                    | S   | 64 bit | 64 bit
 *  ULONGLONG, QWORD            | U   | 64 bit | 64 bit
 *
 */

/*********************************************\
 *
 * Defines
 *
\*********************************************/

#define CONST const

#define WINAPI	 __stdcall
#define CALLBACK __stdcall
#define APIENTRY WINAPI

/*********************************************\
 *
 * Integers - unsigned
 *
\*********************************************/

typedef unsigned char BYTE;
typedef BYTE *LPBYTE;

typedef unsigned char UCHAR;
typedef UCHAR *PUCHAR;

typedef unsigned short USHORT;
typedef USHORT *PUSHORT;

typedef unsigned int UINT;
typedef UINT *PUINT;

typedef unsigned char UINT8;
typedef unsigned short UINT16;
typedef unsigned int UINT32;
typedef uint64_t UINT64;

typedef UINT8 *PUINT8;
typedef UINT16 *PUINT16;
typedef UINT32 *PUINT32;
typedef UINT64 *PUINT64;

typedef unsigned long ULONG;
typedef ULONG *PULONG;

typedef unsigned int ULONG32;
typedef ULONG32 *PULONG32;

typedef uint64_t ULONG64;
typedef ULONG64 *PULONG64;

#if !defined(_M_IX86)
typedef uint64_t ULONGLONG;
#else
typedef double ULONGLONG;
#endif

typedef ULONGLONG *PULONGLONG;

/*********************************************\
 *
 * Integers - signed
 *
\*********************************************/

typedef BYTE *PBYTE;

typedef char CHAR;
typedef CHAR *PCHAR;

typedef short SHORT;
typedef SHORT *PSHORT;

typedef int INT;
typedef int *PINT;
typedef int *LPINT;

typedef signed char INT8;
typedef signed short INT16;
typedef signed int INT32;
typedef int64_t INT64;

typedef INT8 *PINT8;
typedef INT16 *PINT16;
typedef INT32 *PINT32;
typedef INT64 *PINT64;

typedef long LONG;
typedef LONG *PLONG;
typedef long *LPLONG;

typedef signed int LONG32;
typedef LONG32 *PLONG32;

typedef int64_t LONG64;
typedef LONG64 *PLONG64;

#if !defined(_M_IX86)
typedef int64_t LONGLONG;
#else
typedef double LONGLONG;
#endif

typedef LONGLONG *PLONGLONG;

/*********************************************\
 *
 * Boolean types
 *
\*********************************************/

typedef int BOOL;
typedef BOOL *PBOOL;
typedef BOOL *LPBOOL;

typedef BYTE BOOLEAN;
typedef BOOLEAN *PBOOLEAN;

/*********************************************\
 *
 * Floating point
 *
\*********************************************/

typedef float FLOAT;
typedef FLOAT *PFLOAT;

/*********************************************\
 *
 * Words / Double Words / Quad Words
 *
\*********************************************/

typedef unsigned short WORD;
typedef WORD *PWORD;
typedef WORD *LPWORD;

typedef unsigned long DWORD;
typedef DWORD *PDWORD;
typedef DWORD *LPDWORD;

typedef uint64_t QWORD;

typedef unsigned int DWORD32;
typedef DWORD32 *PDWORD32;

typedef uint64_t DWORD64;
typedef DWORD64 *PDWORD64;

typedef uint64_t DWORDLONG;
typedef DWORDLONG *PDWORDLONG;

/*********************************************\
 *
 * (Double/Quad)Word based
 *
\*********************************************/

typedef WORD ATOM;
typedef WORD LANGID;

typedef DWORD COLORREF;
typedef DWORD *LPCOLORREF;

typedef DWORD LCID;
typedef PDWORD PLCID;

typedef DWORD LCTYPE;
typedef DWORD LGRPID;

/*********************************************\
 *
 * Pointers
 *
\*********************************************/

#if defined(_WIN64)
#define POINTER_32 __ptr32
#else
#define POINTER_32
#endif

#if (_MSC_VER >= 1300)
#define POINTER_64 __ptr64
#else
#define POINTER_64
#endif

#define POINTER_SIGNED	 __sptr
#define POINTER_UNSIGNED __uptr

#define VOID void
typedef void *PVOID;
typedef void *LPVOID;
typedef CONST void *LPCVOID;

#if defined(_WIN64)
typedef int HALF_PTR;
typedef int64_t INT_PTR;
typedef int64_t LONG_PTR;
#else
typedef short HALF_PTR;
typedef int INT_PTR;
typedef long LONG_PTR;
#endif

typedef HALF_PTR *PHALF_PTR;
typedef INT_PTR *PINT_PTR;
typedef LONG_PTR *PLONG_PTR;

#ifdef _WIN64
typedef unsigned int UHALF_PTR;
typedef uint64_t UINT_PTR;
typedef uint64_t ULONG_PTR;
#else
typedef unsigned short UHALF_PTR;
typedef unsigned int UINT_PTR;
typedef unsigned long ULONG_PTR;
#endif

typedef UHALF_PTR *PUHALF_PTR;
typedef UINT_PTR *PUINT_PTR;
typedef ULONG_PTR *PULONG_PTR;

typedef ULONG_PTR DWORD_PTR;
typedef DWORD_PTR *PDWORD_PTR;

/*********************************************\
 *
 * Handles
 *
\*********************************************/

typedef PVOID HANDLE;
typedef HANDLE *PHANDLE;
typedef HANDLE *LPHANDLE;

typedef HANDLE HACCEL;
typedef HANDLE HBITMAP;
typedef HANDLE HBRUSH;
typedef HANDLE HCOLORSPACE;
typedef HANDLE HCONV;
typedef HANDLE HCONVLIST;
typedef HANDLE HDC;
typedef HANDLE HDDEDATA;
typedef HANDLE HDESK;
typedef HANDLE HDROP;
typedef HANDLE HDWP;
typedef HANDLE HENHMETAFILE;
typedef HANDLE HFONT;
typedef HANDLE HGDIOBJ;
typedef HANDLE HGLOBAL;
typedef HANDLE HHOOK;
typedef HANDLE HICON;
typedef HANDLE HINSTANCE;
typedef HANDLE HKL;
typedef HANDLE HLOCAL;
typedef HANDLE HMENU;
typedef HANDLE HMETAFILE;
typedef HANDLE HPALETTE;
typedef HANDLE HPEN;
typedef HANDLE HRGN;
typedef HANDLE HRSRC;
typedef HANDLE HSZ;
typedef HANDLE HWND;
typedef HANDLE SC_HANDLE;
typedef HANDLE SERVICE_STATUS_HANDLE;
typedef HANDLE WINSTA;

typedef HINSTANCE HMODULE;
typedef HICON HCURSOR;

typedef HANDLE HKEY;
typedef HKEY *PHKEY;

/*********************************************\
 *
 * Unicode and ANSI C strings
 *
\*********************************************/

typedef char CCHAR;

typedef wchar_t WCHAR;
typedef WCHAR *PWCHAR;

#ifdef UNICODE
typedef WCHAR TBYTE;
typedef WCHAR TCHAR;
#else
typedef unsigned char TBYTE;
typedef char TCHAR;
#endif

typedef TBYTE *PTBYTE;
typedef TCHAR *PTCHAR;

typedef CHAR *PSTR;
typedef CHAR *LPSTR;

typedef WCHAR *PWSTR;
typedef WCHAR *LPWSTR;

typedef CONST CHAR *PCSTR;
typedef CONST WCHAR *PCWSTR;

typedef CONST CHAR *LPCSTR;
typedef CONST WCHAR *LPCWSTR;

#ifdef UNICODE
typedef LPWSTR PTSTR;
typedef LPWSTR LPTSTR;

typedef LPCWSTR PCTSTR;
typedef LPCWSTR LPCTSTR;
#else
typedef LPSTR PTSTR;
typedef LPSTR LPTSTR;

typedef LPCSTR PCTSTR;
typedef LPCSTR LPCTSTR;
#endif

typedef struct _UNICODE_STRING {
	USHORT Length;
	USHORT MaximumLength;
	PWSTR Buffer;
} UNICODE_STRING;

typedef UNICODE_STRING *PUNICODE_STRING;
typedef const UNICODE_STRING *PCUNICODE_STRING;

/*********************************************\
 *
 * Miscellanous
 *
\*********************************************/

typedef ULONG_PTR SIZE_T;
typedef SIZE_T *PSIZE_T;

typedef LONG_PTR SSIZE_T;
typedef SSIZE_T *PSSIZE_T;

typedef LONG_PTR LPARAM;
typedef UINT_PTR WPARAM;

typedef int HFILE;
typedef LONG HRESULT;
typedef LONG_PTR LRESULT;

typedef LPVOID SC_LOCK;

typedef LONGLONG USN;
