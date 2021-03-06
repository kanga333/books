package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var preIndex = 0

type tree struct {
	val    int
	parent *tree
	left   *tree
	right  *tree
}

func (t *tree) print() {
	fmt.Print(t.val)
	fmt.Print(" left ")
	if t.left != nil {
		fmt.Print(t.left.val)
	} else {
		fmt.Print("nil")
	}
	fmt.Print(" right ")
	if t.right != nil {
		fmt.Print(t.right.val)
	} else {
		fmt.Print("nil")
	}
	fmt.Println("")
}

func (t *tree) insert(z *tree) {
	var p *tree
	x := t
	for x != nil {
		p = x
		if z.val < x.val {
			x = x.left
			continue
		}
		x = x.right
	}
	z.parent = p

	if z.val < p.val {
		p.left = z
		return
	}
	p.right = z
}

func (t *tree) find(val int) *tree {
	if t.val == val {
		return t
	}

	if t.left != nil && t.val > val {
		return t.left.find(val)
	}

	if t.right != nil && t.val < val {
		return t.right.find(val)
	}

	return nil
}

func (t *tree) delete() {
	if t.left != nil && t.right != nil {
		n := t.right.nextInOrder()
		t.val = n.val
		n.delete()
		return
	}

	var tmp *tree
	if t.left != nil {
		tmp = t.left
	}

	if t.right != nil {
		tmp = t.right
	}

	p := t.parent
	if tmp != nil {
		tmp.parent = p
	}

	if p.left != nil && p.left.val == t.val {
		p.left = tmp
		return
	}
	p.right = tmp

	return
}

func (t *tree) nodeType() string {
	if t.parent == nil && t.val != 0 {
		return "root"
	}
	if t.left == nil && t.right == nil {
		return "leaf"
	}

	return "internal node"
}

func (t *tree) preOrder(w io.Writer) {
	fmt.Fprintf(w, " %d", t.val)
	if t.left != nil {
		t.left.preOrder(w)
	}
	if t.right != nil {
		t.right.preOrder(w)
	}

	return
}

func (t *tree) nextInOrder() *tree {
	if t.left != nil {
		return t.left.nextInOrder()
	}

	return t
}

func (t *tree) inOrder(w io.Writer) {
	if t.left != nil {
		t.left.inOrder(w)
	}
	fmt.Fprintf(w, " %d", t.val)

	if t.right != nil {
		t.right.inOrder(w)
	}

	return
}

func parseTree(s string) *tree {
	a := strings.Split(s, " ")
	val, err := strconv.Atoi(a[1])
	if err != nil {
		panic(err)
	}
	return &tree{
		val: val,
	}
}

func parseVal(s string) int {
	a := strings.Split(s, " ")
	val, err := strconv.Atoi(a[1])
	if err != nil {
		panic(err)
	}
	return val
}

//  ALDS1_8_C: 二分木探索:削除
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	stdin.Scan() //skip scan n

	var p *tree
	for stdin.Scan() {
		txt := stdin.Text()

		switch {
		case strings.HasPrefix(txt, "insert"):
			t := parseTree(txt)
			if p == nil {
				p = t
				continue
			}
			p.insert(t)
		case strings.HasPrefix(txt, "print"):
			p.inOrder(w)
			fmt.Fprintf(w, "\n")
			p.preOrder(w)
			fmt.Fprintf(w, "\n")
		case strings.HasPrefix(txt, "delete"):
			val := parseVal(txt)
			if p == nil {
				continue
			}
			t := p.find(val)
			if t == nil {
				continue
			}
			t.delete()
		case strings.HasPrefix(txt, "find"):
			val := parseVal(txt)
			if p != nil && p.find(val) != nil {
				fmt.Fprintf(w, "yes\n")
				continue
			}
			fmt.Fprintf(w, "no\n")
		default:
			panic(fmt.Sprintf("Unexpected case: %s", txt))
		}
	}
}

func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	t := sc.Text()
	i, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}
	return i
}

func scanIntArray(sc *bufio.Scanner) []int {
	sc.Scan()
	t := sc.Text()
	ss := strings.Split(t, " ")
	is := []int{}
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		is = append(is, i)
	}
	return is
}
