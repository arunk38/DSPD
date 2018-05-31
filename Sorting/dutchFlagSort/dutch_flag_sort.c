/*
 * Program to sort array with 0,1 and 2's in
 * single pass
 */

#include <stdio.h>

void printArray(int arr[], int arr_size)
{
    int i;
    for (i = 0; i < arr_size; i++)
        printf("%d ", arr[i]);
    
    printf("\n");
}

void swap (int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

void sort012(int a[], int arr_size)
{
    int lo = 0;
    int hi = arr_size - 1;
    int mid = 0;

    while (mid <= hi) {

        switch (a[mid]) {

            case 0:
                swap(&a[lo++], &a[mid++]);
                break;

            case 1:
                mid++;
                break;

            case 2:
                swap(&a[mid], &a[hi--]);
                break;
        }
    }
}

int
main()
{

    int arr[] = {0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1};
    int size = sizeof (arr)/sizeof (arr[0]);
    int i;

    printf("array before segregation:\t");
    printArray(arr, size);

    sort012(arr, size);
    
    printf("array After segregation:\t");
    printArray(arr, size);
    return 0;

}
