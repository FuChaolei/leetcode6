/*
 * @lc app=leetcode.cn id=718 lang=cpp
 *
 * [718] 最长重复子数组
 *
 * https://leetcode.cn/problems/maximum-length-of-repeated-subarray/description/
 *
 * algorithms
 * Medium (56.94%)
 * Likes:    911
 * Dislikes: 0
 * Total Accepted:    184.8K
 * Total Submissions: 324.4K
 * Testcase Example:  '[1,2,3,2,1]\n[3,2,1,4,7]'
 *
 * 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
 * 输出：3
 * 解释：长度最长的公共子数组是 [3,2,1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length, nums2.length <= 1000
 * 0 <= nums1[i], nums2[i] <= 100
 *
 *
 */

// @lc code=start
class Solution
{
public:
    // int findLength(vector<int> &nums1, vector<int> &nums2)
    // {
    //     int m = nums1.size(), n = nums2.size();
    //     vector<vector<int>> dp(m + 1, vector<int>(n + 1));
    //     int res = 0;
    //     for (int i = 1; i <= m; i++)
    //     {
    //         for (int j = 1; j <= n; j++)
    //         {
    //             if (nums1[i - 1] == nums2[j - 1])
    //             {
    //                 dp[i][j] = dp[i - 1][j - 1] + 1;
    //                 res = max(res, dp[i][j]);
    //             }
    //         }
    //     }
    //     return res;
    // }
    int findLength(vector<int> &nums1, vector<int> &nums2)
    {
        int m = nums1.size(), n = nums2.size();
        function<int(int, int, int)> maxl = [&](int addA, int addB, int len)
        {
            int ret = 0;
            int k = 0;
            for (int i = 0; i < len; i++)
            {
                if (nums1[addA + i] == nums2[addB + i])
                {
                    k++;
                }
                else
                {
                    k = 0;
                }
                ret = max(k, ret);
            }
            return ret;
        };
        int res = 0;
        for (int i = 0; i < m; i++)
        {
            int len = min(m - i, n);
            res = max(res, maxl(i, 0, len));
        }
        for (int i = 0; i < n; i++)
        {
            int len = min(n - i, m);
            res = max(res, maxl(0, i, len));
        }
        return res;
    }
};
// @lc code=end
