#ifndef MQTT_MANAGER_H
#define MQTT_MANAGER_H

#include <PubSubClient.h>

void initMQTT(const char* server);
void mqttLoop();
bool reconnectMQTT();
void sendData(String deviceID, float v, float i, float p, float e);
bool isMQTTConnected();

#endif