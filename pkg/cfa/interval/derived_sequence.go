package interval

import (
	"fmt"

	"github.com/graphism/simple"
	"github.com/mewmew/lnp/pkg/cfa"
	"github.com/mewmew/lnp/pkg/cfg"
)

// DerivedSequence returns the derived sequence of graphs G^1, ..., G^n, based
// on the intervals of the given control flow graph, and the associated unique
// sets of intervas, Is^1, ..., Is^n.
//
// Pre: G is a control flow graph.
//
// Post: the derived sequence of G, G^1 ... G^n, n >= 1 has been constructed.
//
// ref: Figure 6-10; Cifuentes' Reverse Comilation Techniques.
func DerivedSequence(g cfa.Graph) ([]cfa.Graph, [][]*Interval) {
	// Note, the Go code is zero-indexed, as compared to the Cifuentes' algorithm
	// notation which is 1-indexed.
	// G^1 = G
	Gs := []cfa.Graph{g}
	// IIs^1 = intervals(G^1)
	IIs := [][]*Interval{Intervals(Gs[0])}
	// i = 2
	// repeat, Construction of G^i
	intNum := 1
	for i := 1; ; i++ {
		Gprev := Gs[i-1]
		// Make each interval of G^{i-1} a node in G^i.
		//
		//    N^i = {n^i | I^{i-1}(n^{i-1}) \in IIs^{i-1}}
		Gi := NewGraph()
		var ns []*Node
		for _, I := range IIs[i-1] {
			n := Gi.NewNode().(*Node)
			n.SetDOTID(fmt.Sprintf("I_%d", intNum))
			ns = append(ns, n)
			Gi.AddNode(n)
			if g.Entry().ID() == I.head.ID() {
				Gi.SetEntry(n)
			}
			intNum++
		}
		initDFSOrder(Gi)
		// The collapsed node n of an interval I(h) has the immediate predecessors
		// of h not part of the interval I(h).
		//
		//    immedPreds(n) n \in G^i = immedPreds(h) :
		//       immedPred(h) \not \in I^{i-1}(h)
		//
		//    \forall n \in N^i, p \in immedPred(n) <=> (
		//          \exists m \in N^{i-1},
		//          m \in I^{i-1}(m)
		//          \land p \in immedPred(m)
		//          \land p \not \in I^{i-1}(m)
		//       )
		for j, I := range IIs[i-1] {
			n := ns[j]
			for preds := Gprev.To(I.head.ID()); preds.Next(); {
				p := preds.Node().(cfa.Node)
				if I.Node(p.ID()) != nil {
					// skip predecessor p if present in interval I(h).
					continue
				}
				// Find interval to which p belongs, so that we can connect the
				// header node of that interval with p in the derived graph.
				var pred cfa.Node
				for k, J := range IIs[i-1] {
					if J.Node(p.ID()) != nil {
						pred = ns[k]
						break
					}
				}
				if pred == nil {
					panic(fmt.Errorf("unable to locate interval to which node %q belong", p.DOTID()))
				}
				e := &cfg.Edge{
					Edge: simple.Edge{F: pred, T: n},
				}
				Gi.SetEdge(e)
			}
		}
		// The collapsed node n of an interval I(h) has the immediate successors
		// of the exit nodes of I(h) not part of the interval I(h).
		//
		//    (a, b) \in G^i iff \exists n \in I^{i-1}(h)
		//       \land m = header(I^{i-1}(m)) : (m, n) \in G^{i-1}
		//
		//    (h_j^i, h_k^i) \in E^i <=> (
		//       \exists n, m, h_j^{i-1}, h_k^{i-1} \in N^{i-1},
		//       h_j^{i-1} = I^{i-1}(h_j^{i-1})
		//       \land h_k^{i-1} = I^{i-1}(h_k^{i-1})
		//       \land m \in I^{i-1}(h_j^{i-1})
		//       \land n \in I^{i-1}(h_k^{i-1})
		//       \land (m, n) \in E^{i-1}
		//    )
		for j, I := range IIs[i-1] {
			n := ns[j]
			for succs := Gprev.From(I.head.ID()); succs.Next(); {
				s := succs.Node().(cfa.Node)
				if I.Node(s.ID()) != nil {
					// skip successor s if present in interval I(h).
					continue
				}
				// Find interval to which s belongs, so that we can connect the
				// header node of that interval with s in the derived graph.
				var succ cfa.Node
				for k, J := range IIs[i-1] {
					if J.Node(s.ID()) != nil {
						succ = ns[k]
						break
					}
				}
				if succ == nil {
					panic(fmt.Errorf("unable to locate interval to which node %s belong", s.DOTID()))
				}
				e := &cfg.Edge{
					Edge: simple.Edge{F: n, T: succ},
				}
				Gi.SetEdge(e)
			}
		}
		if Gi.Nodes().Len() == Gs[i-1].Nodes().Len() {
			// until G^i == G^{i-1}
			break
		}
		Gs = append(Gs, Gi)
		// Is^i = intervals(G^i)
		IIs = append(IIs, Intervals(Gi))
		// i = i + 1
	}
	return Gs, IIs
}
