#include "memberlist.h"
#include "_cgo_export.h"

void
rb_raise2(VALUE exception, const char *str) {
    rb_raise(exception, "%s", str);
}

const char *
rstring_ptr(VALUE str) {
    return RSTRING_PTR(str);
}

long
rstring_len(VALUE str) {
	return RSTRING_LEN(str);
}

int
rstring_lenint(VALUE str) {
	return RSTRING_LENINT(str);
}

void goobj_retain(void *);
void goobj_free(void *);

static const rb_data_type_t go_type = {
    "GoStruct",
    {NULL, goobj_free, NULL},
    0, 0, RUBY_TYPED_FREE_IMMEDIATELY|RUBY_TYPED_WB_PROTECTED
};

VALUE
NewGoStruct(VALUE klass, void *p)
{
    goobj_retain(p);
    return TypedData_Wrap_Struct((klass), &go_type, p);
}

void *
GetGoStruct(VALUE obj)
{
    void *val;
    return TypedData_Get_Struct((obj), void *, &go_type, (val));
}
