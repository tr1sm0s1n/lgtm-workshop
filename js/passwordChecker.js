function isStrongPassword(pwd) {
  const hasUppercase = /[A-Z]/.test(pwd);
  const hasDigit = /\d/.test(pwd);
  const isLongEnough = pwd.length >= 8;

  return hasUppercase && hasDigit && isLongEnough;
}
