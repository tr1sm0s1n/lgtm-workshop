import test from 'node:test'
import assert from 'node:assert/strict'
import { sum, difference, product, quotient } from './math.js'

test('sum function should correctly return the sum', () => {
  assert.strictEqual(sum(2, 3), 5, '2 + 3 should be 5')
})

test('difference function should correctly return the difference', () => {
  assert.strictEqual(difference(7, 6), 1, '7 - 6 should be 1')
})

test('product function should correctly return the product', () => {
  assert.strictEqual(product(3, 3), 9, '3 * 3 should be 9')
})

test('quotient function should correctly return the quotient', () => {
  assert.strictEqual(quotient(8, 4), 2, '8 / 4 should be 2')
})
