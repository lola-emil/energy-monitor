#ifndef PZEM_MANAGER_H
#define PZEM_MANAGER_H

#include <PZEM004Tv30.h>

void initPZEM();
bool readPZEM(float &v, float &i, float &p, float &e);

#endif