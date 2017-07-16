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

type Table struct {
	ptr unsafe.Pointer
}

func (t *Table) Destroy() {
	if t.ptr != nil {
		C.table_destroy(t.ptr)
		t.ptr = nil
	}
}

func (t *Table) GetName() (string, error) {
	var errString *C.char

	cName := C.table_get_name(t.ptr, &errString)
	if errString == nil {
		defer C.free(unsafe.Pointer(cName))
		return C.GoString(cName), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return "", errors.New(C.GoString(errString))
}

func (t *Table) ColumnIsNullable(idx int) (bool, error) {
	var errString *C.char

	cNullable := C.table_column_is_nullable(t.ptr, C.size_t(idx), &errString)
	if errString == nil {
		return bool(cNullable), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return false, errors.New(C.GoString(errString))
}

func (t *Table) GetColumnCount() (int, error) {
	var errString *C.char

	cCount := C.table_get_column_count(t.ptr, &errString)
	if errString == nil {
		return int(cCount), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return 0, errors.New(C.GoString(errString))
}

func (t *Table) GetColumnType(idx int) (int, error) {
	var errString *C.char

	cType := C.table_get_column_type(t.ptr, C.size_t(idx), &errString)
	if errString == nil {
		return int(cType), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return TypeInt, errors.New(C.GoString(errString))
}

func (t *Table) GetColumnName(idx int) (string, error) {
	var errString *C.char

	cName := C.table_get_column_name(t.ptr, C.size_t(idx), &errString)
	if errString == nil {
		defer C.free(unsafe.Pointer(cName))
		return C.GoString(cName), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return "", errors.New(C.GoString(errString))
}

func (t *Table) GetColumnIndex(name string) (int, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var errString *C.char
	idx := C.table_get_column_index(t.ptr, cName, &errString)
	if errString == nil {
		return int(idx), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return 0, errors.New(C.GoString(errString))
}

func (t *Table) AddColumn(dataType int, name string, nullable bool) (int, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var errString *C.char
	idx := C.table_add_column(t.ptr, C.int(dataType), cName, C.bool(nullable), &errString)
	if errString == nil {
		return int(idx), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return 0, errors.New(C.GoString(errString))
}

func (t *Table) InsertColumn(idx int, dataType int, name string, nullable bool) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var errString *C.char
	C.table_insert_column(t.ptr, C.size_t(idx), C.int(dataType), cName, C.bool(nullable), &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (t *Table) RemoveColumn(idx int) error {
	var errString *C.char
	C.table_remove_column(t.ptr, C.size_t(idx), &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (t *Table) RenameColumn(idx int, name string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var errString *C.char
	C.table_rename_column(t.ptr, C.size_t(idx), cName, &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (t *Table) AddEmptyRow() (int, error) {
	var errString *C.char
	idx := C.table_add_empty_row(t.ptr, &errString)
	if errString == nil {
		return int(idx), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return 0, errors.New(C.GoString(errString))
}

func (t *Table) InsertEmptyRow(idx int) error {
	var errString *C.char
	C.table_insert_empty_row(t.ptr, C.size_t(idx), &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (t *Table) RemoveRow(idx int) error {
	var errString *C.char
	C.table_remove_row(t.ptr, C.size_t(idx), &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (t *Table) RemoveLastRow() error {
	var errString *C.char
	C.table_remove_last_row(t.ptr, &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (t *Table) MoveLastRowOver(idx int) error {
	var errString *C.char
	C.table_move_last_row_over(t.ptr, C.size_t(idx), &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (t *Table) Clear() error {
	var errString *C.char
	C.table_clear(t.ptr, &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (t *Table) SwapRows(idx1, idx2 int) error {
	var errString *C.char
	C.table_swap_rows(t.ptr, C.size_t(idx1), C.size_t(idx2), &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (t *Table) IsEmpty() (bool, error) {
	var errString *C.char
	cEmpty := C.table_is_empty(t.ptr, &errString)
	if errString == nil {
		return bool(cEmpty), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return false, errors.New(C.GoString(errString))
}

func (t *Table) GetSize() (int, error) {
	var errString *C.char
	cSize := C.table_get_size(t.ptr, &errString)
	if errString == nil {
		return int(cSize), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return 0, errors.New(C.GoString(errString))
}

func (t *Table) SetInt(col, row int, value int64) error {
	var errString *C.char
	C.table_set_int(t.ptr, C.size_t(col), C.size_t(row), C.int64_t(value), &errString)
	if errString == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(errString))
	return errors.New(C.GoString(errString))
}

func (t *Table) GetInt(col, row int) (int64, error) {
	var errString *C.char
	cInt64 := C.table_get_int(t.ptr, C.size_t(col), C.size_t(row), &errString)
	if errString == nil {
		return int64(cInt64), nil
	}
	defer C.free(unsafe.Pointer(errString))
	return 0, errors.New(C.GoString(errString))
}
