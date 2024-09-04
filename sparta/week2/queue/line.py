from collections import deque
from structures import Queue

def test_problem_queue(num):
    queue = Queue()
    for i in range(1, num+1):
        queue.push(i)
        
    result = 0;
    while True:
        result = queue.pop()
        if queue.is_empty():
            break
        queue.push(queue.pop())

    return result


def test_problem_queue2(num):
    deq = deque([i for i in range(1, num+1)])
    while len(deq) > 1:
        deq.popleft()
        deq.append(deq.popleft())
        
    return deq[0]
        
