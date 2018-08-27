// basic file operations
#include <stdio.h>
#include <stdlib.h>


int main(int argc, char *argv[]) {
    FILE *fp;
    for ( ; ; ) {
        for (int i = 1; i < argc; i++) {
            fp = fopen("argv[i]", "r+");
            if (fp == NULL) {
                    fprintf(stderr, "Invalid file %s", argv[i]);
                exit(EXIT_FAILURE);
            }
            printf("Open successful");
            fclose(fp);
        }
    }
    
    return 0;
}