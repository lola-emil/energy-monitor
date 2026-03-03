#pragma once

#include <PZEM004Tv30.h>

struct SensorData {
  float voltage;
  float current;
  float power;
  float energy;
  float frequency;
  float pf;
};

class EnergySensor {

private:
  PZEM004Tv30& pzem;

public:
  EnergySensor(PZEM004Tv30& p) : pzem(p) {}

  SensorData getData() {
    return SensorData{ pzem.voltage(), pzem.current(),   pzem.power(),
                      pzem.energy(),  pzem.frequency(), pzem.pf() };
  }

  boolean isSensorDataValid() {
    return isnan(pzem.voltage())
      || isnan(pzem.current())
      || isnan(pzem.power())
      || isnan(pzem.energy())
      || isnan(pzem.frequency())
      || isnan(pzem.pf());
  }
};