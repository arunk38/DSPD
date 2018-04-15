#define MAX_ROW     20
#define MAX_COL        60

class Life {
public:
    void initialize();
    void print();
    void update();

private:
    int grid [MAX_ROW + 2][MAX_COL + 2];

    int neighbor_count(int orw, int col);
};
