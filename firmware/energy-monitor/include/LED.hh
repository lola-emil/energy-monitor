#pragma once

#include <Arduino.h>

// LED CONFIG
constexpr unsigned long LED_DURATION = 100;

constexpr uint8_t GREEN_LIGHT = 32;
constexpr uint8_t YELLOW_LIGHT = 33;

void initLED() {
    pinMode(GREEN_LIGHT, OUTPUT);
    pinMode(YELLOW_LIGHT, OUTPUT);
}

void turnOnLED(uint8_t pin) {
    digitalWrite(pin, HIGH);
}


void turnOffLED(uint8_t pin) {
    digitalWrite(pin, LOW);
}
