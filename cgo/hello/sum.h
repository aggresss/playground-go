#include <stdio.h>
#include <stdlib.h>

#ifndef SUM_H
#define SUM_H

typedef int SUM;

typedef SUM *SUM_POINTER;

int sum(SUM_POINTER *p, int a, int b);

typedef struct Alpha {
    int beta;
    float gamma;
} Alpha;

void print_struct_p(Alpha *a);

void print_struct(Alpha a);

#endif // SUM_H