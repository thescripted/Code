with open('input.txt') as f:
    data = [int(number) for number in f.read().splitlines()]

complement = {}
for number in data:
    if number in complement:
        print(number*complement[number])
        break
    complement[2020-number] = number
else:
    print(-1)

# Part 2:
