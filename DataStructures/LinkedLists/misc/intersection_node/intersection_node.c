/*
 * Find intersection of two linked lists
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
        printf("[%p    %3d] -> ", node, node->data);
        node = node->next;
    }
    printf("NULL\n");
}

int getNodeCount(ll_node *h)
{
    int c = 0;

    while (h != NULL) {
        c++;
        h = h->next;
    }
    return c;
}

void getIntersectionNode(ll_node *h1, ll_node *h2)
{
    int c1 = getNodeCount(h1);
    int c2 = getNodeCount(h2);
    int d = 0, i = 0, small_count;
    ll_node *large_head, *small_head;

    if (c1 > c2) {
        d = c1 - c2;
        large_head = h1;
        small_head = h2;
        small_count = c2;
    } else {
        d = c2 - c1;
        large_head = h2;
        small_head = h1;
        small_count = c1;
    }

    for (i = 0; i < d; i++) {
        large_head = large_head->next;
    }

    for (i = 0; large_head != small_head && i < small_count; i++) {
        large_head = large_head->next;
        small_head = small_head->next;
    }

    printf("\n\nIntersection node: [%p    %3d]\n", large_head, large_head->data);


    
}

int main()
{
    /*
    Create two linked lists

    1st 3->6->9->15->30
    2nd 10->15->30

    15 is the intersection point
    */

    ll_node* newNode;
    ll_node* head1 = (ll_node*) malloc(sizeof(ll_node));
    head1->data  = 10;

    ll_node* head2 = (ll_node*) malloc(sizeof(ll_node));

    head2->data  = 3;

    newNode = (ll_node*) malloc(sizeof(ll_node));
    newNode->data = 6;
    head2->next = newNode;

    newNode = (ll_node*) malloc(sizeof(ll_node));
    newNode->data = 9;
    head2->next->next = newNode;

    newNode = (ll_node*) malloc(sizeof(ll_node));
    newNode->data = 15;
    head1->next = newNode;
    head2->next->next->next  = newNode;

    newNode = (ll_node*) malloc(sizeof(ll_node));
    newNode->data = 30;
    head1->next->next= newNode;

    head1->next->next->next = NULL;

    print_list(head1);
    print_list(head2);

    getIntersectionNode(head2, head1);

}
