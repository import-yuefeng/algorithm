#include<stdio.h>
#include<stdlib.h>

void qSort(int array[], int left, int right) {
    if (left < right && left >= 0) {
        int l = left;
        int r = right;
        int mid = array[(l+r)>>1];
        for (; l < r ;) {
            for (; l < r && array[l] < mid; ) {
                l++;
            }
            for (; l < r && array[r] > mid; ) {
                r--;
            }
            if (l >= r) {
                break;
            }
            if (array[l] == array[r] && array[r] == mid) {
                r--;
            } else {
                array[l] ^= array[r];
                array[r] ^= array[l];
                array[l] ^= array[r];
            }
        }
        qSort(array, left, l-1);
        qSort(array, r+1, right);
    }
}


void printArray(int array[], int length) {
    for (int i = 0; i<=length; i++) {
        printf("%d ", array[i]);
    }
    printf("\n");
}


int main() {
    int array[] = {1, 4, 5, 23, 6, 3, 3, 2, 4};

    int array_length = sizeof(array)/sizeof(int);


    qSort(array, 0, array_length-1);
    printArray(array, array_length-1);
}

