package main

import (
	"go/ast"
	"go/token"
)

func init() {
	register(killposFix)
}

var killposFix = fix{
	"killpos",
	"2018-01-09",
	killpos,
	`Remove position information.`,
	true, // disabled by default.
}

func killpos(file *ast.File) bool {
	fixed := false
	walk(file, func(n interface{}) {
		switch n := n.(type) {
		case *ast.ArrayType:
			if n == nil {
				return
			}
			if n.Lbrack != token.NoPos {
				fixed = true
			}
			n.Lbrack = token.NoPos
		case *ast.AssignStmt:
			if n == nil {
				return
			}
			if n.TokPos != token.NoPos {
				fixed = true
			}
			n.TokPos = token.NoPos
		case *ast.BadDecl:
			if n == nil {
				return
			}
			if n.From != token.NoPos {
				fixed = true
			}
			n.From = token.NoPos
			if n.To != token.NoPos {
				fixed = true
			}
			n.To = token.NoPos
		case *ast.BadExpr:
			if n == nil {
				return
			}
			if n.From != token.NoPos {
				fixed = true
			}
			n.From = token.NoPos
			if n.To != token.NoPos {
				fixed = true
			}
			n.To = token.NoPos
		case *ast.BadStmt:
			if n == nil {
				return
			}
			if n.From != token.NoPos {
				fixed = true
			}
			n.From = token.NoPos
			if n.To != token.NoPos {
				fixed = true
			}
			n.To = token.NoPos
		case *ast.BasicLit:
			if n == nil {
				return
			}
			if n.ValuePos != token.NoPos {
				fixed = true
			}
			n.ValuePos = token.NoPos
		case *ast.BinaryExpr:
			if n == nil {
				return
			}
			if n.OpPos != token.NoPos {
				fixed = true
			}
			n.OpPos = token.NoPos
		case *ast.BlockStmt:
			if n == nil {
				return
			}
			if n.Lbrace != token.NoPos {
				fixed = true
			}
			n.Lbrace = token.NoPos
			if n.Rbrace != token.NoPos {
				fixed = true
			}
			n.Rbrace = token.NoPos
		case *ast.BranchStmt:
			if n == nil {
				return
			}
			if n.TokPos != token.NoPos {
				fixed = true
			}
			n.TokPos = token.NoPos
		case *ast.CallExpr:
			if n == nil {
				return
			}
			if n.Lparen != token.NoPos {
				fixed = true
			}
			n.Lparen = token.NoPos
			if n.Ellipsis != token.NoPos {
				fixed = true
			}
			n.Ellipsis = token.NoPos
			if n.Rparen != token.NoPos {
				fixed = true
			}
			n.Rparen = token.NoPos
		case *ast.CaseClause:
			if n == nil {
				return
			}
			if n.Case != token.NoPos {
				fixed = true
			}
			n.Case = token.NoPos
			if n.Colon != token.NoPos {
				fixed = true
			}
			n.Colon = token.NoPos
		case *ast.ChanType:
			if n == nil {
				return
			}
			if n.Begin != token.NoPos {
				fixed = true
			}
			n.Begin = token.NoPos
			if n.Arrow != token.NoPos {
				fixed = true
			}
			n.Arrow = token.NoPos
		case *ast.CommClause:
			if n == nil {
				return
			}
			if n.Case != token.NoPos {
				fixed = true
			}
			n.Case = token.NoPos
			if n.Colon != token.NoPos {
				fixed = true
			}
			n.Colon = token.NoPos
		case *ast.Comment:
			if n == nil {
				return
			}
			if n.Slash != token.NoPos {
				fixed = true
			}
			n.Slash = token.NoPos
		case *ast.CompositeLit:
			if n == nil {
				return
			}
			if n.Lbrace != token.NoPos {
				fixed = true
			}
			n.Lbrace = token.NoPos
			if n.Rbrace != token.NoPos {
				fixed = true
			}
			n.Rbrace = token.NoPos
		case *ast.DeferStmt:
			if n == nil {
				return
			}
			if n.Defer != token.NoPos {
				fixed = true
			}
			n.Defer = token.NoPos
		case *ast.Ellipsis:
			if n == nil {
				return
			}
			if n.Ellipsis != token.NoPos {
				fixed = true
			}
			n.Ellipsis = token.NoPos
		case *ast.EmptyStmt:
			if n == nil {
				return
			}
			if n.Semicolon != token.NoPos {
				fixed = true
			}
			n.Semicolon = token.NoPos
		case *ast.FieldList:
			if n == nil {
				return
			}
			if n.Opening != token.NoPos {
				fixed = true
			}
			n.Opening = token.NoPos
			if n.Closing != token.NoPos {
				fixed = true
			}
			n.Closing = token.NoPos
		case *ast.File:
			if n == nil {
				return
			}
			if n.Package != token.NoPos {
				fixed = true
			}
			n.Package = token.NoPos
		case *ast.ForStmt:
			if n == nil {
				return
			}
			if n.For != token.NoPos {
				fixed = true
			}
			n.For = token.NoPos
		case *ast.FuncType:
			if n == nil {
				return
			}
			if n.Func != token.NoPos {
				fixed = true
			}
			n.Func = token.NoPos
		case *ast.GenDecl:
			if n == nil {
				return
			}
			if n.TokPos != token.NoPos {
				fixed = true
			}
			n.TokPos = token.NoPos
			if n.Lparen != token.NoPos {
				fixed = true
			}
			n.Lparen = token.NoPos
			if n.Rparen != token.NoPos {
				fixed = true
			}
			n.Rparen = token.NoPos
		case *ast.GoStmt:
			if n == nil {
				return
			}
			if n.Go != token.NoPos {
				fixed = true
			}
			n.Go = token.NoPos
		case *ast.Ident:
			if n == nil {
				return
			}
			// Keep position information for AST identifiers. Otherwise, selection
			// expressions get broken into two lines. E.g.
			//
			//    os.
			//    	Args

			//if n.NamePos != token.NoPos {
			//	fixed = true
			//}
			//n.NamePos = token.NoPos
		case *ast.IfStmt:
			if n == nil {
				return
			}
			if n.If != token.NoPos {
				fixed = true
			}
			n.If = token.NoPos
		case *ast.ImportSpec:
			if n == nil {
				return
			}
			if n.EndPos != token.NoPos {
				fixed = true
			}
			n.EndPos = token.NoPos
		case *ast.IncDecStmt:
			if n == nil {
				return
			}
			if n.TokPos != token.NoPos {
				fixed = true
			}
			n.TokPos = token.NoPos
		case *ast.IndexExpr:
			if n == nil {
				return
			}
			if n.Lbrack != token.NoPos {
				fixed = true
			}
			n.Lbrack = token.NoPos
			if n.Rbrack != token.NoPos {
				fixed = true
			}
			n.Rbrack = token.NoPos
		case *ast.InterfaceType:
			if n == nil {
				return
			}
			if n.Interface != token.NoPos {
				fixed = true
			}
			n.Interface = token.NoPos
		case *ast.KeyValueExpr:
			if n == nil {
				return
			}
			if n.Colon != token.NoPos {
				fixed = true
			}
			n.Colon = token.NoPos
		case *ast.LabeledStmt:
			if n == nil {
				return
			}
			if n.Colon != token.NoPos {
				fixed = true
			}
			n.Colon = token.NoPos
		case *ast.MapType:
			if n == nil {
				return
			}
			if n.Map != token.NoPos {
				fixed = true
			}
			n.Map = token.NoPos
		case *ast.ParenExpr:
			if n == nil {
				return
			}
			if n.Lparen != token.NoPos {
				fixed = true
			}
			n.Lparen = token.NoPos
			if n.Rparen != token.NoPos {
				fixed = true
			}
			n.Rparen = token.NoPos
		case *ast.RangeStmt:
			if n == nil {
				return
			}
			if n.For != token.NoPos {
				fixed = true
			}
			n.For = token.NoPos
			if n.TokPos != token.NoPos {
				fixed = true
			}
			n.TokPos = token.NoPos
		case *ast.ReturnStmt:
			if n == nil {
				return
			}
			if n.Return != token.NoPos {
				fixed = true
			}
			n.Return = token.NoPos
		case *ast.SelectStmt:
			if n == nil {
				return
			}
			if n.Select != token.NoPos {
				fixed = true
			}
			n.Select = token.NoPos
		case *ast.SendStmt:
			if n == nil {
				return
			}
			if n.Arrow != token.NoPos {
				fixed = true
			}
			n.Arrow = token.NoPos
		case *ast.SliceExpr:
			if n == nil {
				return
			}
			if n.Lbrack != token.NoPos {
				fixed = true
			}
			n.Lbrack = token.NoPos
			if n.Rbrack != token.NoPos {
				fixed = true
			}
			n.Rbrack = token.NoPos
		case *ast.StarExpr:
			if n == nil {
				return
			}
			if n.Star != token.NoPos {
				fixed = true
			}
			n.Star = token.NoPos
		case *ast.StructType:
			if n == nil {
				return
			}
			if n.Struct != token.NoPos {
				fixed = true
			}
			n.Struct = token.NoPos
		case *ast.SwitchStmt:
			if n == nil {
				return
			}
			if n.Switch != token.NoPos {
				fixed = true
			}
			n.Switch = token.NoPos
		case *ast.TypeAssertExpr:
			if n == nil {
				return
			}
			if n.Lparen != token.NoPos {
				fixed = true
			}
			n.Lparen = token.NoPos
			if n.Rparen != token.NoPos {
				fixed = true
			}
			n.Rparen = token.NoPos
		case *ast.TypeSpec:
			if n == nil {
				return
			}
			if n.Assign != token.NoPos {
				fixed = true
			}
			n.Assign = token.NoPos
		case *ast.TypeSwitchStmt:
			if n == nil {
				return
			}
			if n.Switch != token.NoPos {
				fixed = true
			}
			n.Switch = token.NoPos
		case *ast.UnaryExpr:
			if n == nil {
				return
			}
			if n.OpPos != token.NoPos {
				fixed = true
			}
			n.OpPos = token.NoPos
		}
	})
	return fixed
}
