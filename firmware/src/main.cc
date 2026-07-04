#include <Arduino.h>

#include "device_utils.hh"
#include "mqtt_manager.hh"
#include "pins.hh"
#include "pzem_manager.hh"
#include "wifi_manager.hh"

const char *mqtt_server = "192.168.254.119";

SemaphoreHandle_t mutex;

uint32_t lastPublish = 0;
uint32_t lastReconnect = 0;

String DEVICE_ID;

PZEMData pzemData;

void handlePZEM(void *) {
  while (true) {
    xSemaphoreTake(mutex, portMAX_DELAY);
    readPZEM(pzemData);
    xSemaphoreGive(mutex);

    vTaskDelay(pdMS_TO_TICKS(200));
  }
}

void handleMQTT(void *pvParameters) {
  PZEMData dataCopy;

  while (true) {
    if (!isMQTTConnected()) {
      if (millis() - lastReconnect > 3000) {
        lastReconnect = millis();
        reconnectMQTT();
      }

      vTaskDelay(pdMS_TO_TICKS(100));
      continue;
    }

    mqttLoop();

    if (millis() - lastPublish >= 1000) {
      lastPublish = millis();

      if (xSemaphoreTake(mutex, pdMS_TO_TICKS(100)) == pdTRUE) {
        dataCopy = pzemData;
        xSemaphoreGive(mutex);
      }

      if (isReadingValid(dataCopy)) {
        sendData(DEVICE_ID, dataCopy);

        digitalWrite(LED_MQTT_PIN, LED_ON);
        vTaskDelay(pdMS_TO_TICKS(50));
        digitalWrite(LED_MQTT_PIN, LED_OFF);
      }
    }

    vTaskDelay(pdMS_TO_TICKS(10));
  }
}

void setup() {
  Serial.begin(115200);
  delay(300);

  DEVICE_ID = getDeviceID();

  pinMode(LED_MQTT_PIN, OUTPUT);
  pinMode(LED_POWER_PIN, OUTPUT);

  digitalWrite(LED_POWER_PIN, LED_ON);
  digitalWrite(LED_MQTT_PIN, LED_OFF);

  initPZEM();
  setupWiFi();
  initMQTT(mqtt_server);

  mutex = xSemaphoreCreateMutex();

  if (mutex == NULL) {
    Serial.println("Failed to create mutex");
    while (true)
      ;
  }

  xTaskCreatePinnedToCore(handlePZEM, "PZEM", 2048, NULL, 1, NULL, 0);
  xTaskCreatePinnedToCore(handleMQTT, "MQTT", 8192, NULL, 1, NULL, 1);
}

void loop() {}