#include <cctype>
#include <iostream>
#include <fstream>
#include <string.h>
#include <string>
#include <unordered_map>

const std::string INPUT_DIR = "./inputs/day1_1.txt";
const std::string OUTPUT_DIR = "./outputs/day1_1.txt";
std::unordered_map<std::string, int> lookup = {
    {"one", 1},
    {"two", 2},
    {"three", 3},
    {"four", 4},
    {"five", 5},
    {"six", 6},
    {"seven", 7},
    {"eight", 8},
    {"nine", 9}
};


int extract_number(std::string line) {
    int num = 0;
    int last_num = -1;
    for (size_t i = 0; i < line.size(); ++i) {
        char &c = line[i];
        if (std::isdigit(c)) {
            if (last_num == -1) {
                num += (c - '0') * 10;
            }
            last_num = c - '0';
        }
        std::string subs;
        for (size_t j = 3; j < 6; ++j) {
            subs = line.substr(i, j);
            if (lookup.count(subs) != 0) {
                if (last_num == -1) {
                    num += lookup[subs] * 10;
                }
                last_num = lookup[subs];
                break;
            }
        }
    }
    return num + last_num;
}

int main(){
    std::ifstream file(INPUT_DIR);
    std::string line;

    int ans = 0;
    while (std::getline(file, line)) {
        std::cout << line << std::endl;
        int n = extract_number(line);
        std::cout << n << std::endl;
        ans += n;
    }
    file.close();

    std::ofstream output_file(OUTPUT_DIR);
    output_file << ans;
    std::cout << ans;
    return 0;
}