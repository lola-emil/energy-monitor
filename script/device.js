const mqtt = require("mqtt");
const devices = require("./devices.json");

const MQTT_BROKER = "mqtt://localhost:1883";
const REGISTRATION_CODE = "REG456";

const client = mqtt.connect(MQTT_BROKER);

// store per-device state
const deviceState = {};

client.on("connect", () => {
    console.log("✅ Connected");

    devices.forEach((DEVICE_CODE) => {
        const responseTopic = `device/${DEVICE_CODE}/register/response`;

        // initialize state
        deviceState[DEVICE_CODE] = {
            deviceId: null,
            token: null,
        };

        client.subscribe(responseTopic, (err) => {
            if (err) {
                console.error(`❌ Subscribe failed for ${DEVICE_CODE}:`, err);
                return;
            }

            console.log(`📥 Waiting response for ${DEVICE_CODE}...`);

            const payload = {
                s: DEVICE_CODE,
                c: REGISTRATION_CODE,
            };

            client.publish("device/register", JSON.stringify(payload));
        });
    });
});

client.on("message", (topic, message) => {
    try {
        const data = JSON.parse(message.toString());

        console.log(data);

        // match which device this response belongs to
        const match = topic.match(/^device\/(.+)\/register\/response$/);

        if (match) {
            const DEVICE_CODE = match[1];

            if (!deviceState[DEVICE_CODE]) return;

            deviceState[DEVICE_CODE].deviceId = data.device_id;
            deviceState[DEVICE_CODE].token = data.token;

            console.log(`✅ ${DEVICE_CODE} registered with ID:`, data.device_id);

            startSensorLoop(DEVICE_CODE);
        }
    } catch (err) {
        console.error("❌ Invalid JSON:", err);
    }
});

function startSensorLoop(DEVICE_CODE) {
    const state = deviceState[DEVICE_CODE];

    setInterval(() => {
        if (!state.deviceId) return;

        const payload = {
            token: state.token,
            voltage: randomFloat(210, 240),
            current: randomFloat(0.1, 10),
            power: randomFloat(0.01, 5),
        };

        const topic = `device/${state.deviceId}/sensor`;

        client.publish(topic, JSON.stringify(payload));

        console.log(`📊 [${DEVICE_CODE}] Sent:`, topic, payload);
    }, 2000);
}

function randomFloat(min, max) {
    return parseFloat((Math.random() * (max - min) + min).toFixed(3));
}