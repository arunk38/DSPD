/*
 * Remove duplicates from unsorted linked list
 */

#include <stdio.h>
#include <stdlib.h>

typedef struct _node {
    int             data;
    struct _node    *next;
}ll_node;

void print_list(ll_node *head)
{
    ll_node *node = head;
    int i = 0;

    while (node && i < 10) {
        printf("%d -> ", node->data);
        node = node->next;
        i++;
    }
    printf("\n");
}

void removeDuplicates(ll_node *head)
{
	ll_node *ptr1, *ptr2, *dup;
	ptr1 = head;

	while (ptr1 != NULL && ptr1->next != NULL) {
		ptr2 = ptr1;

		while (ptr2->next != NULL) {
			if (ptr1->data == ptr2->next->data) {
				dup = ptr2->next;
				ptr2->next = ptr2->next->next;
				free(dup);
			} else {
				ptr2 = ptr2->next;
			}
		}
		ptr1 = ptr1->next;
	}
}

int main()
{
    ll_node *head = NULL, *end_node;
    int data, cnt = 1, ele = 0;

    while (cnt) {
        printf("Enter data to insert: ");
        ll_node *node = (ll_node*)malloc(sizeof(ll_node));
        scanf("%d", &data);
        node->data = data;
        node->next = NULL;
        
        if (head == NULL) {
            head = node;
        } else {
            node->next = head;
            head = node;
        }
        ele++;
        print_list(head);
        printf("Continue adding? (1/0) ");
        scanf("%d", &cnt);
    }

    removeDuplicates(head);
    print_list(head);
}
