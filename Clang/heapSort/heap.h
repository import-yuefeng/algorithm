#include <stdio.h>
#include <stdlib.h>
typedef struct
{
    int data[100];
    int size;
    int capacity;
} Heap;

Heap *init(int array[], int length);
int Delete(Heap *h);
int Insert(Heap *h, int x);
void upNode(Heap *h, int i);
void downNode(Heap *h, int i);
void printHeap(Heap *h);
void buildMinHeap(Heap *h);
