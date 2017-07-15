#include <cstdio>
#include <realm.hpp>

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

extern "C" void *shared_group_create(char *filename, bool nocreate, char **err) {
	TRY_SET_ERR({
		return new realm::SharedGroup(std::string(filename), nocreate);
	}, err);
	return NULL;
}

extern "C" void shared_group_destroy(void *sg) {
	try {
		delete SGROUP(sg);
	} catch (...) {
	}
}

extern "C" void *shared_group_begin_read(void *sg, char **err) {
	TRY_SET_ERR({
		return GROUP(&SGROUP(sg)->begin_read());
	}, err);
	return NULL;
}

extern "C" void *shared_group_begin_write(void *sg, char **err) {
	TRY_SET_ERR({
		return GROUP(&SGROUP(sg)->begin_write());
	}, err);
	return NULL;
}

extern "C" void shared_group_end_read(void *sg, char **err) {
	TRY_SET_ERR({
		SGROUP(sg)->end_read();
	}, err);
}

extern "C" void shared_group_commit(void *sg, char **err) {
	TRY_SET_ERR({
		SGROUP(sg)->commit();
	}, err);
}

extern "C" void shared_group_rollback(void *sg, char **err) {
	TRY_SET_ERR({
		SGROUP(sg)->rollback();
	}, err);
}

extern "C" bool group_has_table(void *g, char *name, char **err) {
	TRY_SET_ERR({
		return GROUP(g)->has_table(name);
	}, err);
	return false;
}

extern "C" void *group_get_table_by_idx(void *g, size_t idx, char **err) {
	TRY_SET_ERR({
		auto table_ref = GROUP(g)->get_table(idx);
		return new realm::TableRef(table_ref);
	}, err);
	return NULL;
}

extern "C" void *group_get_table_by_name(void *g, char *name, char **err) {
	TRY_SET_ERR({
		auto table_ref = GROUP(g)->get_table(realm::StringData(name));
		return new realm::TableRef(table_ref);
	}, err);
	return NULL;
}

extern "C" void *group_add_table(void *g, char *name, bool uniq, char **err) {
	TRY_SET_ERR({
		auto table_ref = GROUP(g)->add_table(name, uniq);
		return new realm::TableRef(table_ref);
	}, err);
	return NULL;
}

extern "C" void *group_get_or_add_table(void *g, char *name, bool *added, char **err) {
	TRY_SET_ERR({
		auto table_ref = GROUP(g)->get_or_add_table(name, added);
		return new realm::TableRef(table_ref);
	}, err);
	return NULL;
}

extern "C" void group_destroy(void *g) {
	try {
		delete GROUP(g);
	} catch (...) {
	}
}

extern "C" void table_destroy(void *t) {
	try {
		delete TABLEREF(t);
	} catch (...) {
	}
}

