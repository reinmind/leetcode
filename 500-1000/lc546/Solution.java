package lc546;
import java.util.LinkedList;

public class Solution {

    public int removeBoxes(int[] boxes) {
        LinkedList<Integer> list = new LinkedList<>();
        for(int i = 0;i < boxes.length;i++){
            list.add(boxes[i]);
        }
        for(Integer i :list){
            System.out.println(i);
        }
        return 0;
    }
    public static void main(String[] args) {
        int[] a = new int[]{1, 3, 2, 2, 2, 3, 4, 3, 1};
        Solution s  = new Solution();
        s.removeBoxes(a);
    }
}
// 1,2,1,2