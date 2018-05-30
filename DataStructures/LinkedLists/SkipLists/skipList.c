/*
 * Insert into a skip list
 */

#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

#define SKIPLIST_MAX_LEVEL 6

typedef struct _node {
    int             data;
    struct _node    **forward;      // Array to hold pointers to node of different level
}sl_node;

typedef struct skiplist {
    int             level;
    sl_node         *header;
} skiplist;

/*
 * print a skip list
 */

void
ak_sl_print(
    skiplist        list)
{
    int i, level = list.level;
    printf("\n\t*****Skip List*****\n");
    printf("Total levels in skip list:\t%d\n", list.level);

    printf("\nData in skip list\n\n");

    for (i = level; i > 0; i--) {
        sl_node *node = list.header->forward[i];

        printf("Level %d:", (i - 1));
        while(node != NULL) {
            printf("\t%d\t", node->data);
            node = node->forward[i];
        }
        printf("\n");
    }
}

/*
 * Initalize a skip list
 */

void
ak_sl_init(
    skiplist       *list)
{
    int i;

    /*
     * create and attach a header node to the list
     */

    sl_node *header = (sl_node *) malloc(sizeof (sl_node));
    list->header = header;

    header->data = INT_MAX;

    /*
     * Create and initialize all forward pointers of header to NULL
     */

    header->forward = (sl_node **) malloc(sizeof (sl_node *) * (SKIPLIST_MAX_LEVEL + 1));
    for (i = 0; i <= SKIPLIST_MAX_LEVEL ; i++) {
        header->forward[i] = NULL;
    }

    list->level = 1;

}

/*
 * create random level for skiplist node
 */

int
randomLevel()
{
    int lvl = 1;
    while (rand() < RAND_MAX / 2 && lvl < SKIPLIST_MAX_LEVEL)
        lvl++;

    return lvl;


}

/*
 * Insert into skip list at appropriate level(s)
 */

void
_ak_sl_insert(
    skiplist        *list,
    int             data)
{
    sl_node *update[SKIPLIST_MAX_LEVEL + 1];
    sl_node *x = list->header;

    int i, level;

    /*
     * Traverse linked list(forward) at each level and store
     * the previous node pointer at which new data has to be
     * inserted
     */

    for (i = list->level; i >= 1; i--) {

        /*
         * Traverse this level upto node with data greater than
         * the new `data`
         */

        while (x->forward[i] && (x->forward[i]->data < data))
            x = x->forward[i];

        update[i] = x;
    }

    if (x->forward[1] && x->forward[1]->data == data) {
        return;
    }


    /*
     * get level at which new data is to be inserted
     */

    level = randomLevel();

    /*
     * If level is greater than current level,
     * point update to list header (Head of that level)
     */

    if (level > list->level) {
        for (i = list->level + 1; i <= level; i++) {
            update[i] = list->header;
        }
        list->level = level;
    }

    x = (sl_node *)malloc(sizeof (sl_node));
    x->data = data;
    x->forward = (sl_node **)malloc(sizeof (sl_node*) * (level + 1));

    for (i = 1; i <= level; i++) {
        x->forward[i] = update[i]->forward[i];
        update[i]->forward[i] = x;
    }

}

void
ak_sl_insert(
    skiplist        *list)
{
    int cnt = 1, data;

    while (cnt) {
        printf("\nEnter data to insert: ");
        scanf("%d", &data);
        _ak_sl_insert(list, data);
        printf("\nContinue adding? (1/0) ");
        scanf("%d", &cnt);
    }
}

void
_ak_sl_delete(
    skiplist        *list,
    int             data)
{
}

void
ak_sl_delete(
    skiplist        *list)
{
    int cnt = 1, data;

    while (cnt) {
        printf("\nEnter data to Delete: ");
        scanf("%d", &data);
        _ak_sl_delete(list, data);
        printf("\nContinue deleting? (1/0) ");
        scanf("%d", &cnt);
    }
}

sl_node *
_ak_sl_search(
    skiplist        list,
    int             data)
{
    return NULL;
}

void
ak_sl_search(
    skiplist        list)
{
    int data;
    sl_node *x;

    printf("\nEnter data to search: ");
    scanf("%d", &data);
    x = _ak_sl_search(list, data);

    if (x)
        printf("Element found:\t%d\n", data);
    else
        printf("Element not found in skip list:\t%d\n", data);

}

void
print_info()
{
    printf("*****************************************************************\n");
    printf("*  Supported operation to perform:\t\t\t\t*\n");
    printf("*\t\t\t\t\t\t\t\t*\n");
    printf("*    1. Insert Data\t\t\t\t\t\t*\n");
    printf("*    2. Delete Data\t\t\t\t\t\t*\n");
    printf("*    3. Print skip list data\t\t\t\t\t*\n");
    printf("*    4. Search in skip list\t\t\t\t\t*\n");
    printf("*****************************************************************\n");
    printf("\nSelect a option to continue: [1-4] ");
}

int
main()
{
    skiplist list;
    int data, cnt = 1, choice;

    ak_sl_init(&list);

    while (cnt) {

        print_info();

        scanf("%d", &choice);

        switch (choice) {

            case 1:
                ak_sl_insert(&list);
                break;

            case 2:
                ak_sl_delete(&list);
                break;

            case 3:
                ak_sl_print(list);
                break;

            case 4:
                ak_sl_search(list);
                break;

            default:
                printf("\nInvalid option, try again?\n");
                break;
        }

        printf("\nPress 1 to view menu, 0 to exit (1/0) ");
        scanf("%d", &cnt);
    }
    return 0;

}
