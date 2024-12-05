const input = [1, 1, 1, 2, 2, 3];

// Time complexity: O(n)
// Space complexity: O(n)
function topKFrequent(nums, k) {
  const count = new Map();
  nums.forEach((num) => {
    count.set(num, (count.get(num) || 0) + 1);
  }); // { 1 => 3, 2 => 2, 3 => 1 }

  const frequent = Array.from({ length: nums.length + 1 }, () => []); // [[], [], [], [], ...] 처럼 0부터 nums.length까지 빈 배열을 생성

  for (const [num, freq] of count) {
    frequent[freq].push(num);
  } // [[], [3], [2], [1], ...] 처럼 빈 배열에 빈도수에 해당하는 숫자를 넣음

  const result = [];

  // index가 빈도수, value가 숫자이기 때문에 뒤부터 돌아야 빈도수가 높은 숫자부터 result를 넣을 수 있다.
  for (let i = frequent.length - 1; i >= 0; i--) {
    // 빈도수가 높은 숫자부터 result에 넣음
    for (const num of frequent[i]) {
      result.push(num);
      // k개만큼 result에 넣으면 종료
      if (result.length === k) {
        return result;
      }
    }
  }

  return result;
}

console.log(topKFrequent(input, 2));
