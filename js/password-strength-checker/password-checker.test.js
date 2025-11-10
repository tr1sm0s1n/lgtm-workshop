import test from "node:test";
import assert from "node:assert/strict";
import isStrongPassword from "./password-checker.js";

test("should return true for valid strong passwords", () => {
  assert.strictEqual(
    isStrongPassword("Password1"),
    true,
    '"Password1" is valid'
  );
  assert.strictEqual(
    isStrongPassword("StrongPass9"),
    true,
    '"StrongPass9" is valid'
  );
});

test("should return false if no uppercase letter", () => {
  assert.strictEqual(
    isStrongPassword("password1"),
    false,
    "no uppercase letter"
  );
});

test("should return false if no digit", () => {
  assert.strictEqual(isStrongPassword("Password"), false, "no number present");
});

test("should return false if length < 8", () => {
  assert.strictEqual(isStrongPassword("Pass1"), false, "length is less than 8");
});

test("should return false if input is not a string", () => {
  assert.strictEqual(isStrongPassword(null), false, "null is not valid");
  assert.strictEqual(isStrongPassword(12345678), false, "number is not valid");
});
