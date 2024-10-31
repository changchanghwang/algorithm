function solution(k, tangerine) {
  let result = 0
  const count = {}

  tangerine.forEach((size)=>{
    count[size] = (count[size] || 0) + 1
  })

  const sortedCount = Object.values(count).sort((a,b)=>b-a)

  for (let i = 0; i < sortedCount.length; i++) {
      result++
    if(k > sortedCount[i]) {
        k -= sortedCount[i]
    }else if(k > 0){
        break;
    }
  }

  return result
}

console.log(solution(6,[1, 3, 2, 5, 4, 5, 2, 3]	))