/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://github.com/generalzgd/sqlmake
 * @software: GoLand
 * @file: update_case_test.go.go
 * @time: 2021/1/13 17:02
 */

package sqlmake

import (
	"testing"
)

func TestUpdateCase_SqlStr(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "TestUpdateCase_SqlStr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &UpdateCase{
				TableName:    "tableName",
				FieldName:    "abc",
				KeyFieldName: "id",
				CaseMap: map[interface{}]interface{}{
					1:  "0",
					34: "35",
				},
			}
			if got, _ := p.SqlStr(); got != tt.want {
				t.Logf("SqlStr() = %v", got)
			}
		})
	}
}
