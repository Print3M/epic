#pragma once
#include "../include/common.h"

FUNC int strcmp(const char *s1, const char *s2);
FUNC int wcscmp(const wchar_t *s1, const wchar_t *s2);

FUNC void strcpy(const char *src, char *dest);
FUNC void wcscpy(const wchar_t *src, wchar_t *dest);

FUNC void str_to_lowercase(const char *src, char *dest);
FUNC void wcs_to_lowercase(const wchar_t *src, wchar_t *dest);

FUNC int str_length(const char *str);