#include "pins.hh"
#include <Arduino.h>

#include "button_handler.hh"
#include "device_utils.hh"
#include "mqtt_manager.hh"
#include "pzem_manager.hh"
#include "wifi_manager.hh"

const char *mqtt_server = "192.168.194.181";

unsigned long lastRead = 0;
unsigned long lastReconnectAttempt = 0;
String DEVICE_ID;

void setup() {
  Serial.begin(115200);
  delay(300);

  DEVICE_ID = getDeviceID();
  Serial.println("Device ID: " + DEVICE_ID);

  pinMode(LED_MQTT_PIN, OUTPUT);
  pinMode(LED_POWER_PIN, OUTPUT);

  digitalWrite(LED_POWER_PIN, LED_ON);
  digitalWrite(LED_MQTT_PIN, LED_OFF);

  initButton();
  initPZEM();
  setupWiFi();
  initMQTT(mqtt_server);
}

void loop() {
  if (!isMQTTConnected()) {
    unsigned long now = millis();

    if (now - lastReconnectAttempt > 3000) {
      lastReconnectAttempt = now;

      if (reconnectMQTT()) {
        lastReconnectAttempt = 0;
      }
    }
  } else {
    mqttLoop();
  }

  handleButton();

  if (millis() - lastRead >= 2000) {
    lastRead = millis();

    float v, i, p, e;

    if (readPZEM(v, i, p, e)) {
      sendData(DEVICE_ID, v, i, p, e);

      digitalWrite(LED_MQTT_PIN, LED_ON);
      delay(50);
      digitalWrite(LED_MQTT_PIN, LED_OFF);
    } else {
      Serial.println("PZEM read error");
    }
  }
}