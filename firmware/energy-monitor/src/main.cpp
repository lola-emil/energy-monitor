
#include <PZEM004Tv30.h>

#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>
#include <Wire.h>

#include <PubSubClient.h>
#include <WiFi.h>

#include "EnergySensor.hh"
#include "NetworkComm.hh"

#if defined(ESP32)
PZEM004Tv30 pzem(Serial2, 16, 17);
#else
PZEM004Tv30 pzem(Serial2);
#endif

EnergySensor sensor(pzem);

const uint16_t SCREEN_WIDTH = 128;
const uint16_t SCREEN_HEIGHT = 64;
const int8_t OLED_RESET = -1;

Adafruit_SSD1306 display(SCREEN_WIDTH, SCREEN_HEIGHT, &Wire, OLED_RESET);

WiFiClient wifi;
PubSubClient client(wifi);
NetworkComm netc(client);

static unsigned long lastReadMillis = 0;

void setup() {
  Serial.begin(115200);

  netc.initConnection();
  netc.setChipID(ESP.getEfuseMac());

  if (!display.begin(SSD1306_SWITCHCAPVCC, 0x3C)) {
    Serial.println("OLED not found");
    while (true)
      ;
  }

  display.setTextSize(2);
  display.setTextColor(SSD1306_WHITE);
}

void loop() {
  netc.connectMQTT();

  // Run kada 2 secs
  if (millis() - lastReadMillis >= 2000) {

    lastReadMillis = millis();

    /**
     * TODO: need validation ang reading result
     */

    SensorData sensorData = sensor.getData();
    netc.publishEnergyData(sensorData);

    display.clearDisplay();
    display.setCursor(10, 20);

    display.printf("%.2fV\n %.3fW", sensorData.voltage, sensorData.power);
    display.display();
  }
}