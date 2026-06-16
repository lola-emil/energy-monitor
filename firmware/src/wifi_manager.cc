#include "wifi_manager.hh"
#include <Arduino.h>
#include "pins.hh"

void configModeCallback(WiFiManager *wm) {
    Serial.println("Entered AP mode");
    digitalWrite(LED_MQTT_PIN, LED_ON);
}

void setupWiFi() {
    WiFiManager wm;
    wm.setAPCallback(configModeCallback);

    wm.autoConnect("ESP32-Setup");

    Serial.println("WiFi connected");
    digitalWrite(LED_MQTT_PIN, LED_OFF);
}