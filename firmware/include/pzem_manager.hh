#ifndef PZEM_MANAGER_H
#define PZEM_MANAGER_H

#include <PZEM004Tv30.h>

struct PZEMData {
    float voltage;
    float current;
    float power;
    float energy;
};

void initPZEM();
bool readPZEM(PZEMData& data);

#endif