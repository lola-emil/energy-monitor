#pragma once

#include <PubSubClient.h>
#include <WiFi.h>

#include "EnergySensor.hh"

class NetworkComm {

private:
  const char *ssid = "GlobeAtHome_A0177_2.4";
  const char *password = "letmein123";
  const char *mqttServer = "192.168.254.101";

  unsigned long lastMqttAttempt = 0;
  const unsigned long mqttRetryInterval = 5000; // 5 seconds

  PubSubClient &mqttClient;

  char powerReadingTopic[50];
  char chipID[17];

public:
  NetworkComm(PubSubClient &m) : mqttClient(m) {}

  /**
   * TODO: dapat pa ni butngan ug AP mode
   * para maka input sa bago nga wifi password
   */
  void initConnection() {
    WiFi.begin(ssid, password);

    unsigned long startAttemptTime = millis();
    const unsigned long timeout = 10000; // 10 seconds

    while (WiFi.status() != WL_CONNECTED &&
           millis() - startAttemptTime < timeout) {
      delay(100);
    }

    if (WiFi.status() == WL_CONNECTED) {
      mqttClient.setServer(mqttServer, 1883);
    } else {
      Serial.println("WiFi Failed!");
    }
  }
  void setChipID(uint64_t chipID) {
    snprintf(this->chipID, sizeof(this->chipID), "%04X%08X",
             (uint16_t)(chipID >> 32), (uint32_t)chipID);

    snprintf(powerReadingTopic, sizeof(powerReadingTopic), "sensors/%s/power",
             chipID);
  }

  void publishEnergyData(const SensorData &data) {
    if (!mqttClient.connected())
      return;

    char payload[100];

    /**
     * TODO: kailangan pud labuton ang uban nga value
     *  sa sensor data
     */
    snprintf(payload, sizeof(payload), "{\"voltage\":%.2f,\"power\":%.3f}",
             data.voltage, data.power);

    mqttClient.publish(powerReadingTopic, payload);
  }

  void reconnect() {
    if (mqttClient.connected())
      return;

    if (millis() - lastMqttAttempt < mqttRetryInterval)
      return;

    lastMqttAttempt = millis();

    Serial.print("Attempting MQTT connection...");

    if (mqttClient.connect(chipID)) {
      Serial.println("connected");
    } else {
      Serial.print("failed, rc=");
      Serial.println(mqttClient.state());
    }
  }

  void connectMQTT() {
    if (!mqttClient.connected())
      reconnect();
    mqttClient.loop();
  }
};