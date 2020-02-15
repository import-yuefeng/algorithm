#include <stdio.h>
#include <stdlib.h>

void swap(int array[], int l, int r)
{
    int tmp = array[l];
    array[l] = array[r];
    array[r] = tmp;
    return;
}

void perm(int array[], int cur, int end)
{
    if (cur == end)
    {
        for (int i = 0; i <= end; i++)
        {
            printf("%d ", array[i]);
        }
        printf("\n");
    }
    else
    {
        for (int i = cur; i <= end; i++)
        {
            swap(array, i, cur);
            perm(array, cur + 1, end);
            swap(array, cur, i);
        }
    }
}

void main()
{
    int array[] = {1, 2, 3};
    perm(array, 0, 2);
}
