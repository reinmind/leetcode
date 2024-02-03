package lc104;


public class Solution {

    public int maxDepth(TreeNode root) {
        return depth(root);
    }

    int depth(TreeNode treeNode){
        if(treeNode == null){
            return 0;
        }
        return Math.max(depth(treeNode.left),depth(treeNode.right)) + 1;
    }
}


class TreeNode {
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
