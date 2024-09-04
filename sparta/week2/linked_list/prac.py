def is_palindrome(ln):
    arr = []
    head = ln.head

    if not head:
        return True
    
    node = head
    while node:
        arr.append(node.val)
        node = node.next
    
    while len(arr) > 1:
        first = arr.pop(0)
        last = arr.pop()
        if first != last:
            return False
        
    return True