function isStrongPassword(pwd) {
  return (
    typeof pwd === "string" &&
    pwd.length >= 8 &&
    /[A-Z]/.test(pwd) && // at least one uppercase letter
    /\d/.test(pwd) // at least one digit
  );
}
module.exports = isStrongPassword;
