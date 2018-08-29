#include <stdio.h>

// START OMIT
int foo() {
    int *a = new(int);  // HL
    return *a;
}
// END OMIT

int main(void) { return 0;}