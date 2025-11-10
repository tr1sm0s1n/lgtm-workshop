// account.js

function createAccount(initialBalance = 0) {
  if (initialBalance < 0) {
    throw new Error("Initial balance cannot be negative");
  }

  let balance = initialBalance;

  function deposit(amount) {
    if (amount <= 0) {
      console.log("Deposit amount must be positive");
      return;
    }
    balance += amount;
    console.log(`Deposited ₹${amount}. Current balance: ₹${balance}`);
  }

  function withdraw(amount) {
    if (amount <= 0) {
      console.log("Withdrawal amount must be positive");
      return;
    }

    if (amount > balance) {
      console.log("Insufficient funds. Withdrawal denied.");
      return;
    }

    balance -= amount;
    console.log(`Withdrew ₹${amount}. Current balance: ₹${balance}`);
  }

  function getBalance() {
    return balance;
  }

  return {
    deposit,
    withdraw,
    getBalance,
  };
}

// Example usage:
const myAccount = createAccount(1000);
myAccount.deposit(500);
myAccount.withdraw(200);
myAccount.withdraw(2000); // Prevents overdraft
console.log("Final Balance:", myAccount.getBalance());
