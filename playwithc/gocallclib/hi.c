// gcc -c -fPIC -o hi.o hi.c
// gcc -shared -o libhi.so hi.o
/*
 * hi.c
 * created on: July 1, 2017
 *      author: mark
 */

#include <stdio.h>

void hi() {
    printf("Hello Cgo!\n");
}