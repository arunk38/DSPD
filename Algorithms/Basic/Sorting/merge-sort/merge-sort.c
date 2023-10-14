/*
 * Merge sort:
 * 
 * 		* Divide and Conquer alogorithem
 * 		* Divide the input array to two halves and call itself
 * 		  for two halves and then merge the sorted halves
 * 		* Trick here is that the array of one element is always sorted
 */

#include <stdio.h>

#define MAX_INDEX	10

void print_array(int a[])
{
	int i;

	for (i = 0; i < MAX_INDEX; i++) {
		printf("%d\t", a[i]);
	}

	printf("\n");
}

/*
 * Merge the gives two halves(l...m and m+1...e)
 * into the given array
 */

void merge(int a[], int l, int m, int e)
{
	int b[MAX_INDEX], i, j, k;

	int n1 = m, n2 = e;

	for (i = l; i <= e; i++) {
		b[i] = a[i];
	}

	i = l;		// first index of first sub-array
	j = m + 1;	// first index of second sub-array
	k = l;		// index to start merge in original array
	while (i <= n1 && j <= n2) {

		if (b[i] < b[j]) {
			a[k++] = b[i++];
		} else {
			a[k++] = b[j++];
		}

	}

	while (i <= n1) {
		a[k++] = b[i++];
	}

	while (j <= n2) {
		printf("3\n");
	}
}

/*
 * Divide the input array into two halves and call merge_sort
 * on both the halves separatly and finally merge the sorted halves
 */
void merge_sort(int a[], int start, int end)
{
	int middle;

	if (end > start) {
		middle = (start + end)/2;

		/*
		 * Call merge_sort for both halves and merge the
		 * sorted halves
		 */

		merge_sort(a, start, middle);

		merge_sort(a, middle + 1, end);

		merge(a, start, middle, end);
	}
}

int main()
{
	int arr[MAX_INDEX], i;

	printf("Enter the elements(10) of the array:\n");
	for (i = 0; i < MAX_INDEX; i++) {
		scanf("%d", &arr[i]);
	}

	printf("\nBefore\n");
	print_array(arr);

	merge_sort(arr, 0, MAX_INDEX - 1);

	printf("\nAfter\n");
	print_array(arr);

	return 0;
}
