package filetree

import (
	"github.com/pkg/errors"
	"path/filepath"
	"testing"
)

func Test_BuildByOs(t *testing.T) {
	rootPath := "/Users/3phone/Downloads/test"
	rootNode, maxDepth, err := BuildByOs(rootPath)
	if err != nil {
		t.Errorf("BuildByOs(%s) maxDepth:%+v, \n 简易报错error: %v\n 含栈信息error:%+v",
			rootPath, maxDepth, errors.Cause(err), errors.Wrap(err, ""))
	}
	t.Logf("BuildByOs(%s) maxDepth:%+v, succeeded: %v", rootPath, maxDepth, rootNode)

	t.Logf("***********************************************")
	printNode(rootNode, t)
}

// printNode 递归函数，用来打印每个节点的信息
func printNode(n *Node, t *testing.T) {
	// 打印当前节点的深度和名称
	t.Logf("dep:%d  name:%s, relativePath:%s\n", n.Depth, n.Name, n.RelativePath)
	//t.Logf("dep:%d  name:%s, parentPath:%s\n", n.Depth, n.Name, n.ParentPath)
	//fmt.Printf("dep:%s  name:%s\n", generateIndent(n.Depth), n.Name)
	// 递归调用，打印所有子节点的信息
	for _, child := range n.Children {
		printNode(child, t)
	}
}
func Test_filePath_Split(t *testing.T) {
	p := "/Users/3phone/Downloads/test/cc/dd:/Users/3phone/Downloads/test/cc/dd/ee/ff"
	for _, a := range filepath.SplitList(p) {
		//t.Logf("SplitList(%s) = %v", p, a)
		t.Logf("a:%+v\n", a)
	}
}
func Test_filePath_Rel(t *testing.T) {
	p := "/Users/3phone/Downloads/test/"

	relPath, err := filepath.Rel(p, "/Users/3phone/Downloads/test/cc/dd/2.mp4")
	if err != nil {
		t.Errorf("Rel(%s) failed: %v", p, err)
	}
	t.Logf("Rel(%s) succeeded: %v", p, relPath)

}
func Test_filePath_Dir(t *testing.T) {
	// /Users/3phone/Downloads/test/cc/dd
	relPath := filepath.Dir("/Users/3phone/Downloads/test/cc/dd/2.mp4")
	t.Logf(" succeeded: %v", relPath)
}

func Test_filePath_Base(t *testing.T) {

	relPath := filepath.Base("/Users/3phone/Downloads/test/cc/dd/2.mp4")
	t.Logf(" succeeded: %v", relPath)
}
