import test from 'node:test'
import assert from 'node:assert/strict'
import { createAccount } from './bank.js'

// Test: creating account with default balance
test('Account should initialize with default balance 0', () => {
  const acc = createAccount()
  assert.strictEqual(acc.getBalance(), 0)
})

// Test: deposit functionality
test('deposit() should correctly add amount to balance', () => {
  const acc = createAccount()
  acc.deposit(500)
  assert.strictEqual(acc.getBalance(), 500)
})

// Test: deposit invalid amount
// test('deposit() should throw error for non-positive amount', () => {
//   const acc = createAccount()
//   assert.throws(() => acc.deposit(-100), /Deposit amount must be positive/)
// })

// Test: withdraw valid amount
test('withdraw() should correctly subtract amount from balance', () => {
  const acc = createAccount(1000)
  acc.withdraw(400)
  assert.strictEqual(acc.getBalance(), 600)
})

// Test: withdraw more than balance
// test('withdraw() should throw error when balance is insufficient', () => {
//   const acc = createAccount(200)
//   assert.throws(() => acc.withdraw(500), /Insufficient funds/)
// })

// Test: withdraw invalid amount
// test('withdraw() should throw error for non-positive amount', () => {
//   const acc = createAccount(500)
//   assert.throws(() => acc.withdraw(0), /Withdrawal amount must be positive/)
// })

// Test: invalid initial balance
test('createAccount() should throw error for negative initial balance', () => {
  assert.throws(() => createAccount(-100), /Initial balance cannot be negative/)
})
