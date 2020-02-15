// #include <stdio.h>
// #include <stdlib.h>
#include "heap.h"

Heap *init(int array[], int length)
{
    Heap *h = (Heap *)malloc(sizeof(Heap));
    h->size = 0;
    // hardCode...
    h->capacity = 100;
    for (int i = 0; i < length; i++)
    {
        h->data[i] = array[i];
        h->size++;
    }
    buildMinHeap(h);
    return h;
}

void printHeap(Heap *h)
{
    for (int i = 0; i < h->size; i++)
    {
        printf("%d ", h->data[i]);
    }
    printf("\n");
}

void downNode(Heap *h, int i)
{
    int parent = i;
    int child = parent * 2;
    for (; child < h->size;)
    {
        if (child + 1 < h->size && h->data[child + 1] < h->data[child])
        {
            child++;
        }
        if (h->data[child] >= h->data[parent])
        {
            break;
        }
        int tmp = h->data[child];
        h->data[child] = h->data[parent];
        h->data[parent] = tmp;
        parent = child;
        child = parent * 2;
    }
}

void upNode(Heap *h, int i)
{
    int child = i;
    int parent = child / 2;
    for (; child >= 0;)
    {
        if (child + 1 < h->size && h->data[child + 1] < h->data[child])
        {
            child++;
        }
        if (h->data[child] >= h->data[parent])
        {
            break;
        }
        int tmp = h->data[child];
        h->data[child] = h->data[parent];
        h->data[parent] = tmp;
        child = parent;
        parent = child / 2;
    }
}

void buildMinHeap(Heap *h)
{
    if (h->size <= 1)
    {
        return;
    }
    for (int i = h->size / 2; i >= 0; i--)
    {
        downNode(h, i);
    }
}

int Delete(Heap *h)
{
    if (h->size == 0)
    {
        return -1;
    }
    h->size--;
    int res = h->data[0];
    h->data[0] = h->data[h->size];
    downNode(h, 0);
    return res;
}

int Insert(Heap *h, int x)
{
    if (h->size >= h->capacity)
    {
        return -1;
    }
    h->data[h->size] = x;
    upNode(h, h->size);
    h->size++;
    return 0;
}

void main()
{
    int array[] = {1, 4, 2, 242, 1, 23, 2};
    int length = sizeof(array) / sizeof(int);
    Heap *h = init(array, length);
    for (int i = 0; i < length; i++)
    {
        int res = Delete(h);
        printf("%d\n", res);
    }
}
