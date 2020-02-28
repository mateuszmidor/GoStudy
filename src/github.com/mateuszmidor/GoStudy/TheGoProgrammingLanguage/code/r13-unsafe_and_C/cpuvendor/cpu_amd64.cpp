// This file will only be compiled for amd64 architecture.
// We can use C++, but the exposed interface must be plain C as cgo only uderstands C
#include <string>

/**
 * @brief   Get CPU vendor string
 */
extern "C" const char* get_vendor() {
    static std::string result { "____________"}; // placeholder for up to 12 characters

    asm volatile (
         "movq %0, %%rdi;            "  // put buffer address from general register %0 into rdi
         "xor %%rax, %%rax;          "
         "cpuid;                     "  // populate ebx, edx, ecx with vendor string
         "mov %%ebx, (%%rdi);        "  // move vendor string from registers into buffer
         "mov %%edx, 4(%%rdi);       "
         "mov %%ecx, 8(%%rdi);       "
        : // no output used
        : "g" (result.c_str())          // put buffer address into general register %0
        : "memory", "%eax", "%ebx", "%ecx", "%edx", "%rdi"
    );

    return result.c_str();
}