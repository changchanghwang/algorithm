# 56. Merge Intervals

> Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

```js
/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */
function merge(intervals) {
  intervals.sort((a, b) => a[0] - b[0]); // 항상 왼쪽이 start가 작은 순서로 정렬
  const result = [];

  for (let interval of intervals) {
    if (result.length === 0 || interval[0] > result[result.length - 1][1]) {
      result.push(interval);
    } else {
      // 겹치는 경우, 이전 구간의 끝 시간을 업데이트
      result[result.length - 1][1] = Math.max(
        result[result.length - 1][1],
        interval[1]
      );
    }
  }

  return result;
}
```

```js
intervals.sort((a, b) => a[0] - b[0]);
```

1. 항상 왼쪽에 start가 작은 배열이 오도록 정렬함

> [[1,5],[0,6]] -> [[0,6],[1,5]]

```js
if (result.length === 0 || interval[0] > result[result.length - 1][1]) {
  result.push(interval);
}
```

2. interval 배열을 돌면서 만약 첫번째 interval이면 무조건 result에 푸시, 그게 아니라면 interval의 start와 result의 마지막 원소의 end를 비교해서 interval의 start가 더 크다면 겹치지 않는다고 판단, result에 push

> result : [[2,4]]
> interval : [5,6]
> --> result = [[2,4],[5,6]]
> 이 경우 interval의 start는 5, result 마지막 원소의 end가 4인 경우이기 때문에 합치지 않고 그대로 push한다.

```js
else {
    result[result.length - 1][1] = Math.max(result[result.length - 1][1], interval[1]);
}
```

3. 2번에 해당하지 않는 경우는 전부 겹친다고 판단하여 result의 마지막 원소의 end를 큰쪽으로 변경해준다.

> 이 경우 2가지 가능성이 있다.
>
> 1. result의 마지막 원소:[3,7], interval: [4,9] -> result의 마지막 원소: [3,9]
> 2. result의 마지막 원소: [3,7], interval: [4,5] -> result의 마지막 원소: [3,7]
