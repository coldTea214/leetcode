func closestValue(root *TreeNode, target float64) int {
	closest, cur := root, root

	for cur != nil {
		if target == float64(cur.Val) {
			return cur.Val
		}

		if math.Abs(target-float64(cur.Val)) < math.Abs(target-float64(closest.Val)) {
			closest = cur
		}

		if target < float64(cur.Val) {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	return closest.Val
}
