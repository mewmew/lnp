// Package interval implements control flow recovery based on the interval
// method, as described in C. Cifuentes, "A Structuring Algorithm for
// Decompilation", 1993.
//
// At a high-level, the Interval method TODO...
//
// ref: Cifuentes, Cristina. Reverse Compilation Techniques. Queensland
// University of Technology, Brisbane, 1994 [1].
//
// [1] http://www.phatcode.net/res/228/files/decompilation_thesis.pdf
package interval

import (
	"log"
	"os"

	"github.com/mewkiz/pkg/term"
	"github.com/mewmew/lnp/pkg/cfa"
	"github.com/mewmew/lnp/pkg/cfa/primitive"
)

var (
	// dbg represents a logger with the "interval:" prefix, which logs debug
	// messages to standard error.
	dbg = log.New(os.Stderr, term.YellowBold("interval:")+" ", 0)
	// warn represents a logger with the "interval:" prefix, which logs warning
	// messages to standard error.
	warn = log.New(os.Stderr, term.RedBold("interval:")+" ", 0)
)

// Analyze analyzes the given control flow graph and returns the list of
// recovered high-level control flow primitives. The before and after functions
// are invoked if non-nil before and after merging the nodes of located
// primitives.
func Analyze(g cfa.Graph, before, after func(g cfa.Graph, prim *primitive.Primitive)) []*primitive.Primitive {
	prims := []*primitive.Primitive{}
	// Initialize depth-first search visit order.
	initDFSOrder(g)

	// The Structuring Algorithm is not finite Church-Rosser. Thus an ordering is
	// to be followed, namely: structure n-way conditionals, followed by loop
	// structuring, and 2-way conditional structuring last.

	// Structure compound conditionals.
	prims = append(prims, structCompCond(g, before, after)...)
	n := len(prims) // track prims added after compound conditions.
	// Structure n-way conditionals.
	dom := cfa.NewDom(g)
	prims = append(prims, structNway(g, dom)...)
	// Structure loops of the control flow graph.
	prims = append(prims, loopStruct(g, dom)...)
	// Structure 2-way conditionals.
	prims = append(prims, struct2way(g, dom)...)
	for i := n; i < len(prims); i++ {
		if before != nil {
			before(g, prims[i])
		}
		if after != nil {
			after(g, prims[i])
		}
	}
	return prims
}
