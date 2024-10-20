
#include <vector>
#include <bits/algorithmfwd.h>

using std::vector;
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        nums.erase(std::unique(nums.begin(), nums.end()), nums.end());
        return nums.size();
    }
};