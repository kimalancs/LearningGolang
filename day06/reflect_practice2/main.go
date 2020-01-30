package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini配置文件解析器
// 使用反射从config.ini文件中获取字段信息传入struct结构体

// MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig Redis配置结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test bool 		`ini:"test"`
}

// Config config
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 参数的校验
	// 传进来的data参数必须是指针类型，函数中要修改变量
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer")
		return
	}
	// 传进来的data参数必须是结构体类型的指针，配置文件中的键值对要赋值给结构体字段
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a struct point")
		return
	}
	// 读取文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\n")
	// 一行一行读数据
	var structName string
	for index, line := range lineSlice {
		// 去掉字符串首尾的空格
		line = strings.TrimSpace(line)
		// 如果是空行就跳过
		if len(line) == 0 {
			continue
		}
		// 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 如果方括号开头的是节Section
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' && line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", index+1)
				// index+1，是因为ini文件中行数从0开始，而遍历获取每行内容时，index从0开始
				return
			}
			// 把首尾的[]去掉，取到中间的内容，再把首尾的空格去掉，拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}
			// 根据sectionName，去data里根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				// 找到对应的嵌套结构体，把字段名记下来
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 不是方括号开头的就是键值对
			// 以等号分割这一行，等号左边是key，右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}
			indexof := strings.Index(line, "=")
			key := strings.TrimSpace(line[:indexof])
			value := strings.TrimSpace(line[indexof+1:])
			// 根据structName去data里面把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			structValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			structType := structValue.Type()                // 拿到嵌套结构体的类型信息
			if structType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			// 遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < structValue.NumField(); i++ {
				field := structType.Field(i) // tag信息存在类型信息中
				if field.Tag.Get("ini") == key {
					// 找到对应的字段，存起来
					fieldName = field.Name
				}
			}

			// 如果key等于tag，把value赋值给这个字段
			// 根据fieldName取出这个字段
			if len(fieldName) == 0 {
				// 结构体中找不到对应的字段就跳过
				continue
			}
			fieldObj := structValue.FieldByName(fieldName)
			// 对其赋值
			fmt.Println(fieldName, fieldObj.Type().Kind(), fieldObj)
			switch fieldObj.Type().Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", index+1)
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", index+1)
					return
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", index+1)
				}
				fieldObj.SetFloat(valueFloat)
			}
		}
	}

	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	fmt.Println(cfg, err)
}
