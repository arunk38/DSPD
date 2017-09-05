/*
 * Selection sort algorithem in C
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

/*
 * Swap the values of the input arguments
 */

void swap(int *x, int *y)
{
	int temp = *x;
	*x = *y;
	*y = temp;
}

void selection_sort(int a[])
{
	int i, j, min_idx;

	/*
	 * One by one move boundary of unsorted array
	 */

	for (i = 0; i < MAX_INDEX; i++) {

		/*
		 * Find the minimum from unsorted array
		 */
			
		min_idx = i;

		for(j = i+1; j < MAX_INDEX; j++) {
			if (a[j] < a[min_idx]) {
				min_idx = j;
			}
		}

		/*
		 * Swap the found minimum element with
		 * current element
		 */

		swap(&a[min_idx], &a[i]);
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

	selection_sort(arr);

	printf("\nAfter sorting:\n");
	print_array(arr);
	return 0;
}
