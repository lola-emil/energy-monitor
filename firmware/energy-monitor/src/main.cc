#include <PZEM004Tv30.h>

#include <PubSubClient.h>
#include <WiFi.h>
#include <WiFiManager.h>

#include "EnergySensor.hh"
#include "NetworkComm.hh"
#include "LED.hh"
#include "Button.hh"

#if defined(ESP32)
PZEM004Tv30 pzem(Serial2, RX2, TX2);
#else
PZEM004Tv30 pzem(Serial2);
#endif

EnergySensor sensor(pzem);

WiFiClient wifi;
WiFiManager wm;
PubSubClient client(wifi);
NetworkComm netc(client, wm);

static unsigned long lastReadMillis = 0;
static unsigned long ledOnMillis = 0;
static unsigned long buttonPressStart = 0;

void setup() {
  Serial.begin(115200);

  initLED();
  initButton();

  Serial2.begin(9600, SERIAL_8N1, RX2, TX2);

  netc.initConnection();
}

void loop() {
  if (!netc.connectMQTT()) {
    Serial.println("Wala pa ma connect");
    return;
  };

  if (isLongPressed()) {
    Serial.println("Long press detected! Resetting Wi-Fi...");
    wm.resetSettings();
    delay(500);
    ESP.restart();
  }

  if (millis() - ledOnMillis >= LED_DURATION) {
    turnOffLED(GREEN_LIGHT);
  }

  // Run kada 2 secs
  if (millis() - lastReadMillis >= 2000) {
    lastReadMillis = millis();
    if (!sensor.isSensorDataValid()) {
      return;
    }

    SensorData sensorData = sensor.getData();

    netc.publishEnergyData(sensorData);

    turnOnLED(GREEN_LIGHT);
    ledOnMillis = millis();
  }
}