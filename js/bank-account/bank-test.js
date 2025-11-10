import test from 'node:test'
import assert from 'node:assert/strict'
import { createAccount, deposit, withdraw, getBalance } from './bank.js'

it('successfully created bank account')