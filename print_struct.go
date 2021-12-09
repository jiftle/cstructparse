package main

import (
	"log"
	"reflect"
)

func PrtStruct(v interface{}) {
	typeTmp := reflect.TypeOf(v)
	log.Printf("in param -----> %v, type=%v\n", v, typeTmp.Kind())

	if typeTmp.Kind() != reflect.Ptr {
		log.Printf("v must be struct ptr, but it's %v\n", typeTmp.Kind())
		return
	}
	if typeTmp.Elem().Kind() != reflect.Struct {
		log.Printf("v.value must be struct\n")
		return
	}

	// 遍历结构体的所有字段
	elem := typeTmp.Elem()
	for i := 0; i < elem.NumField(); i++ {
		// 得到字段
		field := elem.Field(i)
		log.Printf("遍历--> %#v\n", field)

		switch field.Type.Kind() {
		case reflect.Uint8:
			log.Println("  <--- Uint8")
			log.Printf("%v", field.Name)
		case reflect.Struct:
			// 得到结构体的指针
			//Bytes_2_Struct(bytTmp, valueTmp)
			fieldValue := reflect.ValueOf(field)
			log.Printf("字段 %s, 是struct, fieldValue=%#v, field=%v", field.Name, fieldValue, field)
			//log.Println(field.Elem())
			fVal := reflect.ValueOf(v).Elem().FieldByName(field.Name)
			log.Printf("[struct] ---> %#v", fVal)
			PrtStruct(fVal.Addr().Interface())
			//PrtStruct(reflect.ValueOf(v).FieldByIndex(i))
		default:
			log.Println(typeTmp.Kind())
		}
		////log.Printf("%#v\n, type=%v\n", field, field.Type.Kind())

		//valueTmp1 := reflect.ValueOf(v)

		//Bytes_2_Struct(bytTmp, fieldVal)
	}

}

func Val_PrtStruct(v interface{}) {
	typeTmp := reflect.TypeOf(v)
	log.Printf("in param -----> %v, type=%v\n", v, typeTmp.Kind())

	// 遍历结构体的所有字段
	elem := typeTmp.Elem()
	for i := 0; i < elem.NumField(); i++ {
		// 得到字段
		field := elem.Field(i)
		log.Printf("遍历--> %#v\n", field)

		switch field.Type.Kind() {
		case reflect.Uint8:
			log.Println("  <--- Uint8")
			log.Printf("%v", field.Name)
		case reflect.Struct:
			// 得到结构体的指针
			//Bytes_2_Struct(bytTmp, valueTmp)
			fieldValue := reflect.ValueOf(field)
			log.Printf("字段 %s, 是struct, fieldValue=%#v, field=%v", field.Name, fieldValue, field)
			//log.Println(field.Elem())
			fVal := reflect.ValueOf(v).Elem().FieldByName(field.Name)
			log.Printf("[struct] ---> %#v", fVal)
			PrtStruct(fVal.Addr())
			//PrtStruct(reflect.ValueOf(v).FieldByIndex(i))
		default:
			log.Println(typeTmp.Kind())
		}
		////log.Printf("%#v\n, type=%v\n", field, field.Type.Kind())

		//valueTmp1 := reflect.ValueOf(v)

		//Bytes_2_Struct(bytTmp, fieldVal)
	}

}
