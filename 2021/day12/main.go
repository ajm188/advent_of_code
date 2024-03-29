package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
	"github.com/ajm188/advent_of_code/pkg/sets"
)

type node struct {
	name  string
	nodes []*node
}

func (n *node) String() string { return n.name }

type nodes []*node

func (nl nodes) Join(sep string) string {
	names := make([]string, len(nl))
	for i, n := range nl {
		names[i] = n.name
	}

	return strings.Join(names, sep)
}

func (nl nodes) Reverse() nodes {
	nl2 := make([]*node, 0, len(nl))
	for _, n := range nl {
		nl2 = append([]*node{n}, nl2...)
	}

	return nodes(nl2)
}

func (nl nodes) Len() int           { return len(nl) }
func (nl nodes) Swap(i, j int)      { nl[i], nl[j] = nl[j], nl[i] }
func (nl nodes) Less(i, j int) bool { return nl[i].name < nl[j].name }

var (
	Start = &node{name: "start"}
	End   = &node{name: "end"}

	debug = flag.Bool("debug", false, "")
)

type iroute interface {
	Add(*node) iroute
	Nodes() []*node
}

type route struct {
	nodes []*node
	seen  *sets.Strings
}

func (r *route) Add(n *node) iroute {
	if r.seen == nil {
		r.seen = sets.NewStrings()
		for _, n := range r.nodes {
			r.seen.Insert(n.name)
		}
	}

	if strings.ToLower(n.name) == n.name && r.seen.Has(n.name) {
		return nil
	}

	r2 := &route{
		nodes: append([]*node{n}, r.nodes...),
		seen:  r.seen.Copy(),
	}
	r2.seen.Insert(n.name)
	return r2
}

func (r *route) Nodes() []*node { return r.nodes }

type twiceroute struct {
	nodes     []*node
	seen      *sets.Strings
	revisited *node
}

func (r *twiceroute) Add(n *node) iroute {
	if r.seen == nil {
		r.seen = sets.NewStrings()
		for _, n := range r.nodes {
			r.seen.Insert(n.name)
		}
	}

	revisited := r.revisited
	if strings.ToLower(n.name) == n.name && r.seen.Has(n.name) {
		if n == Start || revisited != nil {
			return nil
		}

		revisited = n
	}

	r2 := &twiceroute{
		nodes:     append([]*node{n}, r.nodes...),
		seen:      r.seen.Copy(),
		revisited: revisited,
	}
	r2.seen.Insert(n.name)
	return r2
}

func (r *twiceroute) Nodes() []*node { return r.nodes }

func findPaths(routes []iroute) int {
	var count int
	for len(routes) != 0 {
		var newroutes []iroute
		for _, r := range routes {
			if *debug {
				log.Printf("route: %v\n", nodes(r.Nodes()).Reverse().Join(" -> "))
			}

			if r.Nodes()[0] == End {
				count++
				continue
			}

			for _, next := range r.Nodes()[0].nodes {
				if r2 := r.Add(next); r2 != nil {
					newroutes = append(newroutes, r2)
				}
			}
		}

		routes = newroutes
	}

	return count
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	nodeMap := map[string]*node{
		Start.name: Start,
		End.name:   End,
	}
	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			cli.ExitOnError(fmt.Errorf("malformed input: line %d does not match '<name>-<name>'", i))
		}

		left, ok := nodeMap[parts[0]]
		if !ok {
			left = &node{name: parts[0]}
		}

		right, ok := nodeMap[parts[1]]
		if !ok {
			right = &node{name: parts[1]}
		}

		left.nodes = append(left.nodes, right)
		right.nodes = append(right.nodes, left)

		nodeMap[left.name] = left
		nodeMap[right.name] = right
	}

	for _, node := range nodeMap {
		sort.Sort(nodes(node.nodes))
		if *debug {
			log.Printf("%s\t%v\n", node.name, node.nodes)
		}
	}

	fmt.Println(findPaths([]iroute{
		&route{nodes: []*node{Start}},
	}))

	fmt.Println(findPaths([]iroute{
		&twiceroute{nodes: []*node{Start}},
	}))

}
