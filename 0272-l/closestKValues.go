// 基本思路:
// 没想到特别优雅的，挫一点直接二叉搜索树转有序数组再找
// 较为清晰的一种 java 解法
class Solution {
    public List<Integer> closestKValues(TreeNode root, double target, int k) {
        Stack<TreeNode> smaller = new Stack();
        Stack<TreeNode> larger = new Stack();

        while (root != null) {
            if (root.val <= target) {
                smaller.push(root);
                root = root.right;
            } else {
                larger.push(root);
                root = root.left;
            }
        }

        List<Integer> ans = new ArrayList();
        for (; k > 0; k--) {
            double leftDif = smaller.isEmpty() ? Double.MAX_VALUE : target - smaller.peek().val;
            double rightDif = larger.isEmpty() ? Double.MAX_VALUE : larger.peek().val - target;

            if (leftDif <= rightDif) {
                TreeNode node = smaller.pop();
                ans.add(node.val);
                node = node.left;
                while (node != null) {
                    smaller.push(node);
                    node = node.right;
                }
            } else {
                TreeNode node = larger.pop();
                ans.add(node.val);
                node = node.right;
                while (node != null) {
                    larger.push(node);
                    node = node.left;
                }
            }
        }

        return ans;
    }
}
