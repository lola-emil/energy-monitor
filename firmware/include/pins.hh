#pragma once

#include <Arduino.h>

#ifndef NODE_ID
#define NODE_ID 0
#endif

constexpr gpio_num_t PZEM_RX_PIN = GPIO_NUM_16;
constexpr gpio_num_t PZEM_TX_PIN = GPIO_NUM_17;

constexpr gpio_num_t LED_POWER_PIN = GPIO_NUM_5; // Stable, safe pin
constexpr gpio_num_t LED_MQTT_PIN = GPIO_NUM_4;

constexpr gpio_num_t BUTTON_RESET_PIN = GPIO_NUM_18;
constexpr uint8_t BUTTON_ACTIVE_STATE = LOW;

constexpr bool LED_ON = HIGH;
constexpr bool LED_OFF = LOW;

