#include "device_utils.hh"

String getDeviceID() {
    uint64_t chipid = ESP.getEfuseMac();

    char id[20];
    sprintf(id, "EMS-%04X%08X", (uint16_t)(chipid >> 32), (uint32_t)chipid);

    return String(id);
}