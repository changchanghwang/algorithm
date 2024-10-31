function solution(array, commends) {
  return commends.map(([start, end, index]) => array.slice(start - 1, end).sort((a, b) => a - b)[index - 1]);
}

console.log(
  solution(
    [1, 5, 2, 6, 3, 7, 4],
    [
      [2, 5, 3],
      [4, 4, 1],
      [1, 7, 3],
    ],
  ),
);
