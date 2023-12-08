
#include <cctype>
#include <iostream>
#include <fstream>
#include <string.h>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>

const std::string INPUT_DIR = "./inputs/day7_1.txt";
const std::string OUTPUT_DIR = "./outputs/day7_1.txt";
std::unordered_map<char, int> ranks = {
    {'A', 1},
    {'K', 2},
    {'Q', 3},
    {'J', 4},
    {'T', 5},
    {'9', 6},
    {'8', 7},
    {'7', 8},
    {'6', 9},
    {'5', 10},
    {'4', 11},
    {'3', 12},
    {'2', 13}
};
enum class HAND {
    FIVE,
    FOUR,
    FULL,
    THREE,
    TWOPAIR,
    ONEPAIR,
    HIGH,
    UNDEF
};

std::string hand_to_str(HAND hand) {
    switch (hand) {
        case HAND::FIVE: return "FIVE";
        case HAND::FOUR: return "FOUR";
        case HAND::FULL: return "FULL";
        case HAND::THREE: return "THREE";
        case HAND::TWOPAIR: return "TWOPAIR";
        case HAND::ONEPAIR: return "ONEPAIR";
        case HAND::HIGH: return "HIGH";
        default: return "UNDEF";
    }
}

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

HAND infer_hand(std::string hand) {
    std::unordered_map<char, int> groups;
    for (const char &c : hand) {
        ++groups[c];
    }
    int num_groups = groups.size();
    switch (num_groups) {
        case 5: return HAND::HIGH;
        case 4: return HAND::ONEPAIR;
        // (3,1,1) , (2, 2, 1)
        case 3: {
            for (auto it : groups) {
                if (it.second == 3) {
                    return HAND::THREE;
                }
                if (it.second == 2) {
                    return HAND::TWOPAIR;
                }
            }
        }
        // (3,2), (4,1)
        case 2: {
            for (auto it : groups) {
                if (it.second == 4) {
                    return HAND::FOUR;
                }
                if (it.second == 3) {
                    return HAND::FULL;
                }
            }
        }
        case 1: return HAND::FIVE;
        default: return HAND::UNDEF;
    }
}

long get_all_ranks(std::unordered_map<HAND, std::vector<std::string>> &grouped_hands, std::unordered_map<std::string, int> &hand_bids) {
    for (auto &[k, v] : grouped_hands) {
        // std::cout << "\nBEFORE SORTING : " << hand_to_str(k) << '\n';
        // for (const auto &s : v) {
        //     std::cout << s << " ";
        // }
        sort(v.begin(), v.end(), [](const auto &l, const auto &r){
            for (int i = 0; i < 5; ++i) {
                if (ranks[l[i]] == ranks[r[i]]) continue;
                if (ranks[l[i]] < ranks[r[i]]) return false;
                return true;
            }
            return false;
        });
        // std::cout << "\nAFTER SORTING : " << hand_to_str(k) << '\n';
        // for (const auto &s : v) {
        //     std::cout << s << " ";
        // }
    }

    std::vector<HAND> order = {HAND::HIGH, HAND::ONEPAIR, HAND::TWOPAIR, HAND::THREE, HAND::FULL, HAND::FOUR, HAND::FIVE};
    int rank = 1;
    long ans = 0;
    for (const auto &o : order) {
        // std::cout << hand_to_str(o) << " : ";
        for (const auto &s : grouped_hands[o]) {
            // std::cout << s << " " << rank << " ";
            ans += (rank * hand_bids[s]);
            ++rank;
        }
        // std::cout << '\n';
    }
    return ans;
}

int main(){
    std::ifstream file(INPUT_DIR);
    std::string line;
    long ans = 0, ans_2 = 0;
    std::unordered_map<HAND, std::vector<std::string>> grouped_hands;
    std::unordered_map<std::string, int> hand_bids;

    while (std::getline(file, line)) {
        std::vector<std::string> tokens = tokenize(line, " ");
        hand_bids[tokens[0]] = stoi(tokens[1]);
        HAND hand = infer_hand(tokens[0]);
        grouped_hands[hand].push_back(tokens[0]);
    }
    file.close();

    ans = get_all_ranks(grouped_hands, hand_bids);

    std::ofstream output_file(OUTPUT_DIR);
    output_file << ans;
    std::cout << "Ans Part 1 = " << ans << std::endl;
    std::cout << "Ans Part 2 = " << ans_2 << std::endl;
    return 0;
}