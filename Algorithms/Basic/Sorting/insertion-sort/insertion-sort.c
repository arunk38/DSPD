/*
 * Insertion sort algorithem in C
 */

#include <stdio.h>

#define MAX_INDEX	10

void print_array(int a[])
{
		int i;

		for (i=0; i < MAX_INDEX; i++) {
				printf("%d\t", a[i]);
		}
		printf("\n");
}

void insertion_sort(int a[])
{
		int i, j, key;

		/*
		 * Move elements of a[0...i-1], that are grerater
		 * than key, to onr position ahead of their current
		 * position
		 */

		for (i = 1; i < MAX_INDEX; i++) {
				key = a[i];
				j= i - 1;

				while (j >= 0 && a[j] > key) {
						a[j+1] = a[j];
						j--;
				}
				a[j+1] = key;
		}
}

int main()
{
		int arr[MAX_INDEX], i;

		printf("Enter elements(10) of the array:\n");
		for (i = 0; i < MAX_INDEX; i++) {
				scanf("%d", &arr[i]);
		}

		printf("\nBefore sorting:\n");
		print_array(arr);

		insertion_sort(arr);

		printf("\nAfter sorting:\n");
		print_array(arr);
		return 0;
}
