#include <cstdio>
#include <realm.hpp>

#define BEGIN_EXTERN_C extern "C" {
#define END_EXTERN_C }

#define TRY_SET_ERR(CODE, ERRPTR) \
	try { \
		*(ERRPTR) = NULL; \
		CODE \
	} catch (std::exception const &ex) { \
		*(ERRPTR) = strdup(ex.what()); \
	} catch (...) { \
		*(ERRPTR) = strdup("unknown error"); \
	}

#define SGROUP(PTR) ((realm::SharedGroup*)(PTR))
#define GROUP(PTR) ((realm::Group*)(PTR))
#define TABLEREF(PTR) ((realm::TableRef*)(PTR))
#define TABLE(PTR) (*(realm::TableRef*)(PTR))

static char *string_data_dup(realm::StringData sd) {
	return strndup(sd.data(), sd.size());
}

BEGIN_EXTERN_C

void *shared_group_create(char *filename, bool nocreate, char **err) {
	TRY_SET_ERR({
		return new realm::SharedGroup(std::string(filename), nocreate);
	}, err);
	return NULL;
}

void shared_group_destroy(void *sg) {
	try {
		delete SGROUP(sg);
	} catch (...) {
	}
}

void *shared_group_begin_read(void *sg, char **err) {
	TRY_SET_ERR({
		return GROUP(&SGROUP(sg)->begin_read());
	}, err);
	return NULL;
}

void *shared_group_begin_write(void *sg, char **err) {
	TRY_SET_ERR({
		return GROUP(&SGROUP(sg)->begin_write());
	}, err);
	return NULL;
}

void shared_group_end_read(void *sg, char **err) {
	TRY_SET_ERR({
		SGROUP(sg)->end_read();
	}, err);
}

void shared_group_commit(void *sg, char **err) {
	TRY_SET_ERR({
		SGROUP(sg)->commit();
	}, err);
}

void shared_group_rollback(void *sg, char **err) {
	TRY_SET_ERR({
		SGROUP(sg)->rollback();
	}, err);
}

bool group_has_table(void *g, char *name, char **err) {
	TRY_SET_ERR({
		return GROUP(g)->has_table(name);
	}, err);
	return false;
}

void *group_get_table_by_idx(void *g, size_t idx, char **err) {
	TRY_SET_ERR({
		auto table_ref = GROUP(g)->get_table(idx);
		return new realm::TableRef(table_ref);
	}, err);
	return NULL;
}

void *group_get_table_by_name(void *g, char *name, char **err) {
	TRY_SET_ERR({
		auto table_ref = GROUP(g)->get_table(realm::StringData(name));
		return new realm::TableRef(table_ref);
	}, err);
	return NULL;
}

void *group_add_table(void *g, char *name, bool uniq, char **err) {
	TRY_SET_ERR({
		auto table_ref = GROUP(g)->add_table(name, uniq);
		return new realm::TableRef(table_ref);
	}, err);
	return NULL;
}

void *group_get_or_add_table(void *g, char *name, bool *added, char **err) {
	TRY_SET_ERR({
		auto table_ref = GROUP(g)->get_or_add_table(name, added);
		return new realm::TableRef(table_ref);
	}, err);
	return NULL;
}

void group_destroy(void *g) {
	try {
		delete GROUP(g);
	} catch (...) {
	}
}

char *table_get_name(void *t, char **err) {
	TRY_SET_ERR({
		return string_data_dup(TABLE(t)->get_name());
	}, err);
	return NULL;
}

bool table_column_is_nullable(void *t, size_t idx, char **err) {
	TRY_SET_ERR({
		return TABLE(t)->is_nullable(idx);
	}, err);
	return false;
}

size_t table_get_column_count(void *t, char **err) {
	TRY_SET_ERR({
		return TABLE(t)->get_column_count();
	}, err);
	return false;
}

int table_get_column_type(void *t, size_t idx, char **err) {
	TRY_SET_ERR({
		return TABLE(t)->get_column_type(idx);
	}, err);
	return false;
}

char *table_get_column_name(void *t, size_t idx, char **err) {
	TRY_SET_ERR({
		return string_data_dup(TABLE(t)->get_column_name(idx));
	}, err);
	return NULL;
}

size_t table_get_column_index(void *t, char *name, char **err) {
	TRY_SET_ERR({
		return TABLE(t)->get_column_index(realm::StringData(name));
	}, err);
	return 0;
}

size_t table_add_column(void *t, int type, char *name, bool nullable, char **err) {
	TRY_SET_ERR({
		return TABLE(t)->add_column(realm::DataType(type), realm::StringData(name), nullable);
	}, err);
	return 0;
}

void table_insert_column(void *t, size_t idx, int type, char *name, bool nullable, char **err) {
	TRY_SET_ERR({
		TABLE(t)->insert_column(idx, realm::DataType(type), realm::StringData(name), nullable);
	}, err);
}

void table_remove_column(void *t, size_t idx, char **err) {
	TRY_SET_ERR({
		TABLE(t)->remove_column(idx);
	}, err);
}

void table_rename_column(void *t, size_t idx, char *name, char **err) {
	TRY_SET_ERR({
		TABLE(t)->rename_column(idx, realm::StringData(name));
	}, err);
}

size_t table_add_empty_row(void *t, char **err) {
	TRY_SET_ERR({
		return TABLE(t)->add_empty_row();
	}, err);
	return 0;
}

void table_insert_empty_row(void *t, size_t idx, char **err) {
	TRY_SET_ERR({
		TABLE(t)->add_empty_row(idx);
	}, err);
}

void table_remove_row(void *t, size_t idx, char **err) {
	TRY_SET_ERR({
		TABLE(t)->remove(idx);
	}, err);
}

void table_remove_last_row(void *t, char **err) {
	TRY_SET_ERR({
		TABLE(t)->remove_last();
	}, err);
}

void table_move_last_row_over(void *t, size_t idx, char **err) {
	TRY_SET_ERR({
		TABLE(t)->move_last_over(idx);
	}, err);
}

void table_clear(void *t, char **err) {
	TRY_SET_ERR({
		TABLE(t)->clear();
	}, err);
}

void table_swap_rows(void *t, size_t idx1, size_t idx2, char **err) {
	TRY_SET_ERR({
		TABLE(t)->swap_rows(idx1, idx2);
	}, err);
}

bool table_is_empty(void *t, char **err) {
	TRY_SET_ERR({
		return TABLE(t)->is_empty();
	}, err);
	return false;
}

size_t table_get_size(void *t, char **err) {
	TRY_SET_ERR({
		return TABLE(t)->size();
	}, err);
	return 0;
}

int64_t table_get_int(void *t, size_t col, size_t row, char **err) {
	TRY_SET_ERR({
		return TABLE(t)->get_int(col, row);
	}, err);
	return 0;
}

void table_set_int(void *t, size_t col, size_t row, int_fast64_t value, char **err) {
	TRY_SET_ERR({
		TABLE(t)->set_int(col, row, value);
	}, err);
}

void table_destroy(void *t) {
	try {
		delete TABLEREF(t);
	} catch (...) {
	}
}

END_EXTERN_C
