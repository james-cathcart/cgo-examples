#ifndef _HELLOWORLD_H_
#define _HELLOWORLD_H_

#include <stdio.h>
#include <stdlib.h>

// create typedef for function pointer to pass back to go
typedef int (*intFunc) ();

int bridge_int_func(intFunc f);

int fortytwo();

#endif