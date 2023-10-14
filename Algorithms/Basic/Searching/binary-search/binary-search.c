/*
 * Binary Search
 * 
 * 	# Divide and conquer search
 * 	# Array to be searched has to be sorted, if not sort and
 * 	  then call binary search
 */

#include <stdio.h>

#define MAX_INDEX 10

/*
 * A recursive binary search function, it returns location of x in
 * given (sub-)array a[start...end] is present, otherwise -1
 */

int binary_search(int a[], int start, int end, int x)
{
	if (end >= start) {
		int mid = (start + end)/2;

		/*
		 * If the element is present at the middle itself
		 */

		if (a[mid] == x)
			return mid;

		/*
		 * If element is smaller than mid, then it can only be present
		 * in left subarray
		 */

		if (a[mid] > x)
			return binary_search(a, start, mid-1, x);

		/*
		 * Else the element is in the right sub-array
		 */

		return binary_search(a, mid+1, end, x);
	}

	/*
	 * We reach here when the element is not present in the array
	 */

	return -1;
	
}

int main()
{
	int i, arr[MAX_INDEX], x;

	printf("Enter the elements(10) of the array:\n");
	for (i = 0; i < MAX_INDEX; i++) {
		scanf("%d", &arr[i]);
	}

	printf("\nEnter the element to search:  ");
	scanf("%d", &x);

	i = binary_search(arr, 0, MAX_INDEX-1, x);

	if (i != -1) {
		printf("\nSearch successfull\nelement\tindex\n%d\t%d\n", arr[i], i);
	} else {
		printf("\nRequested element is not present in the array\n");
	}

	return 0;
}
