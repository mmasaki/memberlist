package main

/*
#include "memberlist.h"
*/
import "C"

import (
	"unsafe"
	"fmt"
)

var objects = make(map[interface{}]int)

//export goobj_retain
func goobj_retain(obj unsafe.Pointer) {
	objects[obj]++
}

//export goobj_free
func goobj_free(obj unsafe.Pointer) {
	objects[obj]--
	if objects[obj] <= 0 {
		delete(objects, obj)
	}
}

func GOSTRING_PTR(str string) *C.char {
	bytes := *(*[]byte)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(&bytes[0]))
}

func GOSTRING_LEN(str string) C.long {
	return C.long(len(str))
}

func rb_define_alloc_func(klass C.VALUE, fun unsafe.Pointer) {
	C.rb_define_alloc_func(klass, (*[0]byte)(fun))
}

func rb_define_method(klass C.VALUE, name string, fun unsafe.Pointer, args int) {
	cname := GOSTRING_PTR(name)
	C.rb_define_method(klass, cname, (*[0]byte)(fun), C.int(args))
}

func rb_define_singleton_method(klass C.VALUE, name string, fun unsafe.Pointer, args int) {
	cname := GOSTRING_PTR(name)
	C.rb_define_singleton_method(klass, cname, (*[0]byte)(fun), C.int(args))
}

func LONG2NUM(n C.long) C.VALUE {
	return C.rb_long2num_inline(n)
}

func NUM2LONG(n C.VALUE) C.long {
	return C.rb_num2long(n)
}

func RSTRING_PTR(str C.VALUE) *C.char {
	return C.rstring_ptr(str)
}

func RSTRING_LEN(str C.VALUE) C.long {
	return C.rstring_len(str)
}

func RSTRING_LENINT(str C.VALUE) C.int {
	return C.rstring_lenint(str)
}

func RbGoString(str C.VALUE) string {
	C.rb_string_value(&str)
	return C.GoStringN(RSTRING_PTR(str), RSTRING_LENINT(str))
}

func RbBytes(bytes []byte) C.VALUE {
	if len(bytes) == 0 {
		return C.rb_str_new(nil, C.long(0))
	}
	cstr := (*C.char)(unsafe.Pointer(&bytes[0]))
	return C.rb_str_new(cstr, C.long(len(bytes)))
}

func RbString(str string) C.VALUE {
	if len(str) == 0 {
		return C.rb_utf8_str_new(nil, C.long(0))
	}
	return C.rb_utf8_str_new(GOSTRING_PTR(str), GOSTRING_LEN(str))
}

func rb_define_class(name string, parent C.VALUE) C.VALUE {
	return C.rb_define_class(GOSTRING_PTR(name), parent)
}

func rb_raise(exc C.VALUE, format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	C.rb_raise2(exc, GOSTRING_PTR(str))
}

func INT2NUM(n int) C.VALUE {
	return C.rb_int2inum(C.long(n))
}

func INT64toNUM(n int64) C.VALUE {
	return C.rb_ll2inum(C.longlong(n))
}

func StrSlice2RbArray(slice []string) C.VALUE {
	ary := C.rb_ary_new_capa(C.long(len(slice)))
	for _, val := range slice {
		C.rb_ary_push(ary, RbString(val))
	}
	return ary
}
