package lc3036;

import java.util.*;

/**
 * @author xiang.zhang
 * @since 2024/2/12
 */
public class Solution {
    String a = "ababeabcabd";
    String b = "eabcabd";


    public int countMatchingSubarrays(int[] nums, int[] pattern) {
        final int len = nums.length;
        int plen = pattern.length;
        int count = 0;
        int[] ori = new int[len];
        for(int i = 1;i< len;i++){
            int subtract = nums[i] - nums[i - 1];
            ori[i] = subtract == 0 ? 0 : subtract / Math.abs(subtract);
        }
        int[] next = new int[plen];
        next[0] = 0;
        for(int i = 1;i < plen;i++){
            if(pattern[i] == pattern[next[i-1]]){
                next[i] = next[i-1] + 1;
            }else{
                next[i] = 0;
            }
        }
        System.out.println(Arrays.toString(ori));
        System.out.println(Arrays.toString(pattern));
        System.out.println(Arrays.toString(next));
        int j = 0;
        for(int i = 1;i < len;){
            if(plen == 1){
                if(pattern[j] == ori[i]){
                    count++;
                }
                i++;
                continue;
            }
            if(pattern[j] == ori[i]){
                if(j == plen - 1){
                    // 假设最后一个不相等，比较前一个
                    // 如果 pattern 只有一个元素，有可能直接匹配
                    j = next[j-1];
                    count++;
                }else{
                    j++;
                    i++;
                }
            }else {
                if(j > 0){
                    j = next[j-1];
                }else{
                    i++;
                }
            }
        }
        return count;
    }

    public static void main(String[] args) {
        final Solution solution = new Solution();
        System.out.println(solution.countMatchingSubarrays(new int[]{132030805,226565039,226565039,234134481,362977930,362977930,565625669,651711104,713365290,651711104,713365290,713365290,718731626},new int[]{1,0,1,1}));
        System.out.println(solution.countMatchingSubarrays(new int[]{872500231,190002870},new int[]{-1}));
    }
}
