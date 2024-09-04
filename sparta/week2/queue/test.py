from structures import Queue

queue = Queue()

queue.push(1)
queue.push(2)
queue.push(3)

assert queue.pop() == 1
assert queue.pop() == 2
assert queue.pop() == 3
assert queue.is_empty()