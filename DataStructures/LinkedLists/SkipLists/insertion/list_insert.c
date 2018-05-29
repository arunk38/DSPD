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
sl_print(
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
sl_init(
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

int
main()
{
    skiplist list;
    int data, cnt = 1;

    sl_init(&list);

    while (cnt) {
        printf("Enter data to insert: ");
        scanf("%d", &cnt);
        printf("Continue adding? (1/0) ");
        scanf("%d", &cnt);
    }
    sl_print(list);
}
