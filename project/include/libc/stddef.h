#pragma once

/* Size type */
typedef unsigned long long size_t;

/* Pointer difference type */
typedef long long ptrdiff_t;

/* Wide character type */
#ifndef __cplusplus
typedef unsigned short wchar_t;
#endif

/* Maximum alignment type */
typedef long double max_align_t;

/* Null pointer constant */
#ifndef NULL
#ifdef __cplusplus
#define NULL 0
#else
#define NULL ((void *) 0)
#endif
#endif

/* Offset of member in structure */
#ifndef offsetof
#define offsetof(type, member) ((size_t) & (((type *) 0)->member))
#endif
