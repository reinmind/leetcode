package lc100221;

/**
 * @author xiang.zhang
 * @since 2024/2/17
 */
class Solution {
    public int maxOperations(int[] nums) {
        int sum = nums[0] + nums[1];
        int count = 0;
        for(int i = 0;i < nums.length - 1; i+=2){
            if(nums[i] + nums[i+1] == sum){
                count++;
            }else{
                break;
            }
        }
        return count;
    }
}
