package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/vektah/gqlparser/ast"
)

type CommentQueryFilter struct {
	Query *string
}

func (qf *CommentQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field.")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, "comments", &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *CommentQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["text"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("text")+" LIKE ? OR %[1]s"+dialect.Quote("text")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if _, ok := fieldsMap["reference"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("reference")+" LIKE ? OR %[1]s"+dialect.Quote("reference")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	return nil
}
