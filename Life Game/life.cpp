/*
 * Life game
 */

#include "utility.h"
#include "life.h"

/*
 * Driver function for life game
 */

int main()
{
    Life configuration;
    instructions();
    configuration.initialize();
    configuration.print();
    cout << "Continue viewing new generations?" << endl;

    while (user_says_yes()) {
        configuration.update();
        configuration.print();
        cout << "Continue viewing new generations?" << endl;
    }
 }

 void Life::initialize() {}
 void Life::print() {}
 void Life::update() {}

 int
 Life:neighbor_count(
                     int            row,
                     int            col)
{
    int i, j;
    int count = 0;

    for (i = row − 1; i <= row + 1; i++)
        for (j = col − 1; j <= col - 1; j++)
            count += grid[i][j];  // Increase the count if neighbor is alive.

    count −= grid[row][col]; // Reduce count, since cell is not its own neighbor.
    return count;
}
