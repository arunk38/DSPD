/*
 * Reverse thelinked list
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

void ll_reverse(ll_node **head)
{
    if (*head == NULL)
        return;

    ll_node *prev = NULL, *curr = *head,  *next;

    while (curr != NULL) {
        next = curr->next;
        curr->next = prev;
        prev = curr;
        curr = next;
    }
    *head = prev;
}

int main()
{
    ll_node *head = NULL;
    int data, cnt = 1;

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

    ll_reverse(&head);
    print_list(head);
    return 0;
}
