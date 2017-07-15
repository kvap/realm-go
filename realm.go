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

type SharedGroup struct {
	ptr unsafe.Pointer
}

type Group struct {
	ptr unsafe.Pointer
}

type Table struct {
	ptr unsafe.Pointer
}

func (sg *SharedGroup) Destroy() {
	if sg.ptr != nil {
		C.shared_group_destroy(sg.ptr)
		sg.ptr = nil
	}
}

func (g *Group) Destroy() {
	if g.ptr != nil {
		C.group_destroy(g.ptr)
		g.ptr = nil
	}
}

func (t *Table) Destroy() {
	if t.ptr != nil {
		C.table_destroy(t.ptr)
		t.ptr = nil
	}
}

func NewSharedGroup(fileName string, noCreate bool) (*SharedGroup, error) {
	var sg SharedGroup
	var errString *C.char

	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	sg.ptr = C.shared_group_create(cFileName, C.bool(noCreate), &errString)
	if errString == nil {
		return &sg, nil
	}
	defer C.free(unsafe.Pointer(errString))
	return nil, errors.New(C.GoString(errString))
}

func (sg *SharedGroup) BeginRead() (*Group, error) {
	var errString *C.char
	var g Group
	g.ptr = C.shared_group_begin_read(sg.ptr, &errString)
	if errString == nil {
		return &g, nil
	}
	defer C.free(unsafe.Pointer(errString))
	return nil, errors.New(C.GoString(errString))
}

func (sg *SharedGroup) Write(writer func(*Group) bool) error {
	g, err := sg.BeginWrite()
	if err != nil {
		return err
	}

	if writer(g) {
		return sg.Commit()
	} else {
		return sg.Rollback()
	}
}

func (sg *SharedGroup) BeginWrite() (*Group, error) {
	var errString *C.char
	var g Group
	g.ptr = C.shared_group_begin_write(sg.ptr, &errString)
	if errString == nil {
		return &g, nil
	}
	defer C.free(unsafe.Pointer(errString))
	return nil, errors.New(C.GoString(errString))
}

func (sg *SharedGroup) EndRead() error {
	var errString *C.char
	C.shared_group_end_read(sg.ptr, &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (sg *SharedGroup) Commit() error {
	var errString *C.char
	C.shared_group_commit(sg.ptr, &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (sg *SharedGroup) Rollback() error {
	var errString *C.char
	C.shared_group_rollback(sg.ptr, &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
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
