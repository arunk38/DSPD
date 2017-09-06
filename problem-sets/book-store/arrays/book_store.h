#include <stdio.h>
#include <string.h>

#define MAX_INDEX	100
#define SIZE		512

/*
 * Structure to hold information of a book
 */

typedef struct {
	int	year;			/* Year of publication */
	char	name[SIZE];		/* Name of the book */
	char	author[SIZE];		/* Author's name */
	char	publisher[SIZE];	/* Publisher name */
}bs_book;

/*
 * Data structure to hold al the relavent information
 * of the book store.
 */

typedef struct {
	bs_book	book_info[MAX_INDEX];	/* All books in the store */
	int		num_books;	/* Total number of books in the store */
}bs_data_node;

bs_data_node		store_data;

/* print all the books from supplied book list */
void print_store_data();
void display_book_info(bs_book);

/* Helpful menu to user */
void print_menu();

/* Total number of books in the store */
void print_book_count();

/* Inserts the book info into given data node*/
void insert_book(bs_data_node*, bs_book);

/* Add book to store data node*/
void add_book_to_store();
