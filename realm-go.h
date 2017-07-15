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

void table_destroy(void *t);

#endif
