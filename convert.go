package main

import (
	"log"
	"reflect"
)

func Byte_2_Struct(byt []byte, v interface{}) {
	typeTmp := reflect.TypeOf(v)
	//log.Printf("in param -----> %v, type=%v\n", v, typeTmp.Kind())

	if typeTmp.Kind() != reflect.Ptr {
		log.Printf("v must be struct ptr, but it's %v\n", typeTmp.Kind())
		return
	}
	if typeTmp.Elem().Kind() != reflect.Struct {
		log.Printf("v.value must be struct\n")
		return
	}

	valueTmp := reflect.ValueOf(v)
	//log.Printf("in param ---->byt=%v, value= %v \n", hex.EncodeToString(byt), valueTmp)
	curIndex := 0
	// 遍历结构体的所有字段
	elem := typeTmp.Elem()
	for i := 0; i < elem.NumField(); i++ {
		// 得到字段
		field := elem.Field(i)
		//log.Printf("遍历--> %#v\n", field)
		tag := field.Tag.Get("ctype")
		fieldSize := GetCtype_Size(tag)
		//log.Printf("  <--- tag: %v\n", tag)

		switch field.Type.Kind() {
		case reflect.Uint8:
			fVal := GetUInt8(byt, curIndex, fieldSize)
			valueTmp.Elem().FieldByName(field.Name).SetUint(uint64(fVal))
		case reflect.Uint32:
			fVal := GetUInt32(byt, curIndex, fieldSize)
			valueTmp.Elem().FieldByName(field.Name).SetUint(uint64(fVal))
		case reflect.String:
			fVal := GetStr(byt, curIndex, fieldSize)
			valueTmp.Elem().FieldByName(field.Name).SetString(fVal)
		case reflect.Struct:
			// 得到结构体的指针
			//Bytes_2_Struct(bytTmp, valueTmp)
			//fieldValue := reflect.ValueOf(field)
			//log.Printf("字段 %s, 是struct, fieldValue=%#v, field=%v", field.Name, fieldValue, field)
			//tag := field.Tag.Get("ctype")
			//log.Printf("  <--- tag: %v\n", tag)
			//fieldSize := GetCtype_Size(tag)
			bytTmp := byt[curIndex : curIndex+fieldSize]
			//log.Println(field.Elem())
			fVal := reflect.ValueOf(v).Elem().FieldByName(field.Name)
			//log.Printf("[struct] ---> %#v", fVal)
			Byte_2_Struct(bytTmp, fVal.Addr().Interface())
			//PrtStruct(reflect.ValueOf(v).FieldByIndex(i))
		default:
			log.Printf("unkown kind type, %v, type=%v", typeTmp.Kind(), typeTmp)
		}
		curIndex += fieldSize
		////log.Printf("%#v\n, type=%v\n", field, field.Type.Kind())

		//valueTmp1 := reflect.ValueOf(v)

		//Bytes_2_Struct(bytTmp, fieldVal)
	}

}
