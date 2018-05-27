/*
 * Swap given two nodes in a linked list
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

    while (node) {
        printf("%d -> ", node->data);
        node = node->next;
    }
    printf("\n");
}

void swap_nodes(ll_node **head, int x, int y)
{
	if (x == y)
		return;
	
	ll_node *prevX = NULL, *curX = *head;
	while (curX && curX->data != x) {
		prevX = curX;
		curX = curX->next;
	}

	ll_node *prevY = NULL, *curY = *head;
	while (curY && curY->data != y) {
		prevY = curY;
		curY = curY->next;
	}

	if (curX == NULL || curY == NULL)
		return;

	if (prevX != NULL)
		prevX->next = curY;
	else
		*head = curY;
	
	if (prevY != NULL)
		prevY->next = curX;
	else
		*head = curX;

	ll_node *temp = curY->next;
	curY->next = curX->next;
	curX->next = temp;
}

int main()
{
    ll_node *head = NULL;
    int data, cnt = 1;
    int x, y;

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
        print_list(head);
        printf("Continue adding? (1/0) ");
        scanf("%d", &cnt);
    }
    printf("Enter the date of nodes to swap: ");
    scanf("%d%d", &x, &y);

    printf("\n\n\nBefore Swapping:\n");
    print_list(head);
    swap_nodes(&head, x, y);
    printf("\n\n\nAfter Swapping:\n");
    print_list(head);
}
