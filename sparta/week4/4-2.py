def optimize_meetings(meetings):
    sorted_meetings = sorted(meetings, key=lambda x: x[1])
    
    selected_meetings = []
    last_end_time = 0
    
    for meeting in sorted_meetings:
        start_time, end_time = meeting
        if start_time >= last_end_time:
            selected_meetings.append(meeting)
            last_end_time = end_time
    
    return len(selected_meetings)

test_cases = [
    [(0, 6), (1, 4), (3, 5), (3, 8), (5, 7), (8, 9)],
    [(1, 3), (2, 4), (5, 8), (6, 10), (8, 11), (10, 12)]
]

for i, meetings in enumerate(test_cases, 1):
    result = optimize_meetings(meetings)
    print(f"테스트 케이스 {i}:")
    print(f"입력: {meetings}")
    print(f"출력: {result}")
    print()