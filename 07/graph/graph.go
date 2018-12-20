package graph

type Node struct {
	Name string
	In   []*Node
	Out  []*Node
}

type Graph struct {
	Nodes []Node
	Nmap  map[string]*Node // Go: MaP vAlUeS aReN't AdDrEsSaBlE wItH pOiNtErS
}

func (g *Graph) Init() {
	g.Nodes = make([]Node, 0)
	g.Nmap = make(map[string]*Node)
}

func (g *Graph) CheckNode(name string) (exists bool) {
	_, exists = g.Nmap[name]
	return
}

func (g *Graph) AddNode(name string) {
	g.Nodes = append(g.Nodes, Node{name, make([]*Node, 0), make([]*Node, 0)})
	g.Nmap[name] = &g.Nodes[len(g.Nodes)-1]
}

func (g *Graph) AddEdge(inName string, outName string) {
	g.Nmap[inName].In = append(g.Nmap[inName].In, g.Nmap[outName])
	g.Nmap[outName].Out = append(g.Nmap[outName].Out, g.Nmap[inName])
}

func (g *Graph) InstAdd(node string, requires string) {
	// Add Nodes if thhey don't exist yet
	if !g.CheckNode(node) {
		g.AddNode(node)
	}
	if !g.CheckNode(requires) {
		g.AddNode(requires)
	}
	// Not going to check if edges exist

	g.AddEdge(node, requires)
}
