const win = {
    2: 5,
    3: 4,
    4: 3,
    5: 2,
    6: 1
}

function solution(lottos, win_nums) {
    const answer = [];
    const zeroCount = lottos.filter((num) => num === 0).length;
    const lottoSet = new Set(lottos);
    let winCount =0;
    win_nums.forEach((num)=>{
        if(lottoSet.has(num)) {
            winCount++;
        } 
    })
    answer.push(win[winCount+zeroCount] || 6)
    answer.push(win[winCount] || 6)
    return answer;
}

function solution2(lottos, win_nums) {
    const rank = [6,6,5,4,3,2,1]
    const min = win_nums.filter((num) =>lottos.includes(num)).length;
    const max = lottos.filter((num) => num === 0).length + min;

    return [rank[max],rank[min]]
}

console.log(solution2([0, 0, 0, 0, 0, 0], [38, 19, 20, 40, 15, 25]))