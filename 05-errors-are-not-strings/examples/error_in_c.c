#include <stdio.h> /* for fprintf and stderr */
#include <stdlib.h> /* for exit */
int main( void )
{
   int dividend = 50;
   int divisor = 0;
   int quotient;

   if (divisor == 0) {
       /* Example error handling. 
        * Writing a message to stderr, and
        * exiting with failure.
        */
       fprintf(stderr, "Division by zero! Aborting...\n");
       exit(EXIT_FAILURE); /* indicate failure.*/
   }
    
   quotient = dividend / divisor;
   exit(EXIT_SUCCESS); /* indicate success.*/
}