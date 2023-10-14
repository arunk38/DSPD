/*
 * Delete a node from linked list
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

void node_delete(ll_node **head, int data)
{
    if (*head == NULL)
        return;

    ll_node *prev = *head;
    ll_node *node = prev->next;

    if (prev->data == data) {
        *head = node;
        free(prev);
        return;
    }

    while (node) {
        if (node->data == data) {
            prev->next = node->next;
            free(node);
            return;
        }
        node = node->next;
        prev = prev->next;
    }
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

    cnt = 1;
    while (cnt) {
        printf("\nEnter data to delete: ");
        scanf("%d", &data);
        node_delete(&head, data);
        print_list(head);
        printf("Continue removing? (1/0) ");
        scanf("%d", &cnt);
    }
}
