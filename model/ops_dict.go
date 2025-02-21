package model

import (
	"database/sql"
	"github.com/dapr-platform/common"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = common.LocalTime{}
)

/*
DB Table Details
-------------------------------------


Table: o_ops_dict
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] created_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] created_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] updated_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] updated_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] dict_type                                      VARCHAR(100)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
[ 6] dict_code                                      VARCHAR(100)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
[ 7] dict_name                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] dict_value                                     VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 9] sort_order                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[10] remarks                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "mmmhYKlmvYxuihvJFvGHLeNIt",    "created_by": "dTvdpSfIoDdZYdoXSmvfgZkwn",    "created_time": 73,    "updated_by": "OIAvkBZOsZHIwAKYjSAthjBpC",    "updated_time": 78,    "dict_type": "utJBWWIDbtfsRNEiiVMDagKen",    "dict_code": "kFqOOvtIRMMCJYybxhjvecNGE",    "dict_name": "ZjqKyWpxLoDIfqxMTVELgmyjk",    "dict_value": "fdNHXsPOoLJdsPtGdfVODlBAc",    "sort_order": 81,    "remarks": "SdyONVUwsoCvSGhYbgAvRHcRP"}



*/

var (
	Ops_dict_FIELD_NAME_id = "id"

	Ops_dict_FIELD_NAME_created_by = "created_by"

	Ops_dict_FIELD_NAME_created_time = "created_time"

	Ops_dict_FIELD_NAME_updated_by = "updated_by"

	Ops_dict_FIELD_NAME_updated_time = "updated_time"

	Ops_dict_FIELD_NAME_dict_type = "dict_type"

	Ops_dict_FIELD_NAME_dict_code = "dict_code"

	Ops_dict_FIELD_NAME_dict_name = "dict_name"

	Ops_dict_FIELD_NAME_dict_value = "dict_value"

	Ops_dict_FIELD_NAME_sort_order = "sort_order"

	Ops_dict_FIELD_NAME_remarks = "remarks"
)

// Ops_dict struct is a row record of the o_ops_dict table in the  database
type Ops_dict struct {
	ID string `json:"id"` //id

	CreatedBy string `json:"created_by"` //created_by

	CreatedTime common.LocalTime `json:"created_time"` //created_time

	UpdatedBy string `json:"updated_by"` //updated_by

	UpdatedTime common.LocalTime `json:"updated_time"` //updated_time

	DictType string `json:"dict_type"` //字典类型

	DictCode string `json:"dict_code"` //字典编码

	DictName string `json:"dict_name"` //字典名称

	DictValue string `json:"dict_value"` //字典值

	SortOrder int32 `json:"sort_order"` //排序号

	Remarks string `json:"remarks"` //备注

}

var Ops_dictTableInfo = &TableInfo{
	Name: "o_ops_dict",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "created_by",
			Comment:            `created_by`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "CreatedBy",
			GoFieldType:        "string",
			JSONFieldName:      "created_by",
			ProtobufFieldName:  "created_by",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "created_time",
			Comment:            `created_time`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreatedTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "created_time",
			ProtobufFieldName:  "created_time",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "updated_by",
			Comment:            `updated_by`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "UpdatedBy",
			GoFieldType:        "string",
			JSONFieldName:      "updated_by",
			ProtobufFieldName:  "updated_by",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "updated_time",
			Comment:            `updated_time`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "updated_time",
			ProtobufFieldName:  "updated_time",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "dict_type",
			Comment:            `字典类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       100,
			GoFieldName:        "DictType",
			GoFieldType:        "string",
			JSONFieldName:      "dict_type",
			ProtobufFieldName:  "dict_type",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "dict_code",
			Comment:            `字典编码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       100,
			GoFieldName:        "DictCode",
			GoFieldType:        "string",
			JSONFieldName:      "dict_code",
			ProtobufFieldName:  "dict_code",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "dict_name",
			Comment:            `字典名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "DictName",
			GoFieldType:        "string",
			JSONFieldName:      "dict_name",
			ProtobufFieldName:  "dict_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "dict_value",
			Comment:            `字典值`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "DictValue",
			GoFieldType:        "string",
			JSONFieldName:      "dict_value",
			ProtobufFieldName:  "dict_value",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "sort_order",
			Comment:            `排序号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SortOrder",
			GoFieldType:        "int32",
			JSONFieldName:      "sort_order",
			ProtobufFieldName:  "sort_order",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "remarks",
			Comment:            `备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Remarks",
			GoFieldType:        "string",
			JSONFieldName:      "remarks",
			ProtobufFieldName:  "remarks",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *Ops_dict) TableName() string {
	return "o_ops_dict"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *Ops_dict) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *Ops_dict) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *Ops_dict) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *Ops_dict) TableInfo() *TableInfo {
	return Ops_dictTableInfo
}
