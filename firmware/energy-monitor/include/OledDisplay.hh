#pragma once

#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>

class OledDisplay {

private:
    Adafruit_SSD1306& display;

    boolean oledNotFound = false;

public:
    OledDisplay(Adafruit_SSD1306& d) : display(d) {}

    void setup() {
        if (!display.begin(SSD1306_SWITCHCAPVCC, 0x3C)) {
            Serial.println("OLED not found");
            oledNotFound = true;
        }

        display.setTextSize(2);
        display.setTextColor(SSD1306_WHITE);
    }

    void printf(const char* format, ...) {
        if (oledNotFound) return;

        char buffer[64];
        va_list args;
        va_start(args, format);

        vsnprintf(buffer, sizeof(buffer), format, args);

        va_end(args);

        display.clearDisplay();
        display.setCursor(10, 20);

        display.printf("%s", buffer);
        display.display();
    }
};