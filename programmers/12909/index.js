function solution(s) {
  const stack = [];

  for (let i = 0; i < s.length; i++) {
    if (['(', '{', '['].includes(s[i])) {
      stack.push(s[i]);
    }
    if (s[i] === ']' && stack.pop() !== '[') return false;
    if (s[i] === '}' && stack.pop() !== '{') return false;
    if (s[i] === ')' && stack.pop() !== '(') return false;
  }

  return stack.length === 0;
}

console.log(solution(')()('));
