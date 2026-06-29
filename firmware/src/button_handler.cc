#include <Arduino.h>
#include <WiFiManager.h>

#include "button_handler.hh"
#include "pins.hh"

static unsigned long buttonPressStart = 0;
static bool buttonPressed = false;

const unsigned long LONG_PRESS_TIME = 5000;

void initButton() {
    pinMode(BUTTON_RESET_PIN, INPUT_PULLUP);
}

void handleButton() {
    bool isPressed = digitalRead(BUTTON_RESET_PIN) == BUTTON_ACTIVE_STATE;

    if (isPressed && !buttonPressed) {
        buttonPressed = true;
        buttonPressStart = millis();
    }

    if (!isPressed && buttonPressed) {
        buttonPressed = false;
    }

    if (buttonPressed && (millis() - buttonPressStart >= LONG_PRESS_TIME)) {
        Serial.println("Long press detected! Resetting WiFi...");

        digitalWrite(LED_MQTT_PIN, LED_ON);

        WiFiManager wm;
        wm.resetSettings();

        delay(1000);
        ESP.restart();
    }
}