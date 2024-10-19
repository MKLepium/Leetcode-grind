#include <vector>

using std::vector;

class Solution {
public:
    vector<vector<int>> findMissingRanges(vector<int>& nums, int lower, int upper) {
        vector<vector<int>> missingRanges{};
        if (nums.empty()) {
            missingRanges.push_back({lower, upper});
            return missingRanges;
        }

        if (nums.front() > lower) {
            missingRanges.push_back({lower, nums.front() - 1});
        }

        // Every other case within the ranges
        for (int i = 0; i < nums.size() - 1; i++)
        {
            // Not range
            if(nums[i] + 1 == nums[i + 1]) {
                continue;
            } 
            // Range between larger than 1
            missingRanges.push_back({nums[i] + 1, nums[i + 1] - 1});
        }

        if (nums.back() < upper) {
            missingRanges.push_back({nums.back() + 1, upper});
        }
        
        return missingRanges;
    }
};