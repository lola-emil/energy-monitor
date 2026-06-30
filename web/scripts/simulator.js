const mqtt = require("mqtt");

const BROKER = "mqtt://localhost:1883";
const DEVICE_ID = "EMS-PZEM001";

const client = mqtt.connect(BROKER);

client.on("connect", () => {
    console.log("Connected");
    setInterval(sendData, 1000);
});

let voltage = 230.0;
let current = 3.25;
let frequency = 60.00;
let energy = 0;

function drift(value, min, max, step) {
    value += (Math.random() * 2 - 1) * step;

    if (value > max) value = max;
    if (value < min) value = min;

    return value;
}

function sendData() {

    voltage = drift(voltage, 228.5, 231.5, 0.08);

    current = drift(current, 3.10, 3.40, 0.03);

    frequency = drift(frequency, 59.98, 60.02, 0.003);

    const power = voltage * current;

    energy += power / 1000 / 3600;

    const payload = {
        id: DEVICE_ID,
        v: Number(voltage.toFixed(2)),
        A: Number(current.toFixed(2)),
        W: Number(power.toFixed(2)),
        e_kWh: Number(energy.toFixed(3)),
        hz: Number(frequency.toFixed(2))
    };

    console.log(payload);

    client.publish(
        `energy/readings/${DEVICE_ID}`,
        JSON.stringify(payload)
    );
}