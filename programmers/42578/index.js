function solution(clothes) {
  const map = new Map();
  clothes.forEach(([name, type]) => {
    map.set(type, [...(map.get(type) ?? []), name]);
  });

  let count = 1;
  if (map.size === 1) {
    return clothes.length;
  }

  for (const [, value] of map) {
    count *= value.length + 1;
  }

  return count - 1;
}

console.log(
  solution([
    ['yellow_hat', 'headgear'],
    ['blue_sunglasses', 'eyewear'],
    ['green_turban', 'headgear'],
  ]),
);
