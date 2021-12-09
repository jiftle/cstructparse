package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func Bytes_2_Struct(bytTmp []byte, v interface{}) {
	//curIndex := 0

	typeTmp := reflect.TypeOf(v)
	valueTmp := reflect.ValueOf(v)

	log.Printf("type=%v, elem=%v, value=%v\n", typeTmp, typeTmp.Elem(), valueTmp)
	count := typeTmp.Elem().NumField()
	fmt.Println(count)
	for i := 0; i < count; i++ {
		field := typeTmp.Elem().Field(i)
		log.Printf("%#v\n", field)
		////log.Printf("%#v\n, type=%v\n", field, field.Type.Kind())

		//valueTmp1 := reflect.ValueOf(v)

		//Bytes_2_Struct(bytTmp, fieldVal)
	}

	switch typeTmp.Kind() {
	case reflect.Array:
		log.Println("Array")
	case reflect.Chan:
		log.Println("Chan")
	case reflect.Func:
		log.Println("Func")
	case reflect.Interface:
		log.Println("Interface")
	case reflect.Map:
		log.Println("Map")
	case reflect.Ptr:
		log.Println("Ptr")
	case reflect.Slice:
		log.Println("Slice")
	case reflect.Struct:
		log.Println("struct")
		//Bytes_2_Struct(bytTmp, valueTmp)
	default:
		log.Println(typeTmp.Kind())
	}

	//log.Printf("typeTmp=%#v, valueTmp=%#v\n", typeTmp.Kind(), valueTmp)
	//count := typeTmp.NumField()
	//for i := 0; i < count; i++ {
	//field := typeTmp.Field(i)

	////field.Type.FieldByIndex(i)

	//log.Printf("%#v\n, type=%v\n", field, field.Type.Kind())
	//tag := field.Tag.Get("ctype")
	//log.Printf("tag: %v\n", tag)

	//fieldSize := GetCtype_Size(tag)

	//// val
	//fieldKind := field.Type.Kind()
	////fieldKind := valueTmp.Kind()
	//log.Printf("kind: %v", fieldKind)
	//switch fieldKind {
	//case reflect.Uint8:
	//fieldVal := GetUInt8(bytTmp, curIndex, fieldSize)
	//log.Printf("  |--->,uint8, i=%v, val=%v", i, fieldVal)
	//valueTmp.Elem().FieldByName(field.Name).SetUint(uint64(fieldVal))
	//case reflect.Struct:
	//log.Println("---------- struct ------------")
	//case reflect.String:
	//fieldVal := GetStr(bytTmp, curIndex, fieldSize)
	//valueTmp.Elem().FieldByName(field.Name).SetString(fieldVal)
	//}

	//curIndex += fieldSize
	//}

}

func GetVal(bytTmp []byte) {
	tmp := UK_CosDEVINFO{}

	curIndex := 0

	typeTmp := reflect.TypeOf(tmp)
	valueTmp := reflect.ValueOf(&tmp)
	count := typeTmp.NumField()
	for i := 0; i < count; i++ {
		field := typeTmp.Field(i)

		//field.Type.FieldByIndex(i)

		log.Printf("%#v\n, type=%v\n", field, field.Type.Kind())
		tag := field.Tag.Get("ctype")
		log.Printf("tag: %v\n", tag)

		fieldSize := GetCtype_Size(tag)

		// val
		fieldKind := field.Type.Kind()
		//fieldKind := valueTmp.Kind()
		log.Printf("kind: %v", fieldKind)
		switch fieldKind {
		//case reflect.Int:
		//fieldVal := GetUInt8(bytTmp, curIndex, fieldSize)
		//log.Printf("  |---> val=%v", fieldVal)
		//valueTmp.SetInt32(int64(fieldVal))
		case reflect.Uint8:
			fieldVal := GetUInt8(bytTmp, curIndex, fieldSize)
			log.Printf("  |--->,uint8, i=%v, val=%v", i, fieldVal)
			valueTmp.Elem().FieldByName(field.Name).SetUint(uint64(fieldVal))
		case reflect.Struct:
			log.Println("---------- struct ------------")

		case reflect.String:
			fieldVal := GetStr(bytTmp, curIndex, fieldSize)
			valueTmp.Elem().FieldByName(field.Name).SetString(fieldVal)
		}

		curIndex += fieldSize
	}

	//valueTmp := reflect.ValueOf(tmp)
	//k := valueTmp.Kind()
	//switch k {
	//case reflect.Int:
	//valueTmp.SetInt(1)
	//}

	log.Printf("---------v: %#v\n", tmp)
}

func GetCtype_Size(tag string) int {
	arTag := strings.Split(tag, ",")
	for _, v := range arTag {
		arTmp := strings.Split(v, "=")
		//log.Println(arTmp)
		if arTmp[0] == "size" {
			n, _ := strconv.Atoi(arTmp[1])
			return n
		}
	}

	return 0
}

func GetCstrLen(byt []byte) int {
	j := 0
	for i := len(byt) - 1; i >= 0; i-- {
		j++
		if byt[i] != 0 {
			break
		}
	}
	cl := len(byt) - j + 1
	return cl
}

func GetStr(bytBuf []byte, start, nlen int) string {
	var ss string
	bytTmp := bytBuf[start : start+nlen]
	endIndex := start + GetCstrLen(bytTmp)
	bytTmp = bytBuf[start:endIndex]
	ss = string(bytTmp)
	return ss
}

func GetUInt8(bytBuf []byte, start, nlen int) uint8 {
	var n uint8
	bytTmp := bytBuf[start : start+nlen]
	n = uint8(bytTmp[0])
	return n
}

func GetUInt32(bytBuf []byte, start, nlen int) uint32 {
	var n uint32
	bytTmp := bytBuf[start : start+nlen]
	n = binary.LittleEndian.Uint32(bytTmp)
	return n
}

//反射修改值
func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		// 反射修改值
		v.SetFloat(6.9)
		fmt.Println("a is ", v.Float())
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}
}
