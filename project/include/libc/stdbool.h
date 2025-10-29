/* Minimalistic stdbool.h - Freestanding, no external dependencies */
/* Suitable for x86_64 Windows (MinGW-w64) and GCC */

#pragma once

#ifndef __cplusplus

/* Define bool as the C99 _Bool type */
#define bool _Bool

/* Boolean constants */
#define true 1
#define false 0

#else /* __cplusplus */

/* C++ has bool, true, false as keywords */
/* Define _Bool for compatibility */
#define _Bool bool

#endif /* __cplusplus */

/* Signal that all definitions are present */
#define __bool_true_false_are_defined 1
