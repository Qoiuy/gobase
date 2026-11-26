# filetree

把本地目录扫描成内存中的文件树，并提供前序遍历。包位于 `gotool/basic` 这个独立 module 中，导入路径为 `github.com/Qoiuy/gotool/basic/filetree`。

## 数据结构

```go
type Node struct {
	Name         string
	Children     []*Node
	Depth        int
	ParentPath   string
	RelativePath string
	Size         uint64
	FileMode     string
	IsFile       uint8
	ParentNode   *Node
}
```

| 字段 | 说明 |
| --- | --- |
| `Name` | 根节点是传入的完整 `rootPath`；子节点是 `filepath.Base` 后的名称 |
| `Children` | 直接子节点列表 |
| `Depth` | 根节点为 0，每下一层加 1 |
| `ParentPath` | 父目录相对 `rootPath` 的路径；顶层节点的父路径为 `.`，根节点为空字符串 |
| `RelativePath` | 相对 `rootPath` 的路径；根节点为 `.` |
| `Size` | `os.FileInfo.Size()` 的结果 |
| `FileMode` | `os.FileInfo.Mode().String()` 的结果 |
| `IsFile` | `0` 表示目录（`IsFileFolder`），`1` 表示文件（`IsFileFile`） |
| `ParentNode` | 父节点；根节点为 `nil` |

## 构建

```go
root, maxDepth, err := filetree.BuildByOs("/path/to/dir")
if err != nil {
	log.Fatal(err)
}

fmt.Println(root.Name, maxDepth)
for _, node := range root.List() {
	fmt.Printf("depth=%d path=%s mode=%s\n", node.Depth, node.RelativePath, node.FileMode)
}
```

- `BuildByOs` 使用 `filepath.Walk` 递归遍历，目录节点先于其子节点被处理，顺序为 `filepath.Walk` 的字典序。
- 返回根节点、最大深度和 error。最大深度以根节点为 0：空目录也是 0，只有一层子项时为 1。
- 构建前会 `os.Stat(rootPath)`；路径不存在或不可读时返回带回溯上下文的 wrapped error，可用 `errors.Cause` 取底层原因。
- 遍历过程中出错会中断并返回，错误信息里带有路径上下文。

## 遍历

`(*Node).List()` 返回前序遍历结果，包含根节点自身；顺序由 `Children` 的追加顺序决定。

## 常量

```go
const (
	IsFileFolder = 0
	IsFileFile   = 1
)
```

`IsFile` 是 `uint8` 的类型别名。

## 限制

- 全量扫描并驻留内存，没有过滤、最大深度或忽略规则；隐藏文件也会进入树。
- 只反映扫描瞬间的状态，不处理扫描后的新增或删除。
- 符号链接按 `filepath.Walk` 的默认行为处理：不会递归进入链接指向的目录。
- 测试 `Test_BuildByOs` 把路径写死为 `/Users/3phone/Downloads/test`；在该目录不存在时测试会失败，换机器需要先改测试路径或准备对应目录。
