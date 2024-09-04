from structures import Stack

def test_problem_stack(str):
    stack = Stack()

    pair = {
        ")": "(",
        "}": "{",
        "]": "["
    }

    for c in str:
        if c in "({[":
            stack.push(c)
        else:
            top = stack.pop()
            if top != pair[c]:
               return False

    return stack.is_empty()