dp-cli gen --connstr "postgresql://things:things2024@ali4:37054/thingsdb?sslmode=disable" \
--tables=o_ops_dict --model_naming "{{ toUpperCamelCase ( replace . \"o_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"o_\" \"\") }}" \
--module ops-service --api RUDB

