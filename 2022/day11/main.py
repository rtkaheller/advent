def inspect(monkeys, m, relief, gcd):
  for old in m["items"]:
    m["inspects"] += 1
    new = int(eval(m["op"].split('=')[1]))
    new //= relief
    if new % m["test"] == 0:
      monkeys[m["true"]]["items"].append(new % gcd)
    else:
      monkeys[m["false"]]["items"].append(new % gcd)
  m["items"] = []


def main(file, rounds, relief):
  monkeys = []
  with open(file, "r") as f:
    lines = f.readlines()
    for i, line in enumerate(lines):
      if len(line) == 0:
        continue
      if line.startswith("Monkey"):
        monkey = {"inspects": 0}
        monkey["items"] = [int(x.strip()) for x in lines[i+1][17:].split(',')]
        monkey["op"] = lines[i+2][12:].strip()
        monkey["test"] = int(lines[i+3][21:].strip())
        monkey["true"] = int(lines[i+4][28:].strip())
        monkey["false"] = int(lines[i+5][29:].strip())
        monkeys.append(monkey)

  gcd = 1
  for m in monkeys:
    gcd *= m["test"]

  for i in range(rounds):
    counts = []
    for m in monkeys:
      inspect(monkeys, m, relief, gcd)
      counts.append(m["inspects"])
  counts.sort()

  return counts[-1] * counts[-2]


if __name__ == "__main__":
  file = "input.txt"
  print(main(file, 20, 3))
  print(main(file, 10000, 1))
