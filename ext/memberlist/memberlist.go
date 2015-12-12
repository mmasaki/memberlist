package main

/*
#include "memberlist.h"

VALUE memberlist_alloc(VALUE);
VALUE rb_memberlist_join(VALUE self, VALUE member);
VALUE rb_memberlist_members(VALUE self);
*/
import "C"

import(
	"unsafe"
	"github.com/hashicorp/memberlist"
)

func main() {
}

//export memberlist_alloc
func memberlist_alloc(klass C.VALUE) C.VALUE {
	list, err := memberlist.Create(memberlist.DefaultLocalConfig())
	if err != nil {
    rb_raise(C.rb_eRuntimeError, "Failed to create memberlist: %s", err.Error())
	}
	return C.NewGoStruct(klass, unsafe.Pointer(list))
}

//export rb_memberlist_join
func rb_memberlist_join(self C.VALUE, member C.VALUE) C.VALUE {
	list := (*memberlist.Memberlist)(C.GetGoStruct(self))
	name := RbGoString(member)
	n, err := list.Join([]string{name})
	if err != nil {
    rb_raise(C.rb_eRuntimeError, "Failed to join cluster: %", err.Error())
	}
	return INT2NUM(n)
}

//export rb_memberlist_members
func rb_memberlist_members(self C.VALUE) C.VALUE {
	list := (*memberlist.Memberlist)(C.GetGoStruct(self))
	members := list.Members()
	ary := C.rb_ary_new_capa(C.long(len(members)))
	for _, member := range list.Members() {
		C.rb_ary_push(ary, RbString(member.Name))
	}
	return ary
}

var rb_cMemberlist C.VALUE

//export Init_memberlist
func Init_memberlist() {
	rb_cMemberlist = rb_define_class("Memberlist", C.rb_cObject)
	rb_define_alloc_func(rb_cMemberlist, C.memberlist_alloc)

	rb_define_method(rb_cMemberlist, "join", C.rb_memberlist_join, 1)
	rb_define_method(rb_cMemberlist, "members", C.rb_memberlist_members, 0)
}
