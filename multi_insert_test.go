/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://github.com/generalzgd/sqlmake
 * @software: GoLand
 * @file: multi_insert_test.go.go
 * @time: 2021/1/13 17:28
 */

package sqlmake

import (
	"testing"
)

func TestMultiInsert_SqlStr(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "TestMultiInsert_SqlStr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MultiInsert{
				TableName: "TableName",
				FieldList: []string{"a", "b", "c"},
				Rows: [][]interface{}{
					{
						1, "a", 4,
					},
					{
						2, "5h", 41,
					},
				},
				UpdateOperate: map[string]string{
					"a": "=",
				},
			}
			got, _ := p.SqlStr()
			t.Logf("SqlStr() got1 = %v", got)
		})
	}
}
