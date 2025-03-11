-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- 字典表
CREATE TABLE o_ops_dict (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL,
    dict_type VARCHAR(100) NOT NULL,
    dict_type_name VARCHAR(255) NOT NULL,
    dict_name VARCHAR(255) NOT NULL,
    dict_value VARCHAR(255) NOT NULL,
    sort_order INTEGER NOT NULL DEFAULT 0,
    remarks TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ops_dict IS '字典表';
COMMENT ON COLUMN o_ops_dict.dict_type IS '字典类型';
COMMENT ON COLUMN o_ops_dict.dict_type_name IS '字典类型名称';
COMMENT ON COLUMN o_ops_dict.dict_name IS '字典名称';
COMMENT ON COLUMN o_ops_dict.dict_value IS '字典值';
COMMENT ON COLUMN o_ops_dict.sort_order IS '排序号';
COMMENT ON COLUMN o_ops_dict.remarks IS '备注';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS o_ops_dict;

-- +goose StatementEnd
