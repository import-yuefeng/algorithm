#include<stdio.h>
#include<stdlib.h>

void printArray(int array[], int length) {
    for (int i = 0; i<=length; i++) {
        printf("%d ", array[i]);
    }
    printf("\n");
}


typedef struct {
    int top;
    int *data;
    int size;
    int capacity;
} Stack;

Stack *NewStack(int capacity) {
    Stack *res = (Stack*)malloc(sizeof(Stack));
    res->capacity = capacity;
    res->data = (int *)malloc(sizeof(int)*capacity);
    res->top = -1;
    res->size = 0;
    return res;
}

_Bool IsEmpty(Stack* s) {
    return s->size == 0;
}

_Bool IsFull(Stack* s) {
    return s->size == s->capacity;
}

int Push(Stack* s, int x) {
    if (IsFull(s)) {
        return -1;
    }
    s->top++;
    s->data[s->top] = x;
    s->size++;
    return x;
}

int Pop(Stack *s) {
    if (IsEmpty(s)) {
        return -1;
    }
    int x = s->data[s->top];
    s->top--;
    s->size--;
    return x;
}

int GetSize(Stack *sp) {
    Stack s = *sp;
    return s.size;
}

int main() {

    Stack *stack = NewStack(100);
    Push(stack, 10);
    Push(stack, 3);
    Push(stack, 5);
    int res = Pop(stack);
    printf("%d\n", res);

    free(stack);
    stack = NULL;
}

