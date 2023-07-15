#include <stdlib.h>

struct foo {
    int id;
    char *info;
    int info_len;
};

struct foo *get_foo();

typedef struct foo Foo;

struct foo *get_foo() {

    struct foo *fooInstance = malloc(sizeof(struct foo));

    fooInstance->id = 1;
    fooInstance->info = "some info";
    fooInstance->info_len = 9;

    return fooInstance;
}