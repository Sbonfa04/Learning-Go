#include <stdio.h>

int main() {
    int n,i;
    while (i<n) { //while (i<n) is equivalent to for (i=0; i<n; i++)
        printf("ciao \n");
        i++;
    }

    for (i=0; i<n; i++) //for (i=0; i<n; i++) is equivalent to while (i<n)
        printf("ciao \n");

    do {
        printf("ciao \n");
        i++;
    } while (i<n); //do while is equivalent to while
    return 0;

    for (;;) { //infinite loop
        printf("ciao \n");
        i++;
        if (i>=n) break; //break is equivalent to go
    }
    return 0;
}