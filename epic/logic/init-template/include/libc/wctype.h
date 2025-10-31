/* Single-header wctype.h library - Freestanding, minimal dependencies */
/* All functions implemented inline for PIC shellcode compatibility */
/* Windows uses UTF-16 (16-bit wchar_t) */
/* Suitable for x86_64 Windows (MinGW-w64) */

#ifndef _WCTYPE_H
#define _WCTYPE_H

#include "wchar.h"
#include "ctype.h"

/* Wide character classification type */
typedef unsigned long wctype_t;

/* Wide character conversion type */
typedef const int *wctrans_t;

/* Wide character integer type (already defined but included for completeness) */
#ifndef WEOF
#define WEOF ((wchar_t)-1)
#endif

/* ============================================================================
 * Wide character classification functions
 * Note: These handle ASCII range only, suitable for most shellcode needs
 * Full Unicode classification would require large lookup tables
 * ============================================================================ */

/* Check if wide character is alphanumeric */
static inline int iswalnum(wint_t wc) {
    if (wc > 127) return 0;  /* Non-ASCII */
    return isalnum((int)wc);
}

/* Check if wide character is alphabetic */
static inline int iswalpha(wint_t wc) {
    if (wc > 127) return 0;
    return isalpha((int)wc);
}

/* Check if wide character is blank (space or tab) */
static inline int iswblank(wint_t wc) {
    return (wc == L' ' || wc == L'\t');
}

/* Check if wide character is control character */
static inline int iswcntrl(wint_t wc) {
    if (wc > 127) return 0;
    return iscntrl((int)wc);
}

/* Check if wide character is decimal digit */
static inline int iswdigit(wint_t wc) {
    return (wc >= L'0' && wc <= L'9');
}

/* Check if wide character has graphical representation */
static inline int iswgraph(wint_t wc) {
    if (wc > 127) return 0;
    return isgraph((int)wc);
}

/* Check if wide character is lowercase */
static inline int iswlower(wint_t wc) {
    return (wc >= L'a' && wc <= L'z');
}

/* Check if wide character is printable */
static inline int iswprint(wint_t wc) {
    if (wc > 127) return 0;
    return isprint((int)wc);
}

/* Check if wide character is punctuation */
static inline int iswpunct(wint_t wc) {
    if (wc > 127) return 0;
    return ispunct((int)wc);
}

/* Check if wide character is whitespace */
static inline int iswspace(wint_t wc) {
    return (wc == L' ' || wc == L'\t' || wc == L'\n' || 
            wc == L'\r' || wc == L'\f' || wc == L'\v');
}

/* Check if wide character is uppercase */
static inline int iswupper(wint_t wc) {
    return (wc >= L'A' && wc <= L'Z');
}

/* Check if wide character is hexadecimal digit */
static inline int iswxdigit(wint_t wc) {
    return ((wc >= L'0' && wc <= L'9') || 
            (wc >= L'A' && wc <= L'F') || 
            (wc >= L'a' && wc <= L'f'));
}

/* ============================================================================
 * Wide character conversion functions
 * ============================================================================ */

/* Convert wide character to lowercase */
static inline wint_t towlower(wint_t wc) {
    if (wc >= L'A' && wc <= L'Z') {
        return wc + 32;
    }
    return wc;
}

/* Convert wide character to uppercase */
static inline wint_t towupper(wint_t wc) {
    if (wc >= L'a' && wc <= L'z') {
        return wc - 32;
    }
    return wc;
}

/* ============================================================================
 * Shellcode-specific utility functions
 * ============================================================================ */

/* Check if wide character is valid ASCII (0-127) */
static inline int iswascii(wint_t wc) {
    return (wc >= 0 && wc <= 127);
}

/* Force wide character to ASCII range */
static inline wint_t toascii_w(wint_t wc) {
    return wc & 0x7F;
}

/* Convert wide hexadecimal character to integer value (0-15) */
static inline int wxdigit_to_int(wint_t wc) {
    if (wc >= L'0' && wc <= L'9') {
        return wc - L'0';
    }
    if (wc >= L'A' && wc <= L'F') {
        return wc - L'A' + 10;
    }
    if (wc >= L'a' && wc <= L'f') {
        return wc - L'a' + 10;
    }
    return -1;
}

/* Convert integer value (0-15) to wide hexadecimal character (uppercase) */
static inline wint_t int_to_wxdigit(int n) {
    if (n >= 0 && n <= 9) {
        return L'0' + n;
    }
    if (n >= 10 && n <= 15) {
        return L'A' + (n - 10);
    }
    return -1;
}

/* Convert integer value (0-15) to wide hexadecimal character (lowercase) */
static inline wint_t int_to_wxdigit_lower(int n) {
    if (n >= 0 && n <= 9) {
        return L'0' + n;
    }
    if (n >= 10 && n <= 15) {
        return L'a' + (n - 10);
    }
    return -1;
}

/* Check if wide character is valid filename character */
static inline int iswfilename(wint_t wc) {
    return (iswalnum(wc) || wc == L'_' || wc == L'-' || wc == L'.');
}

/* Check if wide character is path separator */
static inline int iswpathsep(wint_t wc) {
    return (wc == L'/' || wc == L'\\');
}

/* Check if wide character is null terminator */
static inline int iswnull(wint_t wc) {
    return (wc == L'\0');
}

/* ============================================================================
 * UTF-16 specific utilities for Windows
 * ============================================================================ */

/* Check if wide character is a high surrogate (for UTF-16 encoding) */
static inline int is_high_surrogate(wchar_t wc) {
    return (wc >= 0xD800 && wc <= 0xDBFF);
}

/* Check if wide character is a low surrogate (for UTF-16 encoding) */
static inline int is_low_surrogate(wchar_t wc) {
    return (wc >= 0xDC00 && wc <= 0xDFFF);
}

/* Check if wide character is any surrogate */
static inline int is_surrogate(wchar_t wc) {
    return (wc >= 0xD800 && wc <= 0xDFFF);
}

/* Check if wide character is in Basic Multilingual Plane (BMP) */
static inline int is_bmp(wchar_t wc) {
    return (wc <= 0xFFFF && !is_surrogate(wc));
}

#endif /* _WCTYPE_H */
