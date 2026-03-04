
#include <PZEM004Tv30.h>

#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>
#include <Wire.h>

#include <PubSubClient.h>
#include <WiFi.h>

#include "EnergySensor.hh"
#include "NetworkComm.hh"
#include "OledDisplay.hh"

#if defined(ESP32)
PZEM004Tv30 pzem(Serial2, 16, 17);
#else
PZEM004Tv30 pzem(Serial2);
#endif

EnergySensor sensor(pzem);

const uint16_t SCREEN_WIDTH = 128;
const uint16_t SCREEN_HEIGHT = 64;
const int8_t OLED_RESET = -1;

Adafruit_SSD1306 oled(SCREEN_WIDTH, SCREEN_HEIGHT, &Wire, OLED_RESET);
OledDisplay display(oled);

WiFiClient wifi;
PubSubClient client(wifi);
NetworkComm netc(client);

static unsigned long lastReadMillis = 0;

void setup() {
  Serial.begin(115200);

  netc.initConnection();
  netc.setChipID(ESP.getEfuseMac());

  display.setup();
}

void loop() {
  netc.connectMQTT();

  // Run kada 2 secs
  if (millis() - lastReadMillis >= 2000) {

    lastReadMillis = millis();
    if (!sensor.isSensorDataValid()) {
      display.printf("SENS_ERR");

      return;
    }

    SensorData sensorData = sensor.getData();

    netc.publishEnergyData(sensorData);

    display.printf("%.2fV\n %.3fW", sensorData.voltage, sensorData.power);

  }
}