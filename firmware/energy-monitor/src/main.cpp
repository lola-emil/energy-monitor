
#include <PZEM004Tv30.h>

#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>
#include <Wire.h>

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

void setup() {
  Serial.begin(115200);

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

  Serial.print("Custom Address:");
  Serial.println(pzem.readAddress(), HEX);

  // Read the data from the sensor
  float voltage = pzem.voltage();
  float current = pzem.current();
  float power = pzem.power();
  float energy = pzem.energy();
  float frequency = pzem.frequency();
  float pf = pzem.pf();

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
    Serial.print("Voltage: ");
    Serial.print(voltage);
    Serial.println("V");
    Serial.print("Current: ");
    Serial.print(current);
    Serial.println("A");
    Serial.print("Power: ");
    Serial.print(power);
    Serial.println("W");
    Serial.print("Energy: ");
    Serial.print(energy, 3);
    Serial.println("kWh");
    Serial.print("Frequency: ");
    Serial.print(frequency, 1);
    Serial.println("Hz");
    Serial.print("PF: ");
    Serial.println(pf);

    display.clearDisplay();
    display.setCursor(10, 20);

    display.printf("%.2fV\n %.3fkWh", voltage, energy);
    display.display();
  }

  Serial.println();
  delay(2000);
}