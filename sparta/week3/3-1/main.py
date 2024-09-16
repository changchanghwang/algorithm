import sys
from collections import deque

# 미로 입력 받기
maze = [list(map(int, list(line.strip()))) for line in sys.stdin if line.strip() != '']

N = len(maze)
M = len(maze[0])

# 이동 방향 (상하좌우)
dx = [0, 0, -1, 1]
dy = [-1, 1, 0, 0]

# 방문 여부와 거리 저장
visited = [[False] * M for _ in range(N)]
distance = [[0] * M for _ in range(N)]

# BFS 시작 위치 설정
queue = deque()
queue.append((0, 0))
visited[0][0] = True
distance[0][0] = 1

while queue:
    x, y = queue.popleft()
    
    # 출구에 도달하면 거리 출력 후 종료
    if x == N - 1 and y == M - 1:
        print(distance[x][y])
        sys.exit(0)
    
    # 상하좌우 위치 확인
    for i in range(4):
        nx = x + dx[i]
        ny = y + dy[i]
        
        # 미로 범위 내이고, 이동 가능한 경로이며, 방문하지 않은 위치인 경우
        if 0 <= nx < N and 0 <= ny < M:
            if maze[nx][ny] == 1 and not visited[nx][ny]:
                queue.append((nx, ny))
                visited[nx][ny] = True
                distance[nx][ny] = distance[x][y] + 1

# 출구에 도달할 수 없는 경우
print(-1)
