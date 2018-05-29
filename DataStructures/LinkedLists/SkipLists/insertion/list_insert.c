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

    for (i = 0; i <= level; i++) {
        sl_node *node = list.header->forward[i];

        printf("Level %d:", i);
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

    list->level = 0;

}

/*
 * create random level for skiplist node
 */

int
randomLevel()
{
    int lvl = 0;
    while (rand() < RAND_MAX / 2 && lvl < SKIPLIST_MAX_LEVEL)
        lvl++;

    return lvl;


}

/*
 * Insert into skip list at appropriate level(s)
 */

void
ak_sl_insert(
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

    /*
     * get level at which new data is to be inserted
     */

    level = randomLevel();

    if (level > list->level) {
        for (i = list->level + 1; i <= level; i++) {
            update[i] = NULL;
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

int
main()
{
    skiplist list;
    int data, cnt = 1;

    ak_sl_init(&list);

    while (cnt) {
        printf("Enter data to insert: ");
        scanf("%d", &data);
        ak_sl_insert(&list, data);
        printf("Continue adding? (1/0) ");
        scanf("%d", &cnt);
    }
    ak_sl_print(list);
}
