#pragma once

#include <PubSubClient.h>
#include <WiFi.h>
#include <WiFiManager.h>

#include "EnergySensor.hh"

class NetworkComm {

private:
  const char *ssid = "GlobeAtHome_A0177_2.4";
  const char *password = "Octocat.2024..";
  const char *mqttServer = "192.168.254.114";

  unsigned long lastMqttAttempt = 0;
  const unsigned long mqttRetryInterval = 5000; // 5 seconds

  PubSubClient &mqttClient;
  WiFiManager &wifiManager;

  char powerReadingTopic[50];
  char chipID[17];

public:
  NetworkComm(PubSubClient &m, WiFiManager &wm)
      : mqttClient(m), wifiManager(wm) {}

  /**
   * TODO: dapat pa ni butngan ug AP mode
   * para maka input sa bago nga wifi password
   */
  void initConnection() {
    if (!wifiManager.autoConnect(
            "ESP32_AP")) { 
      Serial.println("Failed to connect and hit timeout");
      delay(3000);

      ESP.restart();
    }

    Serial.print("Connected! IP: ");
    Serial.println(WiFi.localIP());

    mqttClient.setServer(mqttServer, 1883);
  }

  void setChipID(uint64_t chipID) {
    snprintf(powerReadingTopic, sizeof(powerReadingTopic), "sensors/%04X/power",
             (uint16_t)(chipID >> 32));
  }

  void publishEnergyData(const SensorData &data) {
    if (!mqttClient.connected())
      return;

    char payload[100];

    snprintf(payload, sizeof(payload),
             "{\"voltage\":%.2f,"
             "\"current\":%.3f,"
             "\"power\":%.3f,"
             "\"energy\":%.3f,"
             "\"frequency\":%.2f,"
             "\"pf\":%.3f}",
             data.voltage, data.current, data.power, data.energy,
             data.frequency, data.pf);

    Serial.printf("Topic: %s", powerReadingTopic);

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