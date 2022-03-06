import java.security.KeyStore.Entry;
import java.util.List;
import java.util.Map;
class Solution {
    public String[] findRelativeRanks(int[] score) {
        String[] ranks = new String[]{"Gold Medal","Silver Medal","Bronze Medal"};
        List<Integer> scoreList = new List(score);
        List<Integer> sorted = List.sort(scoreList);
        List<String> award = new ArrayList<String>();
        for(int i = 0;i < sorted.size(); ++i){
            if(i > 2) {
                award.add(i, sorted.get(sorted.size() - i - 1));
            }
            else{
                award.add(i,ranks[i]);
            }
        }
        return award;
    }
}