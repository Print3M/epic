/* Minimalistic limits.h - Freestanding, no dependencies */
/* Suitable for x86_64 Windows (MinGW-w64) */
#pragma once

/* Number of bits in a byte */
#define CHAR_BIT 8

/* Minimum and maximum values for char */
#define SCHAR_MIN  (-128)
#define SCHAR_MAX  127
#define UCHAR_MAX  255

/* char is signed by default on x86_64 Windows */
#define CHAR_MIN   SCHAR_MIN
#define CHAR_MAX   SCHAR_MAX

/* Minimum and maximum values for short */
#define SHRT_MIN   (-32768)
#define SHRT_MAX   32767
#define USHRT_MAX  65535

/* Minimum and maximum values for int */
#define INT_MIN    (-2147483647 - 1)
#define INT_MAX    2147483647
#define UINT_MAX   4294967295U

/* Minimum and maximum values for long (32-bit on Windows x64) */
#define LONG_MIN   (-2147483647L - 1)
#define LONG_MAX   2147483647L
#define ULONG_MAX  4294967295UL

/* Minimum and maximum values for long long */
#define LLONG_MIN  (-9223372036854775807LL - 1)
#define LLONG_MAX  9223372036854775807LL
#define ULLONG_MAX 18446744073709551615ULL

/* Maximum number of bytes in a multibyte character */
#define MB_LEN_MAX 5
