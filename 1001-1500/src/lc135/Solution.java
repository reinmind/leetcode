package lc135;

class Solution {
    int[] g;
    int[] r;

    public int candy(int[] ratings) {
        g = new int[ratings.length];
        r = ratings;
        if(ratings.length == 1){
            return 1;
        }
        if (r[1] >= r[0]) {
            g[0] = 1;
        }
        if (r[1] < r[0]) {
            g[0] = gains(1) + 1;
        }
        int sum = 0;
        for (int i = 0; i < ratings.length; i++) {
            sum += gains(i);
        }
        return sum;
    }

    int gains(int p) {
        if (g[p] != 0) {
            return g[p];
        }
        if (p == r.length - 1) {
            if (r[p] == r[p - 1]) {
                g[p] = 1;
            }
            if (r[p] < r[p - 1]) {
                g[p] = 1;
            }
            if (r[p] > r[p - 1]) {
                g[p] = gains(p - 1) + 1;
            }
            return g[p];
        }
        if (r[p] <= r[p + 1] && r[p] <= r[p - 1]) {
            g[p] = 1;
        }
        if (r[p] > r[p + 1] && r[p] > r[p - 1]) {
            g[p] = Math.max(gains(p + 1) + 1,gains(p-1)+1);
        }
        if(r[p] > r[p+1] && r[p] <= r[p-1]){
            g[p] = gains(p+1) + 1;
        }
        if(r[p] > r[p-1] && r[p] <= r[p+1]){
            g[p] = gains(p-1) + 1;
        }
        return g[p];
    }

    public static void main(String[] args) {
        int[] rating = new int[]{1, 0, 2, 3, 4, 5, 4, 3, 2};
        final Solution solution = new Solution();
        final int candy = solution.candy(rating);
        System.out.println(candy);
    }
}
