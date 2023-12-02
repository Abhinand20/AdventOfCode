#include <cctype>
#include <iostream>
#include <fstream>
#include <string.h>
#include <string>
#include <unordered_map>
#include <vector>

const std::string INPUT_DIR = "./inputs/day2_1.txt";
// const std::string INPUT_DIR = "./tests/day2_2.txt";
const std::string OUTPUT_DIR = "./outputs/day2_1.txt";
const bool IS_PART_TWO = true;

std::unordered_map<std::string, int> max_vals = {
    {"blue", 14},
    {"red", 12},
    {"green", 13}
};

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

bool is_turn_valid(const std::string &turn) {
    std::unordered_map<std::string, int> values;
    std::vector<std::string> tokens = tokenize(turn, ",");
    for (const std::string &token : tokens) {
        std::vector<std::string> num_cubes = tokenize(token.substr(1), " ");
        assert(num_cubes.size() == 2);
        values[num_cubes[1]] = stoi(num_cubes[0]);
    }
    for (auto &[key, val] : values) {
        if (max_vals[key] < val) {
            return false;
        }
    }
    return true;
}

bool is_game_valid(const std::string &line) {
    std::string cubes_shown = tokenize(line, ":")[1];
    std::vector<std::string> turns = tokenize(cubes_shown, ";");
    for (auto &turn : turns) {
        if (!is_turn_valid(turn)) {
            return false;
        }
    }
    return true;
}

int calculate_power(const std::string &line) {
    std::string cubes_shown = tokenize(line, ":")[1];
    std::vector<std::string> turns = tokenize(cubes_shown, ";");
    int power = 1;
    std::unordered_map<std::string, int> values;
    for (auto &turn : turns) {
        std::vector<std::string> tokens = tokenize(turn, ",");
        for (const std::string &token : tokens) {
            std::vector<std::string> num_cubes = tokenize(token.substr(1), " ");
            assert(num_cubes.size() == 2);
            values[num_cubes[1]] = std::max(values[num_cubes[1]],stoi(num_cubes[0]));
        } 
    }
    for (auto it : values) {
        power *= it.second;
    }
    return power;
}

int main(){
    std::ifstream file(INPUT_DIR);
    std::string line;

    int ans = 0;
    int game_num = 1;
    while (std::getline(file, line)) {
        if (!IS_PART_TWO) {
            if (is_game_valid(line)) {
                ans += game_num;
            }
            ++game_num;
        } else {
            ans += calculate_power(line);
        }
    }
    file.close();

    std::ofstream output_file(OUTPUT_DIR);
    output_file << ans;
    std::cout << ans;
    return 0;
}