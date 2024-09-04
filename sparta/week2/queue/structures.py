class Node:
    def __init__(self, value =0,next=None):
        self.value = value
        self.next = next

class Queue:
    def __init__(self):
        self.front = None

    def push(self, value):
        if not self.front:
            self.front = Node(value)
            return
        
        current = self.front
        while current.next is not None:
            current = current.next

        current.next = Node(value)
    
    def pop(self):
        if not self.front:
            return None

        front = self.front
        self.front = front.next
        return front.value
    
    def is_empty(self):
        return self.front is None