/*
 * Bubble sort algorithem in C.
 */

#include <stdio.h>

void print_array(int a[])
{
	int i = 0;

	for (i = 0; i < 10; i++) {
		printf("%d\t", a[i]);
	}

	printf("\n");
}

int main()
{
	int array[10], i = 0, j;

	printf("Enter the numbers(10) to sort\n");

	for (i = 0; i < 10; i++) {

		/*
		 * Scan the integers and store them
		 * in the array
		 */

		scanf("%d", &array[i]);
	}

	printf("\nBefore Sorting:\n");
	print_array(array);

	/*
	 * Pick an element(i) of array one-by-one and check
	 * with next element(j), if greater swap those elements
	 * else move on to next comparision
	 */

	for (i = 0; i < 10; i++) {
		for (j = i+1; j < 10; j++) {
			if (array[i] > array[j]) {

				/*
				 * Elements out-of-order, swap them.
				 */

				int temp = array[i];
				array[i] = array[j];
				array[j] = temp;
			}
		}
	}

	printf("\nAfter Sorting:\n");
	print_array(array);

	return 0;
}
