#include <stdlib.h>

extern void goRVExtension(char *output, size_t outputSize, char *input);
extern void goRVExtensionVersion(char *output, size_t outputSize);

#ifdef WIN64
__declspec(dllexport) void RVExtension(char *output, size_t outputSize, char *input) {
	goRVExtension(output, outputSize, input);
}

__declspec(dllexport) void RVExtensionVersion(char *output, size_t outputSize) {
	goRVExtensionVersion(output, outputSize);
}

#else

__declspec(dllexport) void __stdcall _RVExtension(char *output, size_t outputSize, char *input) {
	goRVExtension(output, outputSize, input);
}

__declspec(dllexport) void __stdcall _RVExtensionVersion(char *output, size_t outputSize) {
	goRVExtensionVersion(output, outputSize);
}

#endif