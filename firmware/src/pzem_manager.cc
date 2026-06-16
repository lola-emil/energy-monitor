#include "pzem_manager.hh"
#include "pins.hh"

HardwareSerial pzemSerial(2);
PZEM004Tv30 pzem(pzemSerial, PZEM_RX_PIN, PZEM_TX_PIN);

void initPZEM() {
    pzemSerial.begin(9600, SERIAL_8N1, PZEM_RX_PIN, PZEM_TX_PIN);
}

bool readPZEM(float &v, float &i, float &p, float &e) {
    v = pzem.voltage();
    i = pzem.current();
    p = pzem.power();
    e = pzem.energy();

    return !isnan(v);
}