package ast

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tboronczyk/kiwi/token"
)

func capture(f func()) string {
	// re-assign stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()
	// read output
	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		out <- buf.String()
	}()

	// restore stdout
	w.Close()
	os.Stdout = old

	return <-out
}

func TestPrintValueNodeNumber(t *testing.T) {
	expected := "ValueNode\n" +
		"├ Value: 1\n" +
		"╰ Type: NUMBER\n"
	actual := capture(func() {
		n := &ValueNode{Value: "1.0", Type: token.NUMBER}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintValueNodeString(t *testing.T) {
	expected := "ValueNode\n" +
		"├ Value: \"foo\"\n" +
		"╰ Type: STRING\n"
	actual := capture(func() {
		n := &ValueNode{Value: "foo", Type: token.STRING}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintValueNodeBool(t *testing.T) {
	expected := "ValueNode\n" +
		"├ Value: true\n" +
		"╰ Type: BOOL\n"
	actual := capture(func() {
		n := &ValueNode{Value: "True", Type: token.BOOL}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintCast(t *testing.T) {
	expected := "CastNode\n" +
		"├ Cast: string\n" +
		"╰ Term: ValueNode\n" +
		"        ├ Value: \"foo\"\n" +
		"        ╰ Type: STRING\n"
	actual := capture(func() {
		n := &CastNode{
			Cast: "string",
			Term: &ValueNode{Value: "foo", Type: token.STRING},
		}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintVariableNode(t *testing.T) {
	expected := "VariableNode\n" +
		"╰ Name: foo\n"
	actual := capture(func() {
		n := &VariableNode{Name: "foo"}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintUnaryOpNode(t *testing.T) {
	expected := "UnaryOpNode\n" +
		"├ Op: -\n" +
		"╰ Term: VariableNode\n" +
		"        ╰ Name: foo\n"
	actual := capture(func() {
		n := &UnaryOpNode{
			Op:   token.SUBTRACT,
			Term: &VariableNode{Name: "foo"},
		}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintBinOpNode(t *testing.T) {
	expected := "BinOpNode\n" +
		"├ Op: +\n" +
		"├ Left: VariableNode\n" +
		"│       ╰ Name: foo\n" +
		"╰ Right: VariableNode\n" +
		"         ╰ Name: bar\n"
	actual := capture(func() {
		n := &BinOpNode{
			Op:    token.ADD,
			Left:  &VariableNode{Name: "foo"},
			Right: &VariableNode{Name: "bar"},
		}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintAssignNode(t *testing.T) {
	expected := "AssignNode\n" +
		"├ Name: foo\n" +
		"╰ Expr: VariableNode\n" +
		"        ╰ Name: bar\n"
	actual := capture(func() {
		n := &AssignNode{
			Name: "foo",
			Expr: &VariableNode{Name: "bar"},
		}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintProgramNode(t *testing.T) {
	expected := "ProgramNode\n" +
		"╰ Stmts: AssignNode\n" +
		"         ├ Name: foo\n" +
		"         ╰ Expr: VariableNode\n" +
		"                 ╰ Name: bar\n" +
		"         AssignNode\n" +
		"         ├ Name: baz\n" +
		"         ╰ Expr: VariableNode\n" +
		"                 ╰ Name: quux\n"
	actual := capture(func() {
		n := &ProgramNode{
			Stmts: []Node{
				&AssignNode{
					Name: "foo",
					Expr: &VariableNode{Name: "bar"},
				},
				&AssignNode{
					Name: "baz",
					Expr: &VariableNode{Name: "quux"},
				},
			},
		}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintProgramNodeEmpty(t *testing.T) {
	expected := "ProgramNode\n" +
		"╰ Stmts: 0x0\n"
	actual := capture(func() {
		n := &ProgramNode{}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintFuncDefNodeNoArgsOrBody(t *testing.T) {
	expected := "FuncDefNode\n" +
		"├ Name: foo\n" +
		"├ Args: 0x0\n" +
		"╰ Body: 0x0\n"
	actual := capture(func() {
		n := &FuncDefNode{Name: "foo"}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintFuncDefNode(t *testing.T) {
	expected := "FuncDefNode\n" +
		"├ Name: foo\n" +
		"├ Args: bar\n" +
		"│       baz\n" +
		"╰ Body: AssignNode\n" +
		"        ├ Name: bar\n" +
		"        ╰ Expr: VariableNode\n" +
		"                ╰ Name: baz\n" +
		"        ReturnNode\n" +
		"        ╰ Expr: VariableNode\n" +
		"                ╰ Name: baz\n"
	actual := capture(func() {
		n := &FuncDefNode{
			Name: "foo",
			Args: []string{"bar", "baz"},
			Body: []Node{
				&AssignNode{
					Name: "bar",
					Expr: &VariableNode{Name: "baz"},
				},
				&ReturnNode{Expr: &VariableNode{Name: "baz"}},
			},
		}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintFuncCallNodeNoArgs(t *testing.T) {
	expected := "FuncCallNode\n" +
		"├ Name: foo\n" +
		"╰ Args: 0x0\n"
	actual := capture(func() {
		n := &FuncCallNode{Name: "foo"}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintFuncCallNode(t *testing.T) {
	expected := "FuncCallNode\n" +
		"├ Name: foo\n" +
		"╰ Args: VariableNode\n" +
		"        ╰ Name: foo\n" +
		"        VariableNode\n" +
		"        ╰ Name: bar\n"
	actual := capture(func() {
		n := &FuncCallNode{
			Name: "foo",
			Args: []Node{
				&VariableNode{Name: "foo"},
				&VariableNode{Name: "bar"},
			},
		}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintIfNodeNoBody(t *testing.T) {
	expected := "IfNode\n" +
		"├ Cond: VariableNode\n" +
		"│       ╰ Name: foo\n" +
		"├ Body: 0x0\n" +
		"╰ Else: 0x0\n"
	actual := capture(func() {
		n := &IfNode{Cond: &VariableNode{Name: "foo"}}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintIfNode(t *testing.T) {
	expected := "IfNode\n" +
		"├ Cond: VariableNode\n" +
		"│       ╰ Name: foo\n" +
		"├ Body: AssignNode\n" +
		"│       ├ Name: bar\n" +
		"│       ╰ Expr: VariableNode\n" +
		"│               ╰ Name: baz\n" +
		"│       AssignNode\n" +
		"│       ├ Name: quux\n" +
		"│       ╰ Expr: VariableNode\n" +
		"│               ╰ Name: norf\n" +
		"╰ Else: IfNode\n" +
		"        ├ Cond: ValueNode\n" +
		"        │       ├ Value: true\n" +
		"        │       ╰ Type: BOOL\n" +
		"        ├ Body: 0x0\n" +
		"        ╰ Else: 0x0\n"
	actual := capture(func() {
		n := &IfNode{
			Cond: &VariableNode{Name: "foo"},
			Body: []Node{
				&AssignNode{
					Name: "bar",
					Expr: &VariableNode{Name: "baz"},
				},
				&AssignNode{
					Name: "quux",
					Expr: &VariableNode{Name: "norf"},
				},
			},
			Else: &IfNode{
				Cond: &ValueNode{
					Value: "true",
					Type:  token.BOOL,
				},
			},
		}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintReturnNode(t *testing.T) {
	expected := "ReturnNode\n" +
		"╰ Expr: VariableNode\n" +
		"        ╰ Name: foo\n"
	actual := capture(func() {
		n := &ReturnNode{Expr: &VariableNode{Name: "foo"}}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintWhileNodeNoBody(t *testing.T) {
	expected := "WhileNode\n" +
		"├ Cond: VariableNode\n" +
		"│       ╰ Name: foo\n" +
		"╰ Body: 0x0\n"
	actual := capture(func() {
		n := &WhileNode{Cond: &VariableNode{Name: "foo"}}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}

func TestPrintWhileNode(t *testing.T) {
	expected := "WhileNode\n" +
		"├ Cond: VariableNode\n" +
		"│       ╰ Name: foo\n" +
		"╰ Body: AssignNode\n" +
		"        ├ Name: bar\n" +
		"        ╰ Expr: VariableNode\n" +
		"                ╰ Name: baz\n" +
		"        AssignNode\n" +
		"        ├ Name: quux\n" +
		"        ╰ Expr: VariableNode\n" +
		"                ╰ Name: norf\n"
	actual := capture(func() {
		n := &WhileNode{
			Cond: &VariableNode{Name: "foo"},
			Body: []Node{
				&AssignNode{
					Name: "bar",
					Expr: &VariableNode{Name: "baz"},
				},
				&AssignNode{
					Name: "quux",
					Expr: &VariableNode{Name: "norf"},
				},
			},
		}
		n.Accept(NewAstPrinter())
	})
	assert.Equal(t, expected, actual)
}
