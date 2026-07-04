#include "pzem_manager.hh"
#include "pins.hh"

HardwareSerial pzemSerial(2);
PZEM004Tv30 pzem(pzemSerial, PZEM_RX_PIN, PZEM_TX_PIN);

void initPZEM() {
    pzemSerial.begin(9600, SERIAL_8N1, PZEM_RX_PIN, PZEM_TX_PIN);
}

void readPZEM(PZEMData& data) {
    data.voltage = pzem.voltage();
    data.current = pzem.current();
    data.power = pzem.power();
    data.energy = pzem.energy();
    data.frequency = pzem.frequency();

}

bool isReadingValid(PZEMData& data) {
    return !isnan(data.voltage);
}