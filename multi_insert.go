/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://github.com/generalzgd/sqlmake
 * @software: GoLand
 * @file: multi_insert.go
 * @time: 2021/1/13 17:15
 */

package sqlmake

import (
	"fmt"
	"strings"
)

// 批量插入
type MultiInsert struct {
	// 表名
	TableName string
	// 要插入的字段名
	FieldList []string
	// 要批量插入的数据
	Rows [][]interface{}
	// DUPLICATE KEY UPDATE部分 对应更新字段以及操作符，默认=，可以修改为+= 或-=
	UpdateOperate map[string]string
}

func (p *MultiInsert) SqlStr() (string, []interface{}) {
	b := strings.Builder{}
	b.WriteString("INSERT INTO `")
	b.WriteString(p.TableName)
	b.WriteString("` (")
	for i, field := range p.FieldList {
		b.WriteString("`")
		b.WriteString(field)
		b.WriteString("`")
		if i < len(p.FieldList)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString(") VALUES ")
	for i, row := range p.Rows {
		b.WriteString("(")
		for j, val := range row {
			b.WriteString(encloseFileValue(val))
			if j < len(row)-1 {
				b.WriteString(",")
			}
		}
		b.WriteString(")")
		if i < len(p.Rows)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString(" ON DUPLICATE KEY UPDATE ")
	cnt := 0
	for field, opt := range p.UpdateOperate {
		switch opt {
		case "=":
			b.WriteString(fmt.Sprintf("`%s`=VALUES(`%s`)", field, field))
		case "+=":
			b.WriteString(fmt.Sprintf("`%s`=`%s`+VALUES(`%s`)", field, field, field))
		case "-=":
			b.WriteString(fmt.Sprintf("`%s`=`%s`-VALUES(`%s`)", field, field, field))
		}
		if cnt < len(p.UpdateOperate)-1 {
			b.WriteString(",")
		}
		cnt += 1
	}
	b.WriteString(";")
	return b.String(), nil
}
