
#include <cctype>
#include <iostream>
#include <fstream>
#include <string.h>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>

const std::string INPUT_DIR = "./inputs/day4_1.txt";
const std::string OUTPUT_DIR = "./outputs/day4_1.txt";
std::unordered_map<int, int> num_copies_won;

std::vector<std::string> tokenize(const std::string &line, const std::string &delim) {
    std::vector<std::string> tokens;
    size_t start = 0, end = delim.size();
    std::string token;
    while ((end = line.find(delim, start)) != std::string::npos) {
        token = line.substr(start, end - start);
        start = end + delim.size();
        tokens.push_back(token);
    }
    tokens.push_back(line.substr(start));
    return tokens;
}


int get_num_matches(const std::string &line) {
    std::string game = tokenize(line, ":")[1];
    std::vector<std::string> turns = tokenize(game, "|");
    std::vector<std::string> reference = tokenize(turns[0], " ");
    std::vector<std::string> hand = tokenize(turns[1], " ");
    std::unordered_set<int> true_vals;
    for (const auto x : reference) {
        if (!x.empty()) {
            true_vals.insert(stoi(x));
        }
    }
    int num_matches = 0;
    for (const auto x : hand) {
        if (!x.empty() && (true_vals.count(stoi(x)) != 0)) {
            num_matches += 1;
        }
    }
    return num_matches;
}

int get_points(const std::string &line) {
    int num_matches = get_num_matches(line);
    return num_matches == 0 ? 0 : 1 << (num_matches - 1);
}
int main(){
    std::ifstream file(INPUT_DIR);
    std::string line;
    std::vector<std::string> grid;

    int ans = 0, ans_2 = 0, line_num = 1;
    while (std::getline(file, line)) {
        ans += get_points(line);
        int num_matches = get_num_matches(line);
        ++num_copies_won[line_num];
        if (num_matches != 0) {
            for (int i = 1; i <= num_matches; ++i) {
                num_copies_won[line_num + i] += num_copies_won[line_num];
            }
        }
        ++line_num;
    }
    file.close();

    for (auto it : num_copies_won) {
        ans_2 += it.second;
    }

    std::ofstream output_file(OUTPUT_DIR);
    output_file << ans;
    std::cout << "Ans Part 1 = " << ans << std::endl;
    std::cout << "Ans Part 2 = " << ans_2 << std::endl;
    return 0;
}