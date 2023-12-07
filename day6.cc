
#include <cctype>
#include <climits>
#include <cmath>
#include <iostream>
#include <fstream>
#include <ostream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>
#include <sstream>

const std::string INPUT_DIR = "./tests/day6_1.txt";
const std::string OUTPUT_DIR = "./outputs/day6_1.txt";
const bool IS_DAY_TWO = true;


std::vector<std::string> tokenize(const std::string &line, const std::string &delim) {
    std::vector<std::string> tokens;
    size_t start = 0, end = delim.size();
    std::string token;
    while ((end = line.find(delim, start)) != std::string::npos) {
        size_t prev_end = end;
        while (end < line.size() - 1 && (line[end+1] == delim[0])) {
            ++end;
        }
        token = line.substr(start, prev_end - start);
        start = end + delim.size();
        tokens.push_back(token);
    }
    tokens.push_back(line.substr(start));
    return tokens;
}

long get_num_sols(long T, long D) {
    long root1 = std::ceil(((T + std::sqrt(T*T - 4 * D)) / 2)) - 1;
    long root2 = std::floor((T - std::sqrt(T*T - 4 * D)) / 2) + 1;
    return root1 - root2 + 1;
}


long solve(const std::vector<long> &times, const std::vector<long> &dist) {
    long ans = 1;
    for (int i = 0; i < times.size(); ++i) {
        long T = times[i];
        long D = dist[i];
        ans *= get_num_sols(T, D);
    } 
    return ans;
}

int main(){
    std::ifstream file(INPUT_DIR);
    std::string line;
    std::vector<long> times;
    std::vector<long> distances;
    long ans = 0;

    int idx = 1;
    while (std::getline(file, line)) {
        std::vector<std::string> tokens = tokenize(line, " ");
        for (int i = 1; i < tokens.size(); ++i) {
            if (idx == 1) {
                times.push_back(stol(tokens[i]));
            } else {
                distances.push_back(stol(tokens[i]));
            }
        } 
        ++idx;
    }

    ans = solve(times, distances);

    std::ofstream output_file(OUTPUT_DIR);
    output_file << ans;
    std::cout << "Ans = " << ans << std::endl;
    return 0;
}