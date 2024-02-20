package lc100217;

import java.util.HashMap;
import java.util.Map;

/**
 * @author xiang.zhang
 * @since 2024/2/18
 */
public class Solution {

    Map<Integer,Integer> primeCount = new HashMap<>();

    int max_num = -1;
    int max_count = 0;

    int[][] dir = new int[][]{{-1,0},{1,0},{0,1},{0,-1},{-1,-1},{-1,1},{1,-1},{1,1}};
    int m,n;
    int[][] mat;
    public int mostFrequentPrime(int[][] mat) {
        m = mat.length;
        n = mat[0].length;
        this.mat = mat;
        for(int i = 0;i < m;i++){
            for(int j = 0;j < n;j++){
                for(int p = 0;p < 8;p++){
                    proceed(i,j,p,0);
                }
            }
        }
        return max_num;
    }

    void proceed(int i, int j, int p, int last_sum){
        if(i < 0 || i >=m || j >=n || j <0){
            return;
        }
        int ni = dir[p][0] + i;
        int nj = dir[p][1] + j;
        int cur_sum = last_sum * 10 + mat[i][j];
        if(primeCount.get(cur_sum) != null){
            final int added = primeCount.get(cur_sum) + 1;
            primeCount.put(cur_sum,added);
            if(added > max_count){
                max_num = cur_sum;
                max_count = added;
            }
            if(added == max_count && cur_sum > max_num){
                max_num = cur_sum;
            }
        }else{
            if(isPrime(cur_sum)){
                primeCount.put(cur_sum,1);
                if(1 > max_count){
                    max_num = cur_sum;
                    max_count = 1;
                }
                if(1 == max_count && cur_sum > max_num){
                    max_num = cur_sum;
                }
            }
        }
        proceed(ni,nj,p,cur_sum);
    }

    boolean isPrime(int n){
        if(n < 10){
            return false;
        }
        if(n % 2 == 0 || n % 3 == 0){
            return false;
        }
        int c = (int)Math.sqrt(n);
        for (int i = 5; i <= c; i += 6) {
            if (n % i == 0 || n % (i + 2) == 0) {
                return false;
            }
        }
        return true;
    }
}
