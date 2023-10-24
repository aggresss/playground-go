#include <stdlib.h>
#include "sum.h"

int global_var = 5;

int sum(SUM_POINTER *p, int a, int b) {
  *p = malloc(sizeof(SUM));
  **p = a + b;
  return 0;
}

void print_struct_p(Alpha *a) {
  printf("print_struct: %d,%f\n", a->beta, a->gamma);
}

void print_struct(Alpha a) {
  printf("print_struct_p: %d,%f\n", a.beta, a.gamma);
}

Alpha return_struct(void) {
  Alpha a = { 5, 6.0 };
  return a;
}

Alpha* return_struct_point(void) {
  Alpha *a = (Alpha *)malloc(sizeof(Alpha));
  a->beta = 7;
  a->gamma = 8.0;
  return a;
}

void* print_string(char* s) {
    printf("%s\n", s);
    return NULL;
}