
#include <PZEM004Tv30.h>

#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>
#include <Wire.h>

#include <PubSubClient.h>
#include <WiFi.h>

/* Hardware Serial2 is only available on certain boards.
 * For example the Arduino MEGA 2560
 */
#if defined(ESP32)
PZEM004Tv30 pzem(Serial2, 16, 17);
#else
PZEM004Tv30 pzem(Serial2);
#endif

const uint16_t SCREEN_WIDTH = 128;
const uint16_t SCREEN_HEIGHT = 64;
const int8_t OLED_RESET = -1;

Adafruit_SSD1306 display(SCREEN_WIDTH, SCREEN_HEIGHT, &Wire, OLED_RESET);

const char *ssid = "GlobeAtHome_A0177_2.4";
const char *password = "Octocat.2024..";
const char *mqtt_server = "192.168.254.101";

WiFiClient espClient;
PubSubClient client(espClient);

const char *device_id = "esp32s2-01";

static unsigned long lastRead = 0;

void reconnect() {
  if (client.connected())
    return;

  Serial.print("Attempting MQTT connection...");
  if (client.connect(device_id)) {
    Serial.println("connected");
  } else {
    Serial.print("failed, rc=");
    Serial.println(client.state());
  }
}

void setup() {
  Serial.begin(115200);

  // Serial2.begin(9600, SERIAL_8N1, 16, 17);

  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
  }

  client.setServer(mqtt_server, 1883);

  if (!display.begin(SSD1306_SWITCHCAPVCC, 0x3C)) {
    Serial.println("OLED not found");
    while (true)
      ;
  }

  display.setTextSize(2);
  display.setTextColor(SSD1306_WHITE);

  // Uncomment in order to reset the internal energy counter
  // pzem.resetEnergy();
}

void loop() {
  if (!client.connected())
    reconnect();
  client.loop();

  char topic[50];
  snprintf(topic, sizeof(topic), "sensors/%s/power", device_id);

  if (millis() - lastRead >= 2000) {

    Serial.print("Custom Address:");
    Serial.println(pzem.readAddress(), HEX);

    // Read the data from the sensor
    float voltage = pzem.voltage();
    float current = pzem.current();
    float power = pzem.power();
    float energy = pzem.energy();
    float frequency = pzem.frequency();
    float pf = pzem.pf();

    lastRead = millis();

    // Check if the data is valid
    if (isnan(voltage)) {
      Serial.println("Error reading voltage");
    } else if (isnan(current)) {
      Serial.println("Error reading current");
    } else if (isnan(power)) {
      Serial.println("Error reading power");
    } else if (isnan(energy)) {
      Serial.println("Error reading energy");
    } else if (isnan(frequency)) {
      Serial.println("Error reading frequency");
    } else if (isnan(pf)) {
      Serial.println("Error reading power factor");
    } else {

      // Print the values to the Serial console
      char payload[100];
      snprintf(payload, sizeof(payload), "{\"voltage\":%.2f,\"power\":%.3f}",
               voltage, power);

      client.publish(topic, payload);

      display.clearDisplay();
      display.setCursor(10, 20);

      display.printf("%.2fV\n %.3fW", voltage, power);
      display.display();
    }
  }
}
