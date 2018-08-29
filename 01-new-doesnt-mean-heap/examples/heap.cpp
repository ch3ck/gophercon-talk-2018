#include <stdio.h>

// START OMIT
int foo() {
    // This is a memory leak below
    int *a = new(int);  // HL
    return *a;
}
// END OMIT

int main(void) { return 0;}