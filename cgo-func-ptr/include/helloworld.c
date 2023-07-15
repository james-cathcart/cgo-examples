#include "helloworld.h"

// C can call the function pointer passed back
// to it from Go
int bridge_int_func(intFunc f)
{
	return f();
}

// return the integer 42
int fortytwo()
{
	return 42;
}