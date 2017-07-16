#ifndef REALM_GO_H
#define REALM_GO_H

#include <stdbool.h>

// Most of these methods will allocate a C string with error description.
// The caller must free that string later.

// Return *SharedGroup. The caller must free that later.
void *shared_group_create(char *filename, bool nocreate, char **err);
void shared_group_destroy(void *sg);

// Return *Group. The caller must NOT free that later.
void *shared_group_begin_read(void *sg, char **err);
void *shared_group_begin_write(void *sg, char **err);

void shared_group_end_read(void *sg, char **err);
void shared_group_commit(void *sg, char **err);
void shared_group_rollback(void *sg, char **err);

bool group_has_table(void *g, char *name, char **err);
// Return *TableRef. The caller must free that later.
void *group_get_table_by_idx(void *g, size_t idx, char **err);
void *group_get_table_by_name(void *g, char *name, char **err);
void *group_add_table(void *g, char *name, bool uniq, char **err);
void *group_get_or_add_table(void *g, char *name, bool *added, char **err);
void group_destroy(void *g);

// Returns *char. The caller must free that later.
char *table_get_name(void *t, char **err);
bool table_column_is_nullable(void *t, size_t idx, char **err);
size_t table_get_column_count(void *t, char **err);
int table_get_column_type(void *t, size_t idx, char **err);
// Returns *char. The caller must free that later.
char *table_get_column_name(void *t, size_t idx, char **err);
size_t table_get_column_index(void *t, char *name, char **err);
size_t table_add_column(void *t, int type, char *name, bool nullable, char **err);
void table_insert_column(void *t, size_t idx, int type, char *name, bool nullable, char **err);
void table_remove_column(void *t, size_t idx, char **err);
void table_rename_column(void *t, size_t idx, char *name, char **err);
size_t table_add_empty_row(void *t, char **err);
void table_insert_empty_row(void *t, size_t idx, char **err);
void table_remove_row(void *t, size_t idx, char **err);
void table_remove_last_row(void *t, char **err);
void table_move_last_row_over(void *t, size_t idx, char **err);
void table_clear(void *t, char **err);
void table_swap_rows(void *t, size_t idx1, size_t idx2, char **err);
bool table_is_empty(void *t, char **err);
size_t table_get_size(void *t, char **err);

int64_t table_get_int(void *t, size_t col, size_t row, char **err);
void table_set_int(void *t, size_t col, size_t row, int64_t value, char **err);

void table_destroy(void *t);

#endif
