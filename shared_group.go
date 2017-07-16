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

func (sg *SharedGroup) Destroy() {
	if sg.ptr != nil {
		C.shared_group_destroy(sg.ptr)
		sg.ptr = nil
	}
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

func (sg *SharedGroup) Read(reader func(*Group)) error {
	g, err := sg.BeginRead()
	if err != nil {
		return err
	}

	reader(g)

	return sg.EndRead()
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
