/*
 * Insert into a linked list
 */

#include <stdio.h>
#include <stdlib.h>

typedef struct _node {
    int             data;
    struct _node    *next;
}ll_node;

void detectAndCountLoop(ll_node *head)
{
	ll_node *node = head;
}

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

restart_loop:
    printf("Loop start at node num.?");
    scanf("%d", &cnt);
    if (cnt > ele) {
    	printf("Loop out of bounds...");
    	goto restart_loop;
    }
    print_list(head);

    detectAndCountLoop(head);
}
