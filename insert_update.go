/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://github.com/generalzgd/sqlmake
 * @software: GoLand
 * @file: insert_update.go
 * @time: 2021/1/11 20:37
 */

package sqlmake

import (
	"fmt"
	"reflect"
	"strings"
)

type InsertUpdate struct {
	// 唯一建，或者主键
	UniqueKeys []string
	// 键值对
	Kvs map[string]interface{}
	// 表名
	TableName string
	// DUPLICATE KEY UPDATE部分 对应更新字段以及操作符，默认=，可以修改为+= 或-=
	UpdateOperate map[string]string
}

func (p *InsertUpdate) AddKvs(kvs map[string]interface{}) {
	for k, v := range kvs {
		val := reflect.ValueOf(v)
		if val.IsZero() {
			continue
		}
		p.Kvs[k] = v
	}
}

func (p *InsertUpdate) SqlStr() (string, []interface{}) {
	UniqueMap := make(map[string]struct{}, len(p.UniqueKeys))
	for _, key := range p.UniqueKeys {
		UniqueMap[key] = struct{}{}
	}
	if p.UpdateOperate == nil {
		p.UpdateOperate = map[string]string{}
	}

	argNames := make([]string, 0, len(p.Kvs))
	argPos := make([]string, 0, len(p.Kvs))
	argValues := make([]interface{}, 0, len(p.Kvs))
	updateArgNames := make([]string, 0, len(p.Kvs))
	//
	for k, v := range p.Kvs {
		argNames = append(argNames, fmt.Sprintf("`%s`", k))
		argPos = append(argPos, "?")
		argValues = append(argValues, v)
		if _, ok := UniqueMap[k]; !ok {
			updateArgNames = append(updateArgNames, k)
		}
	}
	//
	b := strings.Builder{}
	b.WriteString("INSERT INTO ")
	b.WriteString(p.TableName)
	b.WriteString("(")
	b.WriteString(strings.Join(argNames, ","))
	b.WriteString(")")
	b.WriteString("VALUES(")
	b.WriteString(strings.Join(argPos, ","))
	b.WriteString(")")
	b.WriteString(" ON DUPLICATE KEY UPDATE ")
	for i, it := range updateArgNames {
		opt := "="
		if v, ok := p.UpdateOperate[it]; ok {
			opt = v
		}
		switch opt {
		case "=":
			b.WriteString(fmt.Sprintf("`%s`=VALUES(`%s`)", it, it))
		case "+=":
			b.WriteString(fmt.Sprintf("`%s`=`%s`+VALUES(`%s`)", it, it, it))
		case "-=":
			b.WriteString(fmt.Sprintf("`%s`=`%s`-VALUES(`%s`)", it, it, it))
		}
		if i < len(updateArgNames)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString(";")
	return b.String(), argValues
}
