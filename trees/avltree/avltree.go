// References: https://zh.wikipedia.org/wiki/AVL%E6%A0%91
package avltree

import (
	"datastruct/utils"
	"fmt"
)

// Tree 保存 AVL 树的元素。
type Tree struct {
	Root       *Node
	Comparator utils.Comparator
	size       int
}

// Node 是树中的单个元素
type Node struct {
	Key      interface{}
	Value    interface{}
	Parent   *Node
	Children [2]*Node
	b        int8
}

func NewWith(comparator utils.Comparator) *Tree {
	return &Tree{Comparator: comparator}
}

func NewWithIntComparator() *Tree {
	return &Tree{Comparator: utils.IntComparator}
}

func NewWithStringComparator() *Tree {
	return &Tree{Comparator: utils.StringComparator}
}

func (t *Tree) Put(key interface{}, value interface{}) {
	if t.Root != nil {
		fmt.Println("key:", key, "value:", value, "root_key:", t.Root.Key)
	} else {
		fmt.Println("key:", key, "value:", value, "root_key:", nil)
	}

	fmt.Println(t.put(key, value, nil, &t.Root))
}

func (t *Tree) put(key, value interface{}, p *Node, qp **Node) bool {
	fmt.Println("	key:", key, "value:", value)
	q := *qp
	// 父节点为空 创建新的节点作为父节点
	if q == nil {
		t.size++
		*qp = &Node{Key: key, Value: value, Parent: p}
		return true
	}

	c := t.Comparator(key, q.Key)
	//key 相等 替换
	if c == 0 {
		q.Key = key
		q.Value = value
		return false
	}

	// 添加的 key 小于 当前比较的 节点 key
	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2 // a [ 0 , 1 ]
	var fix bool
	fix = t.put(key, value, q, &q.Children[a])
	if fix {
		return putFix(int8(c), qp)
	}
	return false
}

// 返回ture 未改变父节点  false 改变父节点
func putFix(c int8, t **Node) bool {
	s := *t

	fmt.Println("		", "s.key:", s.Key, "s.b:", s.b, "c:", c)

	if s.b == 0 {
		s.b = c
		return true
	}

	if s.b == -c {
		s.b = 0
		return false
	}

	if s.Children[(c+1)/2].b == c {
		fmt.Println("			singlerot")
		s = singlerot(c, s) // 父节点有一个节点,执行单个节点
	} else {
		fmt.Println("			doublerot")
		s = doublerot(c, s)
	}
	*t = s
	return false
}

func singlerot(c int8, s *Node) *Node {
	s.b = 0
	s = rotate(c, s)
	s.b = 0
	return s
}

func doublerot(c int8, s *Node) *Node {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = rotate(-c, s.Children[a])
	p := rotate(c, s)

	switch {
	case p.b == c:
		s.b = -c
		r.b = 0
	case p.b == -c:
		s.b = 0
		r.b = c
	default:
		s.b = 0
		r.b = 0
	}
	p.b = 0
	return p
}

// 计算节点
func rotate(c int8, s *Node) *Node {
	a := (c + 1) / 2
	r := s.Children[a]

	fmt.Println("			r.key", r.Key, "a^1", a^1)
	fmt.Println("			s.key", s.Key)
	fmt.Println("			r.Children[a^1]:", r.Children[a^1])

	s.Children[a] = r.Children[a^1]
	if s.Children[a] != nil {
		s.Children[a].Parent = s
	}
	r.Children[a^1] = s
	r.Parent = s.Parent
	s.Parent = r
	return r
}
