const assert = require("assert");
const convertTemperature = require("./script");

// Test 1: Celsius to Fahrenheit
assert.strictEqual(convertTemperature(0, "C"), 32, "0°C should be 32°F");
assert.strictEqual(convertTemperature(100, "C"), 212, "100°C should be 212°F");

// Test 2: Fahrenheit to Celsius
assert.strictEqual(convertTemperature(32, "F"), 0, "32°F should be 0°C");
assert.strictEqual(convertTemperature(212, "F"), 100, "212°F should be 100°C");

// Test 3: Invalid unit
assert.strictEqual(convertTemperature(50, "X"), null, "Invalid unit should return null");

console.log("✅ All tests passed!");
