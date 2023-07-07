package DS

type Avl struct {
	root *node
}

func (avl *Avl) Insert(val int) {
	avl.root = insert(avl.root, val)
}

// 定义节点
type node struct {
	val    int   // 值
	height int   // 高度
	left   *node // 左儿子
	right  *node // 右儿子
}

// 四种旋转情况：左旋转，右旋转，左旋转再右旋转，右旋转再左旋转
func rotateLeft(root *node) *node {
	temp := root.right
	root.right = temp.left
	temp.left = root
	// 更新高度，谁的子节点变了，更新谁的高度
	root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	temp.height = max(getHeight(temp.left), getHeight(temp.right)) + 1
	return temp
}

func rotateRight(root *node) *node {
	temp := root.left
	root.left = temp.right
	temp.right = root
	root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	temp.height = max(getHeight(temp.left), getHeight(temp.right)) + 1
	return temp
}

func rotateLeftRight(root *node) *node {
	root.left = rotateLeft(root.left) // 先把左节点左转，再整体右转
	return rotateRight(root)
}

func rotateRightLeft(root *node) *node {
	root.right = rotateRight(root.right)
	return rotateLeft(root)
}

func insert(root *node, val int) *node {
	// 创建节点
	if root == nil {
		root = &node{
			val:    val,
			left:   nil,
			right:  nil,
			height: 1,
		}
	} else {
		if val < root.val { // 往左插
			root.left = insert(root.left, val)
			root.height = max(getHeight(root.left), getHeight(root.right)) + 1
			// 开始处理旋转, 这一步体现的就是找到失衡的点
			if getHeight(root.left)-getHeight(root.right) == 2 {
				if root.left.val > val { // 左左,右旋
					root = rotateRight(root)
				} else { // 左右,左旋再右旋
					root = rotateLeftRight(root)
				}
			}
		} else { // 往右边插
			root.right = insert(root.right, val)
			root.height = max(getHeight(root.left), getHeight(root.right)) + 1
			if getHeight(root.right)-getHeight(root.left) == 2 {
				if root.right.val <= val { // 右右, 左旋
					root = rotateLeft(root)
				} else {
					root = rotateRightLeft(root)
				}
			}
		}
	}
	return root
}

func getHeight(root *node) int {
	if root == nil {
		return 0
	}
	return root.height
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
