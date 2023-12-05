
#include <cctype>
#include <climits>
#include <iostream>
#include <fstream>
#include <ostream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>
#include <sstream>

const std::string INPUT_DIR = "./inputs/day5_1.txt";
const std::string OUTPUT_DIR = "./outputs/day5_1.txt";
const bool IS_DAY_TWO = false;

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

long find_location(long curr_seed, const std::vector<std::string> &maps, int i) {
    if (i >= maps.size()) { return curr_seed; }
    // Find the current map
    std::string current_map = maps[i];
    std::vector<std::string> mappings = tokenize(current_map, "\n");
    long ans = LONG_MAX;
    // std::cout << curr_seed << " " << mappings[0] << " " << i << std::endl;
    bool found_match = false;
    // find correct mapping
    long matched = curr_seed;

    for (int idx = 1; idx < mappings.size(); ++idx) {
        std::vector<std::string> tokens = tokenize(mappings[idx], " ");
        long dest = stol(tokens[0]);
        long src = stol(tokens[1]);
        long range = stol(tokens[2]);
        long diff = abs(src - dest);

        if (curr_seed >= src && curr_seed < (src + range)) {
            if (dest > src) {
                matched = curr_seed + diff;
            } else {
                matched = curr_seed - diff;
            }
            break;
        }
    }
    return std::min(ans, find_location(matched, maps, i + 1));
}

long process_content(const std::string &content) {
    std::vector<std::string> sections = tokenize(content, "\n\n");
    std::vector<std::string> seeds = tokenize(sections[0], " ");
    std::vector<std::string> maps{sections.begin() + 1, sections.end()};

    long ans = LONG_MAX;
    for (int i = 1; i < seeds.size(); IS_DAY_TWO ? i += 2 : ++i) {
        if (!IS_DAY_TWO) {
            long curr_seed = stol(seeds[i]);
            ans = std::min(ans, find_location(curr_seed, maps, 0));
        } else {
            long start_seed = stol(seeds[i]);
            long end_seed = start_seed + stol(seeds[i+1]);
            for (long j = start_seed; j < end_seed; ++j) {
                ans = std::min(ans, find_location(j, maps, 0));
            }
        }
    }

    return ans;
}

int main(){
    std::ifstream file(INPUT_DIR);
    std::stringstream buffer;
    std::string content;
    std::vector<std::string> grid;
    std::string all_input;

    long ans = 0;
    buffer << file.rdbuf();
    content = buffer.str();
    ans = process_content(content);
    std::ofstream output_file(OUTPUT_DIR);
    output_file << ans;
    std::cout << "Ans = " << ans << std::endl;
    return 0;
}