public class MinInsertions {
    public int minInsertions(String s) {
        int n = s.length();
        if (n <= 1) return 0;

        // 创建一个(n+1) x (n+1)的二维数组，dp[i][j]表示s.substring(0, i)转换为回文串需要的最少插入次数，使得它与s.substring(0, j)相同
        int[][] dp = new int[n + 1][n + 1];

        // 初始化对角线上的元素，因为单个字符本身就是回文串
        for (int i = 0; i <= n; i++) {
            dp[i][i] = 0;
        }

        // 动态规划填充矩阵
        for (int len = 2; len <= n; len++) {
            for (int start = 0; start <= n - len; start++) {
                int end = start + len - 1;
                if (s.charAt(start) == s.charAt(end)) {
                    dp[start][end] = dp[start + 1][end - 1];
                } else {
                    dp[start][end] = Math.min(dp[start + 1][end], dp[start][end - 1]) + 1;
                }
            }
        }

        // 最终答案位于dp[0][n-1]
        return dp[0][n - 1];
    }
}
