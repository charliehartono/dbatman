package sql

import (
	"fmt"
)

/*******************************************
 * Table Interfaces and Structs
 * doc:
 *  - table_references http://dev.mysql.com/doc/refman/5.7/en/join.html
 *  - table_factor     http://dev.mysql.com/doc/refman/5.7/en/join.html
 *  - join_table       http://dev.mysql.com/doc/refman/5.7/en/join.html
 ******************************************/
type ITable interface {
	IsTable()
	GetSchemas() []string
}

type ITables []ITable

func (*JoinTable) IsTable()    {}
func (*ParenTable) IsTable()   {}
func (*AliasedTable) IsTable() {}

type JoinTable struct {
	Left  ITable
	Join  []byte
	Right ITable
	// TODO On    BoolExpr
}

func (j *JoinTable) GetSchemas() []string {

	if j.Left == nil {
		panic("join table must have left value")
	}

	if j.Right == nil {
		panic("join table must have right value")
	}

	l := j.Left.GetSchemas()
	r := j.Right.GetSchemas()

	if l == nil && r == nil {
		return nil
	} else if l == nil {
		return r
	} else {
		return l
	}

	return append(l, r...)
}

type ParenTable struct {
	Table ITable
}

func (p *ParenTable) GetSchemas() []string {
	if p.Table == nil {
		return nil
	}
	return p.Table.GetSchemas()
}

type AliasedTable struct {
	TableOrSubQuery interface{} // here may be the table_ident or subquery
	As              []byte
	// TODO IndexHints
}

func (a *AliasedTable) GetSchemas() []string {
	if t, ok := a.TableOrSubQuery.(ITable); ok {
		return t.GetSchemas()
	} else if s, can := a.TableOrSubQuery.(*SubQuery); can {
		return s.SelectStatement.GetSchemas()
	} else {
		panic(fmt.Sprintf("alias table has no table_factor or subquery, element type[%T]", a.TableOrSubQuery))
	}
}

// SimpleTable contains only qualifier, name and a column field
func (*SimpleTable) IsSimpleTable() {}
func (*SimpleTable) IsTable()       {}

type ISimpleTable interface {
	IsSimpleTable()
	ITable
}

type SimpleTable struct {
	Qualifier []byte
	Name      []byte
	Column    []byte
}

func (s *SimpleTable) GetSchemas() []string {
	if s.Qualifier == nil || len(s.Qualifier) == 0 {
		return nil
	}
	return []string{string(s.Qualifier)}
}

type Spname struct {
	Qualifier []byte
	Name      []byte
}

type SchemaInfo struct {
	Name []byte
}
