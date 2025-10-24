#include "../include/common.h"
#include <stddef.h>
#include <stdint.h>

FUNC int wcscmp(const wchar_t *s1, const wchar_t *s2) {
	while (*s1 != L'\0' && *s2 != L'\0') {
		if (*s1 != *s2) {
			return (*s1 < *s2) ? -1 : 1;
		}

		s1++;
		s2++;
	}

	if (*s1 == L'\0' && *s2 == L'\0') {
		return 0;
	}

	return (*s1 == L'\0') ? -1 : 1;
}

FUNC int strcmp(const char *s1, const char *s2) {
	while (*s1 != '\0' && *s2 != '\0') {
		if (*s1 != *s2) {
			return (*s1 < *s2) ? -1 : 1;
		}

		s1++;
		s2++;
	}

	if (*s1 == '\0' && *s2 == '\0') {
		return 0;
	}

	return (*s1 == '\0') ? -1 : 1;
}

FUNC void strcpy(const char *src, char *dest) {
	while (*src) {
		*dest = *src;
		++dest;
		++src;
	}

	*dest = '\0';
}

FUNC void wcscpy(const wchar_t *src, wchar_t *dest) {
	while (*src) {
		*dest = *src;
		++dest;
		++src;
	}

	*dest = L'\0';
}

FUNC void str_to_lowercase(const char *src, char *dest) {
	while (*src) {
		if (*src >= 'A' && *src <= 'Z') {
			*dest = *src + 32;
		}

		++src;
		++dest;
	}

	*dest = '\0';
}

FUNC void wcs_to_lowercase(const wchar_t *src, wchar_t *dest) {
	while (*src != L'\0') {
		if (*src >= L'A' && *src <= L'Z') {
			*dest = *src + 32;
		}

		++src;
		++dest;
	}

	*dest = L'\0';
}

FUNC int str_length(const char *str) {
	int count = 0;

	while (*str) {
		++count;
		++str;
	}

	return count;
}

FUNC int wcs_length(const wchar_t *str) {
	int count = 0;

	while (*str) {
		++count;
		++str;
	}

	return count;
}