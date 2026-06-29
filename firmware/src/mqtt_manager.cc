#include "mqtt_manager.hh"
#include "pzem_manager.hh"
#include <WiFi.h>

WiFiClient espClient;
PubSubClient client(espClient);

static const char *mqtt_server;

void initMQTT(const char* server) {
    mqtt_server = server;
    client.setServer(mqtt_server, 1883);
}

void mqttLoop() {
    client.loop();
}

bool reconnectMQTT() {
    String clientId = "ESP32_Node_" + String(NODE_ID);

    if (client.connect(clientId.c_str())) {
        Serial.println("MQTT connected");
        return true;
    } else {
        Serial.print("MQTT failed, rc=");
        Serial.println(client.state());
        return false;
    }
}

void sendData(String deviceID, PZEMData& data) {
    String topic = "energy/readings/" + deviceID;

    String payload = "{";
    payload += "\"device_code\":\"" + deviceID + "\",";
    payload += "\"voltage\":" + String(data.voltage, 2) + ",";
    payload += "\"current\":" + String(data.current, 2) + ",";
    payload += "\"power\":" + String(data.power, 2) + ",";
    payload += "\"energy\":" + String(data.energy, 2);
    payload += "}";

    Serial.println(payload);
    client.publish(topic.c_str(), payload.c_str());
}

bool isMQTTConnected() {
    return client.connected();
}