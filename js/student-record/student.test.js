// student.test.js
import test from 'node:test';
import assert from 'node:assert/strict';
import { Student } from './student.js';

test('total() should correctly return the total marks', () => {
  const s1 = new Student('Alice', [80, 90, 100]);
  assert.strictEqual(s1.total(), 270, 'Total should be 270');
});

test('passed() should return true if all marks ≥ 40', () => {
  const s2 = new Student('Bob', [50, 60, 70]);
  assert.strictEqual(s2.passed(), true, 'Should pass when all marks >= 40');
});

test('passed() should return false if any mark < 40', () => {
  const s3 = new Student('Charlie', [45, 30, 60]);
  assert.strictEqual(s3.passed(), false, 'Should fail when any mark < 40');
});
