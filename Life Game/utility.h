#include <iostream>
using namespace std;

enum Error_code {
    SUCCESS,
    FAIL,
    RANGE_ERROR,
    UNDERFLOW,
    OVERFLOW,
    FATAL,
    NOT_PRESENT,
    DUPLICATE_ERROR,
    ENTRY_INSERTED,
    ENTRY_FOUND,
    INTERNAL_ERROR
};

bool user_says_yes();
void instructions();
