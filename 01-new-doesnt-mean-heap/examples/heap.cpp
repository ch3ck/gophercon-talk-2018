#include <stdio.h>

// START OMIT
int foo() {
	int *a = new(int);
        return *a;
}
// END OMIT

int main(void) { return 0;}