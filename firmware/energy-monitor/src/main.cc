#include <PZEM004Tv30.h>
#include <Wire.h>

#include <PubSubClient.h>
#include <WiFi.h>
#include <WiFiManager.h>

#include "EnergySensor.hh"
#include "NetworkComm.hh"

#if defined(ESP32)
PZEM004Tv30 pzem(Serial2, RX2, TX2);
#else
PZEM004Tv30 pzem(Serial2);
#endif
EnergySensor sensor(pzem);

constexpr uint16_t SCREEN_WIDTH = 128;
constexpr uint16_t SCREEN_HEIGHT = 64;
constexpr int8_t OLED_RESET = -1;

WiFiClient wifi;
WiFiManager wm;
PubSubClient client(wifi);
NetworkComm netc(client, wm);

static unsigned long lastReadMillis = 0;

// LED CONFIG
static unsigned long ledOnMillis = 0;
constexpr unsigned long LED_DURATION = 100;
constexpr uint8_t GREEN_LIGHT = 32;
constexpr uint8_t YELLOW_LIGHT = 33;

// PUSH BUTTON CCONFIG
constexpr uint16_t LONG_PRESS_TIME = 5000;
constexpr uint8_t BUTTON_PIN = 14;
static unsigned long buttonPressStart = 0;

bool buttonWasPressed = false;

void setup() {
  Serial.begin(115200);

  pinMode(BUTTON_PIN, INPUT_PULLUP);

  pinMode(GREEN_LIGHT, OUTPUT);
  pinMode(YELLOW_LIGHT, OUTPUT);

  digitalWrite(YELLOW_LIGHT, HIGH);

  Serial2.begin(9600, SERIAL_8N1, RX2, TX2);
  netc.initConnection();

  uint64_t macAddress = ESP.getEfuseMac(); // Gets the unique MAC address
  netc.setChipID(macAddress);

  digitalWrite(YELLOW_LIGHT, LOW);
}

void loop() {
  netc.connectMQTT();

  bool buttonPressed = digitalRead(BUTTON_PIN) == LOW;

  if (buttonPressed) {
    if (!buttonWasPressed) {

      buttonPressStart = millis();
      buttonWasPressed = true;
    } else {
      if (millis() - buttonPressStart >= LONG_PRESS_TIME) {
        Serial.println("Long press detected! Resetting Wi-Fi...");
        wm.resetSettings();
        delay(500);
        ESP.restart();
      }
    }
  } else {
    buttonWasPressed = false;
  }

  if (millis() - ledOnMillis >= LED_DURATION) {
    digitalWrite(GREEN_LIGHT, LOW);
  }
  // Run kada 2 secs
  if (millis() - lastReadMillis >= 2000) {
    lastReadMillis = millis();
    if (!sensor.isSensorDataValid()) {
      return;
    }

    SensorData sensorData = sensor.getData();

    netc.publishEnergyData(sensorData);

    digitalWrite(GREEN_LIGHT, HIGH);
    ledOnMillis = millis();
  }
}