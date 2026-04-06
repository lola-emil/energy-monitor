#pragma once

#include <Arduino.h>

bool buttonWasPressed = false;

constexpr uint16_t LONG_PRESS_TIME = 5000;
constexpr uint8_t BUTTON_PIN = 14;

void initButton() {
    pinMode(BUTTON_PIN, INPUT_PULLUP);
}

bool isLongPressed() {
  static bool buttonWasPressed = false;
  static unsigned long buttonPressStart = 0;

  bool buttonPressed = digitalRead(BUTTON_PIN) == LOW;

  if (buttonPressed) {
    if (!buttonWasPressed) {
      buttonPressStart = millis();
      buttonWasPressed = true;
    } else {
      if (millis() - buttonPressStart >= LONG_PRESS_TIME) {
        buttonWasPressed = false;
        return true;
      }
    }
  } else {
    buttonWasPressed = false;
  }

  return false;
}