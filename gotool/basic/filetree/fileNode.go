package filetree

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

// Node 结构体表示一个文件或目录节点
type Node struct {
	Name         string  // 文件或目录名
	Children     []*Node // 子节点列表 遍历使用
	Depth        int     // 当前节点所在的深度，根目录为0
	ParentPath   string  // 父路径
	RelativePath string  // 相对路径
	Size         uint64  // 文件大小
	FileMode     string  // 文件类型
	IsFile       uint8   // 是否是文件 0: 文件夹 1: 文件
	ParentNode   *Node   // 父文件节点
	//MaxDepth int     // 根节点的最大深度 稍微纠结一下，是否需要放在这里 因为现在基本都是 调用 不作为参数传递给其他参数 所以注释掉，   如果外部调用涉及 多处+传参
}

func (node *Node) children(res *[]*Node) {
	// 首先添加当前节点
	*res = append(*res, node)

	// 遍历当前节点的所有子节点
	for _, child := range node.Children {
		// 对每个子节点递归地调用Children方法
		child.children(res)
	}
}

// List 遍历Node的全部节点存储到数组中
func (node *Node) List() []*Node {
	nodes := make([]*Node, 0)
	node.children(&nodes)
	return nodes
}

// AddChild 向节点添加一个子节点
func (node *Node) addChild(child *Node) {
	if node.Children == nil {
		node.Children = make([]*Node, 0)
	}
	node.Children = append(node.Children, child)
}

// BuildByOs 构建文件树，返回根节点，深度
func BuildByOs(rootPath string) (*Node, int, error) {
	//使用map 提升速度
	nodeMap, maxDepth := make(map[string]nodeAndDepth, 0), 0

	rootInfo, err := os.Stat(rootPath)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "os.Stat rootPath:%+v", rootPath)
	}
	rootNode := &Node{
		Name:         rootPath,
		IsFile:       IsFileFolder,
		RelativePath: ".",
		Size:         uint64(rootInfo.Size()),
		FileMode:     rootInfo.Mode().String(),
	}
	nodeMap["."] = nodeAndDepth{depth: 0, n: rootNode}

	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrapf(err, "filepath.Walk path:%+v, info:%+v", path, info)
		}

		// 忽略根节点本身
		if path == rootPath {
			return nil
		}

		// 从根路径创建相对路径
		relativePath, err := filepath.Rel(rootPath, path)
		if err != nil {
			return errors.Wrapf(err, "filepath.Rel rootPath:%+v, path:%+v", rootPath, path)
		}

		// parentPath的相对路径
		parentPath := filepath.Dir(relativePath)
		parentRelativePath, err := filepath.Rel(".", parentPath)
		if err != nil {
			return errors.Wrapf(err, "filepath.Rel rootPath:%+v, parentPath:%+v", rootPath, parentPath)
		}

		// 创建当前节点的 node
		currNode := &Node{
			Name:         filepath.Base(relativePath),
			RelativePath: relativePath,
			ParentPath:   parentRelativePath,
			Size:         uint64(info.Size()),
			FileMode:     info.Mode().String(),
			IsFile:       IsFileFile,
		}
		if info.IsDir() {
			currNode.IsFile = IsFileFolder
		}

		// 找到父节点
		if parentNode, ok := nodeMap[parentPath]; ok {
			// 更新当前节点的深度
			currNode.Depth = parentNode.depth + 1
			// 将当前node addChild到父节点
			currNode.ParentNode = parentNode.n
			parentNode.n.addChild(currNode)
		} else {
			return errors.Errorf("怪异的问题: 原因: 没有父节点   \n"+
				"一般来说 不会找不到父节点:在标准的 filepath.Walk 使用中，它不会先遍历子文件夹再遍历父文件夹 \n"+
				"复现使用方法_Test_BuildByOs, test_path:%+v  \n"+
				"问题点参数: path:%+v, nodeMap:%+v info:%+v  \n",
				rootPath, path, nodeMap, info)
		}

		// 更新maxDepth
		if currNode.Depth > maxDepth {
			maxDepth = currNode.Depth
			//rootNode.MaxDepth = maxDepth
		}
		// 将当前节点添加到 nodeMap
		nodeMap[relativePath] = nodeAndDepth{depth: currNode.Depth, n: currNode}

		return nil
	})

	if err != nil {
		return nil, 0, errors.Wrapf(err, "BuildByOs path:%+v", rootPath)
	}

	return rootNode, maxDepth, nil
}

// map 存储节点和 存储层 BuildByOs 方法调用此结构体
type nodeAndDepth struct {
	depth int
	n     *Node
}
