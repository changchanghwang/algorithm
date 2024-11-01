function narcissistic(value) {
  const stringValue = String(value);
  const digits = stringValue.length;

  let result = 0;
  for (let i = 0; i < stringValue.length; i++) {
    result += Number(stringValue[i]) ** digits;
  }

  return result === value;
}
