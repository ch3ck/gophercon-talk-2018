// basic file operations
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

//START OMIT
#define TIME_TO_WAIT 1 /* wait for one second */
int main() {
    FILE *fp;
    clock_t last = clock();
    char* directory[2] = {"one.txt", "two.txt"};
    for ( ; ; ) {
        clock_t current = clock();
        if (current >= (last + TIME_TO_WAIT + CLOCKS_PER_SEC)) { // HL
            for (int i = 0; i < 2; i++) {
                fp = fopen(directory[i], "r+");
                printf("\nopening %s", directory[i]); // HL
                if (fp == NULL) {
                        fprintf(stderr, "Invalid file %s", directory[i]); // HL
                    exit(EXIT_FAILURE);
                }
                //some FILE processing happens
                fclose(fp);            // HL
                printf("\nclosing %s", directory[i]); // HL
                last = current;
            }
        }//executes every second
    }
    //END OMIT
    return 0;
}