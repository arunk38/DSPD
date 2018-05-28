/*
 * Swap elements of a linked list pairwise
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

void swap(int *a, int *b)
{
    int t = *a;
    *a = *b;
    *b = t;
}

void pair_swap(ll_node *head)
{
    ll_node *temp = head;
    while (temp != NULL && temp->next != NULL) {
        swap(&temp->data, &temp->next->data);
        temp = temp->next->next;
    }
	
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

    printf("\n\n\nBefore Swapping:\n");
    print_list(head);
    pair_swap(head);
    printf("\n\n\nAfter Swapping:\n");
    print_list(head);
}
