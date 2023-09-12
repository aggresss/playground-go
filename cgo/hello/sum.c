#include <stdio.h>
#include <stdlib.h>

#include "sum.h"

int sum(SUM_POINTER *p, int a, int b) {
    *p = malloc(sizeof(SUM));
    **p = a + b;
}
