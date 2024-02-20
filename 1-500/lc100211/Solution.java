package lc100211;

import java.util.*;

/**
 * @author xiang.zhang
 * @since 2024/2/17
 */
public class Solution {
    public String lastNonEmptyString(String s) {
        final int[] cc = new int[27];
        final char[] ca = s.toCharArray();
        int max = 0;
        List<Character> l = new ArrayList<>();
        for(char c:ca){
            int count = ++cc[c-'a'];
            if(count > max){
                max = count;
                l.clear();
                l.add(c);
                continue;
            }
            if(count == max){
                l.add(c);
            }
        }
        final StringBuilder sb = new StringBuilder();
        for(Character c:l){
            sb.append(c);
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        System.out.println(new Solution().lastNonEmptyString("aabcbbca"));
        System.out.println(new Solution().lastNonEmptyString("abcdzzz"));
    }
}
