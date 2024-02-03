package ms16d17;

public class Solution {

    public int maxSubArray(int[] nums) {
        int max = nums[0];
        int minv = 0;
        if(nums[0] < minv){
            minv = nums[0];
        }
        for(int i =1;i< nums.length;i++){
            if(nums[i] > max){
                max = nums[i];
            }
            nums[i] = nums[i-1] + nums[i];
            if(nums[i] < minv){
                minv = nums[i];
                continue;
            }
            int t = nums[i] - minv;
            if(t > max){
                max = t;
            }
        }
        return max;
    }
}
