package mysql

import "fmt"

// https://dev.mysql.com/doc/internals/en/com-query-response.html#packet-Protocol::ColumnType
func typeStr(t byte) string {
	switch t {
	case fieldTypeDecimal,
		fieldTypeFloat,
		fieldTypeDouble,
		fieldTypeNewDecimal:
		return "float"
	case fieldTypeTiny,
		fieldTypeShort,
		fieldTypeLong,
		fieldTypeLongLong,
		fieldTypeInt24,
		fieldTypeYear:
		return "int"
	case fieldTypeNULL:
		return "null"
	case fieldTypeTimestamp,
		fieldTypeDate,
		fieldTypeTime,
		fieldTypeDateTime,
		fieldTypeNewDate:
		return "time"
	case fieldTypeVarChar,
		fieldTypeTinyBLOB,
		fieldTypeMediumBLOB,
		fieldTypeLongBLOB,
		fieldTypeBLOB,
		fieldTypeVarString,
		fieldTypeString:
		return "string"
	case fieldTypeBit:
		return "bit"
	case fieldTypeJSON:
		return "json"
	case fieldTypeEnum:
		return "enum"
	case fieldTypeSet:
		return "set"
	case fieldTypeGeometry:
		return "geometry"
	default:
		return fmt.Sprintf("%x", t)
	}
}
