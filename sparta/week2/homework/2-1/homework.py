def solution(input):
    stack = []
    for c in input.split():
        if c.isdigit():
            stack.append(c)
        else:
            b = stack.pop()
            a = stack.pop()
            if c == '+':
                stack.append(int(a) + int(b))
            elif c == '-':
                stack.append(int(a) - int(b))
            elif c == '*':
                stack.append(int(a) * int(b))
            elif c == '/':
                stack.append(int(a) / int(b))
    return stack.pop()