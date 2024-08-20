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
            result[result.length - 1][1] = Math.max(result[result.length - 1][1], interval[1]);
        }
    }
    
    return result;
};


const result = merge([[1,4],[0,2],[3,5]]);
console.log(result);
