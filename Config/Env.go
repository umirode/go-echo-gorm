package Config

import (
	"os"
	"strconv"
)

func GetEnv(key string, valueType string) interface{} {
	value := os.Getenv(key)

	result := *new(interface{})
	switch valueType {
	case "bool":
		result, _ = strconv.ParseBool(value)
		break
	case "int":
		result, _ = strconv.Atoi(value)
		break
	case "int64":
		result, _ = strconv.ParseInt(value, 10, 64)
		break
	case "uint64":
		result, _ = strconv.ParseUint(value, 10, 0)
		break
	case "uint":
		result, _ = strconv.ParseUint(value, 10, 0)
		result = uint(result.(uint64))
		break
	case "string":
		result = value
		break
	}

	return result
}
