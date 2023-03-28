package task31;

public class Solution {

    public static void main(String[] args) {
        TreeNode s3 = new TreeNode(3);
        TreeNode s2 = new TreeNode(2);
        /*TreeNode s7 = new TreeNode(7);
        TreeNode s4 = new TreeNode(4, null, s1);
        TreeNode s13 = new TreeNode(13);
        TreeNode s11 = new TreeNode(11, s7, s2);
        TreeNode s8 = new TreeNode(8, s13, s4);
        TreeNode s42 = new TreeNode(4, s11, null);*/
        TreeNode root = new TreeNode(1, s2, null);
        System.out.println(hasPathSum(root, 1));
    }

    public static boolean hasPathSum(TreeNode root, int targetSum) {

        if (root == null)
            return false;

        return helper(root, targetSum, 0);
    }

    public static boolean helper(TreeNode treeNode, int targetSum, int previousSum) {

        if ((treeNode.left == null) && (treeNode.right == null)) {
            if ((treeNode.val + previousSum) == targetSum) {
                return true;
            }
        }

        previousSum += treeNode.val;

        if ((treeNode.left != null) && (helper(treeNode.left, targetSum, previousSum)))
            return true;

        if ((treeNode.right != null) && (helper(treeNode.right, targetSum, previousSum)))
            return true;

        return false;
    }
}
