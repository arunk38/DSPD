/*
 * Book store program
 */

#include "book_store.h"

int main()
{
    int choice, cnt = 1;

    store_data.num_books = 0;

    while (cnt) {
	
	print_menu();
	scanf("%d", &choice);
		
	/* Select operation based on user choice */
	switch(choice) {

	    case 1:		/* Add a book to store */
		add_book_to_store();
		break;

	    case 2:		/* Delete given book(name) from store */
	        delete_book();
	        break;

	    case 3:		/* Display number of books in the store */
		print_book_count();
		break;

	    case 4:		/* Display all books in the store */
		print_store_data();
		break;

	    case 5:		/* Diaplay all books matching user query*/
	    default:
		printf("\nInvalid option, try again?\n");
		break;
	}

	printf("\nPress 1 to view menu, 0 to exit (1/0) ");
	scanf("%d", &cnt);
    }

    return 0;
}

/* Add book to store data */
void add_book_to_store()
{
    bs_book	book;
    int cnt = 1;

    while (cnt) {
	printf("\nEnter following book details:\n");
	
	/* Get details of book from user */
	printf("Name:\t\t\t");
	scanf("%s", book.name);
	printf("Author:\t\t\t");
	scanf("%s", book.author);
	printf("Publisher:\t\t");
	scanf("%s", book.publisher);
        printf("Year of publication:\t");
        scanf("%d", &book.year);

        insert_book(&store_data, book);

	/* Ask for multiple books */
	printf("\nSuccessfully added book to store, continue adding? (1/0) ");
	scanf("%d", &cnt);
    }
}

/* Insert book into given data node */
void insert_book(bs_data_node *data_node, bs_book book)
{
    int indx = data_node->num_books;

    strcpy(data_node->book_info[indx].name, book.name);
    strcpy(data_node->book_info[indx].author, book.author);
    strcpy(data_node->book_info[indx].publisher, book.publisher);

    data_node->book_info[indx].year = book.year;

    data_node->num_books++;
}

/* Displays book store menu with all supported operations */
void print_menu()
{
    printf("*****************************************************************\n");
    printf("*\t\tWELCOME TO BOOK STORE DATA BASE\t\t\t*\n");
    printf("*\t\t\t\t\t\t\t\t*\n");
    printf("*  Supported operation to perform:\t\t\t\t*\n");
    printf("*\t\t\t\t\t\t\t\t*\n");
    printf("*    1. Add a book to store\t\t\t\t\t*\n");
    printf("*    2. Remove a book from store\t\t\t\t*\n");
    printf("*    3. Total number of books in store\t\t\t\t*\n");
    printf("*    4. Diaplay all books in store\t\t\t\t*\n");
    printf("*    5. Search store\t\t\t\t\t\t*\n");
    printf("*****************************************************************\n");
    printf("\nSelect a option to continue: [1-5] ");
}

/* Total number of books in store */
void print_book_count()
{
    printf("\nTotal number of books in store: %d\n", store_data.num_books);
}

void print_book(bs_book book)
{
    printf("%s\t%-12s\t%-20s\t\t%d\n", book.name,
		book.author, book.publisher, book.year);
}

/* Print information about the store */
void print_store_data()
{
    int indx = 0;

    print_book_count();

    printf("\nBook name\tAuthor\t\tPublisher\t\t\tPublication year\n");
    while (indx < store_data.num_books) {
	print_book(store_data.book_info[indx]);
	indx++;
    }
}

void delete_book()
{
    char name[SIZE];
    int indx = 0, max_indx = store_data.num_books;

    printf("\nEnter the name of the book to delete:  ");
    scanf("%s", name);

    while (indx < max_indx) {

        /* Search for given book */
        if (strcmp(store_data.book_info[indx].name, name) == 0) {

            /* If found copy the last book in store to its place and decrement books count */
            strcpy(store_data.book_info[indx].name, store_data.book_info[max_indx-1].name);
            strcpy(store_data.book_info[indx].author, store_data.book_info[max_indx-1].author);
            strcpy(store_data.book_info[indx].publisher, store_data.book_info[max_indx-1].publisher);

            store_data.book_info[indx].year = store_data.book_info[max_indx-1].year;

            store_data.num_books--;

            printf("\nSucessfully deleted book from store:\n");
            print_book(store_data.book_info[max_indx-1]);
            return;
        }
        
        indx++;
    }

    printf("\nBook with given name not found: %s\n", name);
}
