class Node:
    def __init__(self, val=0,next=None):
        self.val = val
        self.next = next

class Stack:
    def __init__(self):
        self.top = None

    def push(self,val):
        self.top = Node(val, self.top)
            

    def pop(self):
        if self.top is None:
            return None
        
        top = self.top
        self.top = top.next
        return top.val

    def is_empty(self):
        return self.top is None