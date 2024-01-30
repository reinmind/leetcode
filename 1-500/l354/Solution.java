package l354;

import java.util.Arrays;

class Solution {
    // [[46,89],[46,46],[50,69],[50,53],[52,68],[72,45],[77,81]]
    public int maxEnvelopes(int[][] envelopes) {
        Arrays.sort(envelopes, (a, b) -> {
            // 首先按照第一列升序排序
            int compareResult = Integer.compare(a[0], b[0]);

            // 如果第一列相等，则按照第二列降序排序
            if (compareResult == 0) {
                compareResult = Integer.compare(b[1], a[1]);
            }
            return compareResult;
        });
        int[] tails = new int[envelopes.length];
        tails[0] = envelopes[0][1];
        int tlen = 1;

        for(int i = 0;i < envelopes.length;i++){
            int num = envelopes[i][1];
            int left = 0,right = tlen;
            while(left < right){
                int mid = (left + right) / 2;
                if(tails[mid] < num){
                    left = mid + 1;
                }else{
                    right = mid;
                }
            }
            if(num < tails[right]){
                tails[right] = num;
            }
            if(left == tlen){
                tlen++;
                tails[left] = num;
            }
        }

        return tlen;
    }
}
