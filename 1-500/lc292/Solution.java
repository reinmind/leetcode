package lc292;

import java.util.Arrays;

public class Solution {


    public boolean canWinNim(int n) {
        boolean[] result = new boolean[4] ;
        result[0] = true;
        result[1] = true;
        result[2] = true;
        result[3] = true;

        for(int t = 4;t<= n;t++){
            // 如果场上有 4 个石头,只要有一种解法成功就可以
            result[t%4] = !result[(t-1)%4] || !result[(t-2)%4] || !result[(t-3)%4];
        }
        return result[n%4];
    }

    public boolean canWin2(int n){
        boolean[] result = new boolean[4] ;
        result[0] = true;
        result[1] = true;
        result[2] = true;
        result[3] = false;
        return result[(n-1) % 4];
    }

    public static void main(String[] args) {
        final Solution solution = new Solution();
        final boolean win = solution.canWinNim(1348820612);
        final boolean b = solution.canWin2(1348820612);
        System.out.println("canWinNim: " + win);
        System.out.println("canWin2: " + b);
    }
}
