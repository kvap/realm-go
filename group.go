package realm

/*
#include <stdlib.h>
#include "realm-go.h"
*/
import "C"

import (
	"unsafe"
	"errors"
)

type Group struct {
	ptr unsafe.Pointer
}

func (g *Group) Destroy() {
	if g.ptr != nil {
		C.group_destroy(g.ptr)
		g.ptr = nil
	}
}

func (g *Group) HasTable(name string) (bool, error) {
	var hasTable C.bool

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var errString *C.char
	hasTable = C.group_has_table(g.ptr, cName, &errString)
	if errString == nil {
		return bool(hasTable), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return false, errors.New(C.GoString(errString))
}

func (g *Group) GetTableByIdx(idx int) (*Table, error) {
	var errString *C.char
	var t Table
	t.ptr = C.group_get_table_by_idx(g.ptr, C.size_t(idx), &errString)
	if errString == nil {
		return &t, nil
	}
	defer C.free(unsafe.Pointer(errString))
	return nil, errors.New(C.GoString(errString))
}

func (g *Group) GetTableByName(name string) (*Table, error) {
	var errString *C.char
	var t Table

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	t.ptr = C.group_get_table_by_name(g.ptr, cName, &errString)
	if errString == nil {
		return &t, nil
	}
	defer C.free(unsafe.Pointer(errString))
	return nil, errors.New(C.GoString(errString))
}

func (g *Group) AddTable(name string, uniq bool) (*Table, error) {
	var errString *C.char
	var t Table

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	t.ptr = C.group_add_table(g.ptr, cName, C.bool(uniq), &errString)
	if errString == nil {
		return &t, nil
	}
	defer C.free(unsafe.Pointer(errString))
	return nil, errors.New(C.GoString(errString))
}

func (g *Group) GetOrAddTable(name string) (*Table, bool, error) {
	var errString *C.char
	var t Table
	var added C.bool

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	t.ptr = C.group_get_or_add_table(g.ptr, cName, &added, &errString)
	if errString == nil {
		return &t, bool(added), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return nil, false, errors.New(C.GoString(errString))
}
