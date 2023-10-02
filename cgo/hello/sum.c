#include "sum.h"

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
