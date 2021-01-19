/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://github.com/generalzgd/sqlmake
 * @software: GoLand
 * @file: insert_update_test.go.go
 * @time: 2021/1/11 20:57
 */

package sqlmake

import (
	"testing"
)

func TestInsertUpdate_SqlStr(t *testing.T) {

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "TestInsertUpdate_SqlStr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := InsertUpdate{
				UniqueKeys: []string{"task_id", "user_id"},
				Kvs: map[string]interface{}{
					"task_id": 3,
					"user_id": 60,
					"status":  2,
				},
				TableName: "user_task_status",
			}
			p.AddKvs(map[string]interface{}{
				"user_code": "code",
				"phone":     "",
				"last_date": "2021-01-11",
			})
			got, got1 := p.SqlStr()
			t.Logf("SqlStr() got = %v, got1 = %v", got, got1)
		})
	}
}
