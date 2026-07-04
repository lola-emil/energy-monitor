#ifndef PZEM_MANAGER_H
#define PZEM_MANAGER_H

#include <PZEM004Tv30.h>

struct PZEMData {
    float voltage;
    float current;
    float power;
    float energy;
    float frequency;
};

void initPZEM();
void readPZEM(PZEMData& data);
bool isReadingValid(PZEMData& data);

#endif