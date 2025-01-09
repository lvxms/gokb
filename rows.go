/******************************************************************************
* 版权信息：北京人大金仓信息技术股份有限公司

* 作者：KingbaseES

* 文件名：rows.go

* 功能描述：

* 其它说明：

* 修改记录：
  1.修改时间：
  
  2.修改人：
  
  3.修改内容：

******************************************************************************/

package gokb

import (
	"math"
	"reflect"
	"time"
	"database/sql"

	"kingbase.com/gokb/oid"
)

func (fd fieldDesc) Type() (rt reflect.Type) {
	switch fd.OID {
	case oid.T_int8:
		rt = reflect.TypeOf(int64(0))
		return
	case oid.T_int4:
		rt = reflect.TypeOf(int32(0))
		return
	case oid.T_int2:
		rt = reflect.TypeOf(int16(0))
		return
	case oid.T_tinyint:
		rt = reflect.TypeOf(int8(0))
		return
	case oid.T_bit:
		var b sql.NullBool
		rt = reflect.TypeOf(b)
		return
	case oid.T_varchar, oid.T_text:
		rt = reflect.TypeOf("")
		return
	case oid.T_bool:
		rt = reflect.TypeOf(false)
		return
	case oid.T_date, oid.T_time, oid.T_timetz, oid.T_timestamp, oid.T_timestamptz:
		rt = reflect.TypeOf(time.Time{})
		return
	case oid.T_bytea, oid.T_numeric, oid.T_money:
		rt = reflect.TypeOf([]byte(nil))
		return
	case oid.T_float4:
		rt = reflect.TypeOf(float32(0))
		return
	case oid.T_float8:
		rt = reflect.TypeOf(float64(0))
		return
	default:
		rt = reflect.TypeOf(new(interface{})).Elem()
		return
	}
}

func (fd fieldDesc) Name() (s string) { return oid.TypeName[fd.OID] }

func (fd fieldDesc) Length() (len int64, state bool) {
	switch fd.OID {
	case oid.T_text, oid.T_bytea:
		len = math.MaxInt64
		state = true
		return
	case oid.T_varchar, oid.T_bpchar, oid.T_bpcharbyte, oid.T_varcharbyte, oid.T_nchar, oid.T_nvarchar, oid.T_binary, oid.T_varbinary:
		len = int64(fd.Mod - headerSize)
		state = true
		return
	default:
		len = 0
		state = false
		return
	}
}

func (fd fieldDesc) PrecisionScale() (precision int64, scale int64, state bool) {
	switch fd.OID {
	case oid.T_numeric, oid.T__numeric:
		mod := fd.Mod - headerSize
		precision = int64((mod >> 16) & 0xffff)
		scale = int64(mod & 0xffff)
		state = true
		return
	default:
		precision = 0
		scale = 0
		state = false
		return
	}
}

func (fd fieldDesc) NullAble() (nullable bool, hasNullable bool) {
	switch fd.OID {
	case oid.T_SIMPLE_INTEGER, oid.T_SIMPLE_FLOAT, oid.T_SIMPLE_DOUBLE, oid.T_positiven, oid.T_NATURALN,
		 oid.T__SIMPLE_INTEGER, oid.T__SIMPLE_FLOAT, oid.T__SIMPLE_DOUBLE, oid.T__positiven, oid.T__NATURALN:
		nullable = false
		hasNullable = true
		return
	default:
		nullable = true
		hasNullable = true
		return
	}
}

func (rs *rows) ColumnTypeNullable(index int) (nullable bool, hasNullable bool) {
	return rs.colTyps[index].NullAble()
}

// ColumnTypeScanType返回适合的SCAN数据类型
func (rs *rows) ColumnTypeScanType(index int) (rt reflect.Type) {
	return rs.colTyps[index].Type()
}

// ColumnTypeDatabaseTypeName返回数据的类型名
func (rs *rows) ColumnTypeDatabaseTypeName(index int) (s string) {
	return rs.colTyps[index].Name()
}

// ColumnTypeLength返回列类型的长度
func (rs *rows) ColumnTypeLength(index int) (int64, bool) {
	return rs.colTyps[index].Length()
}

// ColumnTypePrecisionScale返回decimal类型的精度和刻度
func (rs *rows) ColumnTypePrecisionScale(index int) (int64, int64, bool) {
	return rs.colTyps[index].PrecisionScale()
}
