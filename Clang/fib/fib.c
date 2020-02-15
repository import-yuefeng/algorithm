#include "fib.h"

void main()
{
    // int res = fib_1(100);
    // printf("%d\n", res);
    int array[100];
    for (int i; i < 100; i++)
    {
        array[i] = 0;
    }
    int res = fib_2(20, array);
    printf("%d\n", res);
}

int fib_1(int step)
{
    if (step == 0 || step == 1)
    {
        return 1;
    }
    return fib_1(step - 1) + fib_1(step - 2);
}

int fib_2(int step, int array[])
{
    if (step == 0 || step == 1)
    {
        return 1;
    }
    if (array[step] != 0)
    {
        return array[step];
    }
    array[step] = fib_2(step - 1, array) + fib_2(step - 2, array);
    return array[step];
}