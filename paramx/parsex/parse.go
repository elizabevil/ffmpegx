package parsex

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"reflect"
	"strconv"
	"strings"
)

type DefaultParamParser struct {
}

var DefaultParser = DefaultParamParser{}

func (d DefaultParamParser) ParamParse(input any) typex.Args {
	if input == nil {
		return make(typex.Args, 0)
	}
	of := reflect.ValueOf(input)
	switch of.Kind() {
	case reflect.Struct:
		tx := of.Type()
		args := make(typex.Args, 0, of.NumField()>>2)
		for i := of.NumField() - 1; i > -1; i-- {
			value := of.Field(i)
			isZero := value.IsZero()
			if isZero { //skip: zero and (no default value)
				continue
			}
			tagFlag := tx.Field(i).Tag.Get("flag")  // get flag
			tagIndex := strings.Split(tagFlag, ",") // has default value ?
			lenTag := len(tagIndex)
			if lenTag > 1 {
				tagFlag = tagIndex[0] //just:tagFlag: -v,1 ==> -v
			}
			switch value.Kind() {
			case reflect.Bool:
				if value.Bool() {
					args = append(args, tagFlag)
				}
			case reflect.Float64, reflect.Float32:
				args = append(args, tagFlag, strconv.FormatFloat(value.Float(), 'f', -1, 32))
			case reflect.String:
				if len(tagFlag) > 0 {
					args = append(args, tagFlag, value.String())
				} else {
					args = append(args, value.String())
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if isZero {
					args = append(args, tagIndex...)
				} else {
					args = append(args, tagFlag, strconv.FormatInt(value.Int(), 10))
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if isZero {
					args = append(args, tagIndex...)
				} else {
					args = append(args, tagFlag, strconv.FormatInt(int64(value.Uint()), 10))
				}
			case reflect.Array, reflect.Slice:
				l := value.Len()
				for i := 0; i < l; i++ {
					itemType, b := d.ParamItemType(value.Index(i))
					if b {
						if len(tagFlag) > 0 {
							args = append(args, tagFlag, itemType)
						} else {
							// outputs
							args = append(args, itemType)
						}
					} else {
						args = append(args, tagFlag+itemType)
					}
				}
			case reflect.Map:
				args = append(args, tagFlag)
				mapRange := value.MapRange()
				if mapRange.Next() {
					key, _ := d.ParamItemType(mapRange.Key())
					vv, _ := d.ParamItemType(mapRange.Value())
					args = append(args, fmt.Sprintf("%v=%v", key, vv))
				}
			case reflect.Struct:
				fmt.Println(value.Interface())
				if c, ok := value.Interface().(typex.StreamSpecifier); ok {
					args = append(args, fmt.Sprintf("%v:%v %v", tagFlag, c.K, c.V))
				}
			case reflect.Interface:
				if value.IsZero() {
					args = append(args, tagIndex...)
					continue
				}
				if c, ok := value.Interface().(typex.StreamSpecifier); ok {
					args = append(args, fmt.Sprintf("%v:%v %v", tagFlag, c.K, c.V))
				}
			case reflect.Pointer:
				valu := value.Elem().Interface()
				switch rec := valu.(type) {
				case typex.Bool:
					args = append(args, tagFlag, strconv.FormatInt(int64(rec), 10))
				case typex.Number:
					args = append(args, tagFlag, strconv.FormatInt(int64(rec), 10))
				case typex.UNumber:
					args = append(args, tagFlag, strconv.FormatInt(int64(rec), 10))
				case typex.Flt:
					args = append(args, tagFlag, strconv.FormatFloat(float64(rec), 'f', -1, 32))
				}
			default:

			}
		}
		return args
	default:
		return nil
	}
}

// value --> args（string）, blank(true:-c copy ;false: -c:v)

func (d DefaultParamParser) ParamItemType(of reflect.Value) (string, bool) {
	switch of.Kind() {
	case reflect.Bool:
		if of.Bool() {
			return "", false
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(of.Int(), 10), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatInt(int64(of.Uint()), 10), true
	case reflect.Array, reflect.Slice:
		l := of.Len()
		args := make([]string, 0, l)
		for i := 0; i < l; i++ {
			args = append(args, of.Index(i).String())
		}
		return strings.Join(args, " "), true
	case reflect.String:
		return of.String(), true
	case reflect.Map:
		args := make([]string, 0, of.Len())
		mapRange := of.MapRange()
		if mapRange.Next() {
			args = append(args, fmt.Sprintf("%s=\"%s\"", mapRange.Key().String(), mapRange.Value().String()))
		}
		return strings.Join(args, " "), true
	case reflect.Struct:
		if c, ok := of.Interface().(typex.StreamSpecifier); ok {
			//return ":" + c.K + " " + c.V, false
			return fmt.Sprintf(":%v %v", c.K, c.V), false
		}
	case reflect.Interface:
		return fmt.Sprintf("%v", of), true
	default:
	}
	return "", false
}
