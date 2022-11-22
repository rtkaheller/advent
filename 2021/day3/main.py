#in_file = "fake.txt"
in_file = "input.txt"

with open(in_file, "r") as f:
    lines = list(map(lambda x: x.strip(), f.readlines()))

for i in range(len(lines[0])):
    count1 = 0
    count0 = 0
    for line in lines:
        if line[i] == "1":
            count1 += 1
        if line[i] == "0":
            count0 += 1
    print(count1, count0, len(lines), i)
    if count1 >= count0:
        keep = "1"
    else:
        keep = "0"

    print(keep)
    lines = list(filter(lambda x: x[i] == keep, lines))
    #print(lines)
    if len(lines) == 1:
        break

ox = int(lines[0].strip(), 2)

with open(in_file, "r") as f:
    lines = list(map(lambda x: x.strip(), f.readlines()))

for i in range(len(lines[0])):
    count1 = 0
    count0 = 0

    for line in lines:
        if line[i] == "1":
            count1 += 1
        if line[i] == "0":
            count0 += 1
    print(count1, count0, len(lines), i)
    if count1 < count0:
        keep = "1"
    else:
        keep = "0"

    print(keep)
    lines = list(filter(lambda x: x[i] == keep, lines))
    #print(lines)
    if len(lines) == 1:
        break

co = int(lines[0].strip(), 2)

print(ox, co, ox*co)
