#ifndef MQTT_MANAGER_H
#define MQTT_MANAGER_H

#include <PubSubClient.h>
#include "pzem_manager.hh"

void initMQTT(const char* server);
void mqttLoop();
bool reconnectMQTT();
void sendData(String deviceID, PZEMData& pzemData);
bool isMQTTConnected();

#endif