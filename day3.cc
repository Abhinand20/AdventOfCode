#include <cctype>
#include <iostream>
#include <fstream>
#include <string.h>
#include <string>
#include <unordered_map>
#include <vector>

const std::string INPUT_DIR = "./inputs/day3_1.txt";
const std::string OUTPUT_DIR = "./outputs/day3_1.txt";
const struct {
    int x;
    int y;
} DIRECTIONS[] = {
    {0,1}, 
    {1, 0},
    {0, -1}, 
    {-1, 0},
    {-1, -1}, 
    {-1, 1},
    {1, -1},
    {1,1}
    };


bool check_bounds(const int &x, const int &y, const int &r, const int&c) {
    return x >= 0 && x < r && y >= 0 && y < c;
}

bool can_pick(const int &x, const int &y, const std::vector<std::string> &grid) {
    return grid[x][y] != '.' && !std::isdigit(grid[x][y]);
}

bool is_valid(const int &i, const int &start, const int &end, const std::vector<std::string> &grid) {
    // Validate each position
    const int r = grid.size();
    const int c = grid[0].size();
    for (int j = start; j <= end; ++j) {
        for (int d = 0; d < 8; ++d) {
            int next_x = i + DIRECTIONS[d].x;
            int next_y = j + DIRECTIONS[d].y;
            if (check_bounds(next_x, next_y, r, c) && can_pick(next_x, next_y, grid)) {
                return true;
            }
        }
    }
    return false;
}

int get_sum_of_parts(const std::vector<std::string> &grid) {
    const int r = grid.size();
    const int c = grid[0].size();
    int ans = 0;
    // Extract numbers in each line and validate
    for (int i = 0; i < r; ++i) {
        std::string token;
        for (int j = 0; j < c; ++j) {   
            if (!token.empty() && !std::isdigit(grid[i][j])) {
                int start = j - token.size();
                if (is_valid(i, start, j - 1, grid)) {
                    ans += stoi(token);
                }
                token.clear();
            }
            if (std::isdigit(grid[i][j])) {
                token.push_back(grid[i][j]);
            }
        }
        if (!token.empty()) {
            if (is_valid(i, c - token.size(), c - 1, grid)) {
                ans += stoi(token);
            }
        }
    }
    return ans;
}

int main(){
    std::ifstream file(INPUT_DIR);
    std::string line;
    std::vector<std::string> grid;

    int ans = 0;
    while (std::getline(file, line)) {
        grid.push_back(line);
    }
    file.close();

    ans = get_sum_of_parts(grid);

    std::ofstream output_file(OUTPUT_DIR);
    output_file << ans;
    std::cout << "Ans = " << ans << std::endl;
    return 0;
}