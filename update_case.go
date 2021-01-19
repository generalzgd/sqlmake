/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://github.com/generalzgd/sqlmake
 * @software: GoLand
 * @file: update_case.go
 * @time: 2021/1/13 16:53
 */

package sqlmake

import (
	"fmt"
	"strings"
)

// 更加某字段更新对应值
type UpdateCase struct {
	// 表名
	TableName string
	// 要更新的字段名
	FieldName string
	// 主键/或唯一建字段名
	KeyFieldName string
	// when -> then
	CaseMap map[interface{}]interface{}
}

func (p *UpdateCase) SqlStr() (string, []interface{}) {
	b := strings.Builder{}
	b.WriteString("UPDATE `")
	b.WriteString(p.TableName)
	b.WriteString("` ")
	b.WriteString("SET `")
	b.WriteString(p.FieldName)
	b.WriteString("` = CASE `")
	b.WriteString(p.KeyFieldName)
	b.WriteString("` ")
	for when, then := range p.CaseMap {
		b.WriteString(fmt.Sprintf("WHEN %v THEN %v ", encloseFileValue(when), encloseFileValue(then)))
	}
	b.WriteString("END;")

	return b.String(), nil
}
