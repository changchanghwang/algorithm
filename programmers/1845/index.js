function solution(nums) {
  const speciesCount = new Set(nums).size

  return speciesCount >= nums.length/2 ?  nums.length/2 :speciesCount
}

console.log(solution([3,1,2,3]))