package task30;

public class Solution {

    public static void main(String[] args) {
        TreeNode s7 = new TreeNode(7);
        TreeNode s15 = new TreeNode(15);
        TreeNode s20 = new TreeNode(20, s15, s7);
        TreeNode s9 = new TreeNode(9);
        TreeNode root = new TreeNode(3, s9, s20);
        System.out.println(minDepth(root));
    }

    public static int minDepth(TreeNode root) {

        if (root == null)
            return 0;

        if ((root.left == null) && (root.right == null))
            return 1;

        int min = Integer.MAX_VALUE;

        if (root.left != null) {
            int tmp = minDepth(root.left) + 1;
            if (tmp < min)
                min = tmp;
        }

        if (root.right != null) {
            int tmp = minDepth(root.right) + 1;
            if (tmp < min)
                min = tmp;
        }

        return min;
    }
}
