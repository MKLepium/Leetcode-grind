#include<iostream>
#include<map>
#include<vector>
class RangeModule {
private:
    // Use a sorted map where the key is the start of an interval and the value is the end
    std::map<int, int> intervals;

public:
    RangeModule() {
    }

    // Adds a range [left, right)
    void addRange(int left, int right) {
        auto it = intervals.lower_bound(left);
        // If there is an interval ending before `left`, move to the previous one
        if (it != intervals.begin() && std::prev(it)->second >= left) {
            --it;
        }

        // Merge overlapping intervals
        while (it != intervals.end() && it->first <= right) {
            left = std::min(left, it->first);
            right = std::max(right, it->second);
            it = intervals.erase(it); // Remove the overlapping interval
        }

        intervals[left] = right;
    }

    // Returns true if the range [left, right) is fully covered
    bool queryRange(int left, int right) {
        auto it = intervals.upper_bound(left);
        if (it == intervals.begin()) return false;
        --it;
        return it->second >= right;
    }

    // Removes the range [left, right)
    void removeRange(int left, int right) {
        auto it = intervals.lower_bound(left);
        // If there is an interval ending before `left`, move to the previous one
        if (it != intervals.begin() && std::prev(it)->second > left) {
            --it;
        }

        // Split and remove overlapping intervals
        std::vector<std::pair<int, int>> toAdd;
        while (it != intervals.end() && it->first < right) {
            if (it->first < left) {
                toAdd.emplace_back(it->first, left); // Keep the left part of the interval
            }
            if (it->second > right) {
                toAdd.emplace_back(right, it->second); // Keep the right part of the interval
            }
            it = intervals.erase(it);
        }

        for (const auto& p : toAdd) {
            intervals[p.first] = p.second;
        }
    }
};