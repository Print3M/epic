/* Single-header ctype.h library - Freestanding, minimal dependencies */
/* All functions implemented inline for PIC shellcode compatibility */
/* Suitable for x86_64 Windows (MinGW-w64) */
#pragma once

/* ============================================================================
 * Character classification functions
 * ============================================================================ */

/* Check if character is alphanumeric (letter or digit) */
static inline int isalnum(int c) {
    return ((c >= 'A' && c <= 'Z') || 
            (c >= 'a' && c <= 'z') || 
            (c >= '0' && c <= '9'));
}

/* Check if character is alphabetic */
static inline int isalpha(int c) {
    return ((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z'));
}

/* Check if character is a control character */
static inline int iscntrl(int c) {
    return ((c >= 0 && c <= 31) || c == 127);
}

/* Check if character is a decimal digit */
static inline int isdigit(int c) {
    return (c >= '0' && c <= '9');
}

/* Check if character is printable (including space) */
static inline int isprint(int c) {
    return (c >= 32 && c <= 126);
}

/* Check if character has graphical representation (printable except space) */
static inline int isgraph(int c) {
    return (c >= 33 && c <= 126);
}

/* Check if character is lowercase letter */
static inline int islower(int c) {
    return (c >= 'a' && c <= 'z');
}

/* Check if character is uppercase letter */
static inline int isupper(int c) {
    return (c >= 'A' && c <= 'Z');
}

/* Check if character is punctuation */
static inline int ispunct(int c) {
    return (isgraph(c) && !isalnum(c));
}

/* Check if character is whitespace */
static inline int isspace(int c) {
    return (c == ' ' || c == '\t' || c == '\n' || 
            c == '\r' || c == '\f' || c == '\v');
}

/* Check if character is hexadecimal digit */
static inline int isxdigit(int c) {
    return ((c >= '0' && c <= '9') || 
            (c >= 'A' && c <= 'F') || 
            (c >= 'a' && c <= 'f'));
}

/* Check if character is blank (space or tab) */
static inline int isblank(int c) {
    return (c == ' ' || c == '\t');
}

/* ============================================================================
 * Character conversion functions
 * ============================================================================ */

/* Convert uppercase letter to lowercase */
static inline int tolower(int c) {
    if (c >= 'A' && c <= 'Z') {
        return c + 32;
    }
    return c;
}

/* Convert lowercase letter to uppercase */
static inline int toupper(int c) {
    if (c >= 'a' && c <= 'z') {
        return c - 32;
    }
    return c;
}

/* ============================================================================
 * Shellcode-specific utility functions
 * ============================================================================ */

/* Check if character is valid ASCII (0-127) */
static inline int isascii(int c) {
    return (c >= 0 && c <= 127);
}

/* Force character to ASCII range */
static inline int toascii(int c) {
    return c & 0x7F;
}

/* Convert hexadecimal character to integer value (0-15) */
static inline int xdigit_to_int(int c) {
    if (c >= '0' && c <= '9') {
        return c - '0';
    }
    if (c >= 'A' && c <= 'F') {
        return c - 'A' + 10;
    }
    if (c >= 'a' && c <= 'f') {
        return c - 'a' + 10;
    }
    return -1; /* Invalid hex digit */
}

/* Convert integer value (0-15) to hexadecimal character (uppercase) */
static inline int int_to_xdigit(int n) {
    if (n >= 0 && n <= 9) {
        return '0' + n;
    }
    if (n >= 10 && n <= 15) {
        return 'A' + (n - 10);
    }
    return -1; /* Invalid value */
}

/* Convert integer value (0-15) to hexadecimal character (lowercase) */
static inline int int_to_xdigit_lower(int n) {
    if (n >= 0 && n <= 9) {
        return '0' + n;
    }
    if (n >= 10 && n <= 15) {
        return 'a' + (n - 10);
    }
    return -1; /* Invalid value */
}

/* Check if character is valid filename character (basic set) */
static inline int isfilename(int c) {
    return (isalnum(c) || c == '_' || c == '-' || c == '.');
}

/* Check if character is valid path separator */
static inline int ispathsep(int c) {
    return (c == '/' || c == '\\');
}

/* ============================================================================
 * Wide character versions (useful for Windows API)
 * ============================================================================ */

/* Check if wide character is alphanumeric (ASCII range only) */
static inline int iswalnum(int c) {
    return isalnum(c);
}

/* Check if wide character is alphabetic (ASCII range only) */
static inline int iswalpha(int c) {
    return isalpha(c);
}

/* Check if wide character is digit (ASCII range only) */
static inline int iswdigit(int c) {
    return isdigit(c);
}

/* Check if wide character is lowercase (ASCII range only) */
static inline int iswlower(int c) {
    return islower(c);
}

/* Check if wide character is uppercase (ASCII range only) */
static inline int iswupper(int c) {
    return isupper(c);
}

/* Check if wide character is whitespace (ASCII range only) */
static inline int iswspace(int c) {
    return isspace(c);
}

/* Convert wide character to lowercase (ASCII range only) */
static inline int towlower(int c) {
    return tolower(c);
}

/* Convert wide character to uppercase (ASCII range only) */
static inline int towupper(int c) {
    return toupper(c);
}
