package lc987;

import java.util.ArrayList;
import java.util.List;
import java.util.PriorityQueue;

/**
 * @author xiang.zhang
 * @since 2024/2/13
 */
public class Solution {

     static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode() {
        }

        TreeNode(int val) {
            this.val = val;
        }

        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }
    int [] tag = new int[3000];
    PriorityQueue<Integer> [][] m = new PriorityQueue[3000][3000];
    PriorityQueue<Integer>[] v = new PriorityQueue[3000];
    public List<List<Integer>> verticalTraversal(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();
        traverse(root,1000,0);
        for(int i = 0;i < 3000;i++){
            if(tag[i] == 0) continue;
            List<Integer> t = new ArrayList<>();
            while(v[i].peek() != null){
                final Integer j = v[i].poll();
                while(m[i][j].peek() !=null){
                    t.add(m[i][j].poll());
                }
            }
            result.add(t);
        }
        return result;
    }

    void traverse(TreeNode node,int i,int j){
        if(node == null){
            return;
        }
        if(v[i] == null){
            v[i] = new PriorityQueue<>();
        }
        v[i].add(j);
        tag[i] = 1;
        if(m[i][j] == null){
            m[i][j] = new PriorityQueue<>();
        }
        m[i][j].add(node.val);
        traverse(node.left,i-1,j+1);
        traverse(node.right,i+1,j+1);
    }

    public static void main(String[] args) {
        final TreeNode treeNode = new TreeNode();
        treeNode.val = 100;
        treeNode.left = new TreeNode(3);
        treeNode.right = new TreeNode(4);
        System.out.println(new Solution().verticalTraversal(treeNode));
    }
}
