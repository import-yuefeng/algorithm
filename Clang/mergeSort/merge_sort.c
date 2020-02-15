#include "merge_sort.h"

void merge(int *array, int left, int mid, int right)
{
    int res[right - left + 1];
    int a = left;
    int b = mid + 1;
    int c = 0;
    for (; a <= mid && b <= right; c++)
    {
        if (array[a] > array[b])
        {
            res[c] = array[b];
            b++;
        }
        else
        {
            res[c] = array[a];
            a++;
        }
    }
    for (int i = a; i <= mid; i++)
    {
        res[c] = array[a];
        c++;
    }
    for (int i = b; i <= right; i++)
    {
        res[c] = array[i];
        c++;
    }

    for (int i = 0; i < right - left + 1; i++)
    {
        array[i] = res[i];
    }
    return;
}

void merge_sort(int *array, int left, int right)
{
    if (left < right && left >= 0)
    {
        int mid = (left + right) >> 1;
        merge_sort(array, left, mid);
        merge_sort(array, mid + 1, right);
        merge(array, left, mid, right);
    }
}

int main()
{
    int array[] = {1, 5, 543, 3, 3, 2, 123, 2, 312};
    int len = sizeof(array) / sizeof(int);
    merge_sort(array, 0, len - 1);
    for (int i = 0; i < len; i++)
    {
        printf("%d ", array[i]);
    }
    printf("\n");
}
