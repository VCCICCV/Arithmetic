#include <iostream>
#include <vector>
#include <cassert>

class Solution {
public:
    std::vector<int> twoSum(std::vector<int>& nums, int target) {
        int n = nums.size();
        for (int i = 0; i < n; ++i) {
            for (int j = i + 1; j < n; ++j) {
                if (nums[i] + nums[j] == target) {
                    return { i, j };
                }
            }
        }
        return {};
    }
};

void runTests() {
    Solution solution;

    // Test Case 1
    {
        std::vector<int> nums = { 2, 7, 11, 15 };
        int target = 9;
        std::vector<int> expected = { 0, 1 };
        std::vector<int> result = solution.twoSum(nums, target);

        assert(result == expected);
    }

    // Test Case 2
    {
        std::vector<int> nums = { 2, 7, 11, 15 };
        int target = 10;
        std::vector<int> expected = {};
        std::vector<int> result = solution.twoSum(nums, target);

        assert(result == expected);
    }

    // Add more test cases...

    std::cout << "All tests passed!" << std::endl;
}

int main() {
    runTests();

    return 0;
}