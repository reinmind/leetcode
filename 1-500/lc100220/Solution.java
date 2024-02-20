package lc100220;

/**
 * @author xiang.zhang
 * @since 2024/2/17
 */
public class Solution {
    int[] nums;
    int len ;
    public int maxOperations(int[] nums) {
        this.nums = nums;
        len = nums.length;
        int max = 0;
        int a = dfs(0,len-2,nums[len-1]+nums[len-2]);
        int b = dfs(2,len,nums[0]+nums[1]);
        int c = dfs(1,len-1,nums[0]+nums[len-1]);
        max = Math.max(a,b);
        max = Math.max(max,c);
        return max+1;
    }

    int dfs(int a,int b,int sum){
        if(b - a < 2 || a >= b || a > len-1 || b < 2 ){
            return 0;
        }
        int d1=-1,d2=-1,d3=-1,m=-1;
        if(nums[a] + nums[a+1] == sum) d1 = dfs(a+2,b,sum);
        if(nums[a] + nums[b-1] == sum) d2 = dfs(a+1,b-1,sum);
        if(nums[b-1] + nums[b-2] == sum) d3 = dfs(a,b-2,sum);
        m = Math.max(d1,d2);
        m = Math.max(d3,m);
        return m+1;
    }

    public static void main(String[] args) {
        System.out.println(new Solution().maxOperations(new int[]{3,4,3,2,1,2,3,4}));
        System.out.println(new Solution().maxOperations(new int[]{1,2,3,2,2,1,2,3}));
        System.out.println(new Solution().maxOperations(new int[]{3,2,1,2,3,4}));
        System.out.println(new Solution().maxOperations(new int[]{3,2,6,1,4}));
    }
}
