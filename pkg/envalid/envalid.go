package envalid

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func GetEnvironmentVariables(envType interface{}) {
	if reflect.ValueOf(envType).Kind() != reflect.Ptr || reflect.ValueOf(envType).Elem().Kind() != reflect.Struct {
		log.Fatalln("invalid input for GetEnvironmentVariables")
	}

	structElement := reflect.ValueOf(envType).Elem()
	for i := 0; i < structElement.NumField(); i++ {
		field := reflect.Indirect(reflect.ValueOf(envType)).Type().Field(i)

		envTag := field.Tag.Get("env")
		rules := make(map[string]string)
		for _, rule := range strings.Split(envTag, ",") {
			cleanedRule := strings.TrimSpace(rule)
			arr := strings.SplitN(cleanedRule, ":", 2)

			if len(arr) > 1 {
				rules[arr[0]] = arr[1]
			} else {
				if arr[0] == "default" {
					log.Fatalln("'default' tag should have a value")
				}

				rules[arr[0]] = ""
			}
		}

		envValue, hasEnv := os.LookupEnv(field.Name)

		if _, hasRule := rules["required"]; hasRule && !hasEnv {
			log.Fatalf("environment property %s is required", field.Name)
		}

		if value, hasRule := rules["default"]; hasRule && !hasEnv {
			if !structElement.Field(i).CanSet() {
				log.Fatalln("struct properties must be addresable")
			}

			if reflectValue, err := convertValueToType(field.Name, field.Type, value); err != nil {
				log.Fatalln(err.Error())
			} else {
				structElement.Field(i).Set(reflectValue)
			}
		}

		if hasEnv {
			if reflectValue, err := convertValueToType(field.Name, field.Type, envValue); err != nil {
				log.Fatalln(err.Error())
			} else {
				structElement.Field(i).Set(reflectValue)
			}
		}

	}
}

func convertValueToType(fieldName string, rt reflect.Type, value string) (reflect.Value, error) {
	switch rt.String() {
	case "string":
		return reflect.ValueOf(value), nil
	case "int":
		if result, err := strconv.ParseInt(value, 10, 0); err == nil {
			return reflect.ValueOf(int(result)), nil
		}
	case "int8":
		if result, err := strconv.ParseInt(value, 10, 8); err == nil {
			return reflect.ValueOf(int8(result)), nil
		}
	case "int16":
		if result, err := strconv.ParseInt(value, 10, 16); err == nil {
			return reflect.ValueOf(int16(result)), nil
		}
	case "int32":
		if result, err := strconv.ParseInt(value, 10, 32); err == nil {
			return reflect.ValueOf(int32(result)), nil
		}
	case "int64":
		if result, err := strconv.ParseInt(value, 10, 32); err == nil {
			return reflect.ValueOf(result), nil
		}
	case "float32":
		if result, err := strconv.ParseFloat(value, 32); err == nil {
			return reflect.ValueOf(float32(result)), nil
		}
	case "float64":
		if result, err := strconv.ParseFloat(value, 32); err == nil {
			return reflect.ValueOf(result), nil
		}
	case "bool":
		if result, err := strconv.ParseBool(value); err == nil {
			return reflect.ValueOf(result), nil
		}
	default:
		return reflect.ValueOf(""), fmt.Errorf("type not supported")
	}

	return reflect.ValueOf(""), fmt.Errorf("error while converting %s of field '%s' to %s", value, fieldName, rt.String())

}
