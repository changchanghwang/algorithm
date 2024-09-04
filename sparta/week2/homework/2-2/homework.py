def solution(str:str) -> int:
    dic = {}

    start = [0]
    count = 0
    for i in range(0,len(str)):
        c = str[i]
        if dic.get(c) is None:
            dic[c] = i
            start[count] += 1
        else:
            count += 1
            start.append(1)
    
    return max(start)

