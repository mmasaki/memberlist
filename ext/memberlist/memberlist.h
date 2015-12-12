#ifndef MEMBERLIST_H
#define MEMBERLIST_H 1

#include "ruby.h"

void rb_raise2(VALUE exception, const char *str);
const char *rstring_ptr(VALUE str);
long rstring_len(VALUE str);
int rstring_lenint(VALUE str);

VALUE NewGoStruct(VALUE klass, void *p);
void *GetGoStruct(VALUE obj);

#endif /* MEMBERLIST_H */
