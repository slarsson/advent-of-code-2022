import functools

rows = open('./input.txt', 'r').read().split('\n')

sums = [0]
cursor = 0
for row in rows:
    if row == '':
        sums.append(0)
        cursor += 1
        continue
    sums[cursor] += int(row)

sums.sort(reverse=True)

print(f'a: {sums[0]}')
print(f'b: {functools.reduce(lambda a, b: a + b, sums[:3])}')
