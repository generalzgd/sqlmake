/**
 * @version: v0.1.0
 * @author: zhangguodong
 * @license: LGPL v3
 * @contact: zhangguodong@dobest.com
 * @site: https://github.com/generalzgd/sqlmake
 * @software: GoLand
 * @file: sqlmaker.go
 * @time: 2021/1/13 17:19
 */

package sqlmake

import (
	"fmt"
	"reflect"
)

type SqlMaker interface {
	// 返回sql语句和对应的参数列表
	SqlStr() (string, []interface{})
}

func encloseFileValue(val interface{}) string {
	w := reflect.ValueOf(val)
	switch w.Kind() {
	// case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	// 	fallthrough
	// case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	// 	return fmt.Sprintf("%v", val)
	case reflect.String:
		return fmt.Sprintf("'%v'", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}
