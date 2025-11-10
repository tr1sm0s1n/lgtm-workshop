function convertTemperature(value, unit) {
    if (unit === "C") {
        // Convert Celsius to Fahrenheit
        return (value * 9 / 5) + 32;
    } else if (unit === "F") {
        // Convert Fahrenheit to Celsius
        return (value - 32) * 5 / 9;
    } else {
        // Invalid unit
        return null;
    }
}