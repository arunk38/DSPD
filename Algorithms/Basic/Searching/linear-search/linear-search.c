/*
 * Linear Searching in an array
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
 * Search the whole length of array one element at
 * a time for search element until found.
 */

int linear_search(int a[], int x)
{
	int i;

	for (i = 0; i < MAX_INDEX; i++) {
		if (a[i] == x)
			return i;
	}

	return -1;
}

int main()
{
	int i, arr[MAX_INDEX], x;

	printf("Enter the elemtes(10) of the array:\n");
	for (i = 0; i < MAX_INDEX; i++) {
		scanf("%d", &arr[i]);
	}

	printf("\nEnter the element to search:  ");
	scanf("%d", &x);

	i = linear_search(arr, x);

	if (i != -1) {
		printf("\nSearch successfull\nelement\tindex\n%d\t%d\n", x, i);
	} else {
		printf("Required element is not present in the input array.\n");
	}

	return 0;
}
