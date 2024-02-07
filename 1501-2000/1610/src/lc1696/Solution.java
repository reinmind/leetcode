package lc1696;

import java.util.PriorityQueue;

/**
 * @author xiang.zhang
 * @since 2024/2/5
 */
class Solution {


    public int maxResult(int[] nums, int k) {
        int len = nums.length;
        int[] dp = new int[len];
        PriorityQueue<Integer> maxk = new PriorityQueue<>((a, b) -> dp[b] - dp[a]);
        dp[0] = nums[0];
        maxk.offer(0);
        for (int i = 1; i < len; i++) {
            while (maxk.peek() < i - k) {
                maxk.poll();
            }
            dp[i] = dp[maxk.peek()] + nums[i];
            maxk.offer(i);
        }
        return dp[len - 1];
    }


    public static void main(String[] args) {
        final Solution solution = new Solution();
        final int i = solution.maxResult(new int[]{1, -1, -2, 4, -7, 3}, 2);
        System.out.println(i);
    }
}
