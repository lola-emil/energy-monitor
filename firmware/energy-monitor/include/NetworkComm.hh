#pragma once

#include <ArduinoJson.h>
#include <PubSubClient.h>

#include "EnergySensor.hh"

char registeredID[32];
char powerReadingTopic[50];

class NetworkComm {

private:
  const char *ssid = "GlobeAtHome_A0177_2.4";
  const char *password = "ESP32_AP";
  const char *mqttServer = "192.168.254.132";

  unsigned long lastMqttAttempt = 0;
  const unsigned long mqttRetryInterval = 5000; // 5 seconds

  PubSubClient &mqttClient;
  WiFiManager &wifiManager;

  char chipID[17];

  static void mqttCallback(char *topic, byte *payload, unsigned int length) {
    char msg[32];
    memcpy(msg, payload, length);
    msg[length] = '\0';

    Serial.print("Topic: ");
    Serial.println(topic);

    Serial.print("Payload (ID): ");
    Serial.println(msg);

    if (strstr(topic, "/register/success") != nullptr) {
      Serial.println("Registration success received");

      Serial.print("Assigned ID: ");
      Serial.println(msg);

      setPowerReadingTopic(msg);

      // Example: store it
      // strcpy(chipID, msg);  // if chipID is large enough
    }
  }

public:
  NetworkComm(PubSubClient &m, WiFiManager &wm)
      : mqttClient(m), wifiManager(wm) {}

  void initConnection() {
    if (!wifiManager.autoConnect("ESP32_AP")) {
      Serial.println("Failed to connect and hit timeout");
      delay(3000);

      ESP.restart();
    }

    Serial.print("Connected! IP: ");
    Serial.println(WiFi.localIP());

    mqttClient.setServer(mqttServer, 1883);
    mqttClient.setCallback(mqttCallback);
  }

  void publishEnergyData(const SensorData &data) {
    if (!mqttClient.connected())
      return;

    char payload[100];

    snprintf(payload, sizeof(payload),
             "{\"v\":%.2f,"  // voltage
             "\"A\":%.2f,"   // current
             "\"w\":%.2f,"   // power
             "\"e\":%.2f,"   // energy
             "\"f\":%.2f,"   // frequency
             "\"pf\":%.2f}", // power factor
             data.voltage, data.current, data.power, data.energy,
             data.frequency, data.pf);

    Serial.printf("Topic: %s\n", powerReadingTopic);

    bool ok = mqttClient.publish(powerReadingTopic, payload);

    if (ok) {
      Serial.println("Na send ang payload");
    } else {
      Serial.println("Wala na send ang payload");
    }
  }

  static void setPowerReadingTopic(char id[32]) {
    snprintf(powerReadingTopic, sizeof(powerReadingTopic), "device/%s/sensor",
             id);
  }

  void registerDevice() {

    bool subOk = mqttClient.subscribe("device/ESP-MADAFAK/register/success");

    Serial.print("Subscribe result: ");
    Serial.println(subOk);

    char payload[50];
    const char *regCode = "REG456";

    snprintf(payload, sizeof(payload), "{\"s\":\"%s\",\"c\":\"%s\"}",
             "ESP-MADAFAK", regCode);

    bool ok = mqttClient.publish("device/register", payload);

    Serial.print("Payload: ");
    Serial.println(payload);

    if (ok) {
      Serial.println("Publish OK");
    } else {
      Serial.println("Publish FAILED");
    }
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
      registerDevice();
    } else {
      Serial.print("failed, rc=");
      Serial.println(mqttClient.state());
    }
  }

  bool connectMQTT() {
    if (!mqttClient.connected()) {
      reconnect();
      return false;
    }
    mqttClient.loop();
    return true;
  }
};