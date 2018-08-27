// basic file operations
#include <stdio.h>
#include <stdlib.h>


int main (char *argv, int argc) {
    FILE *fp;
    for (int i = 1; i < argc; i++) {
        fp = fopen(argv[i], "r+");
        ...
        fclose(fp);
    }
    return 0;
}