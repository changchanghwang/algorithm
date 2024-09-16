def generate_permutations(arr):
    def backtrack(start):
        if start == len(arr):
            permutations.append(arr[:])
        else:
            for i in range(start, len(arr)):
                arr[start], arr[i] = arr[i], arr[start]
                backtrack(start + 1)
                arr[start], arr[i] = arr[i], arr[start]  # 원래 상태로 복구

    permutations = []
    backtrack(0)
    return permutations

def print_permutations(arr):
    perms = generate_permutations(arr)
    for perm in perms:
        print(perm)

# 테스트
print("Example 1:")
print_permutations([1, 2, 3])

print("\nExample 2:")
print_permutations([1, 2, 3, 4])