def overlap(range1, range2):
    is_overlapping = True
    out_range = []
    if range1[0] > range2[1] or range2[0] > range1[1]:
        is_overlapping = False
    if is_overlapping:
        out_range = [min(range1[0], range2[0]), max(range1[1], range2[1])]
    return is_overlapping, out_range

fresh = []
ingredients = []

with open("5/5.txt") as fs:
    for line in fs.readlines():
        if "-" in line:
            fresh.append([int(n) for n in line[:-1].split("-")])
        elif line != "\n":
            ingredients.append(int(line[:-1]))

while True:
    last_n_fresh = len(fresh)
    for idx in range(last_n_fresh):
        for jdx in range(last_n_fresh):
            if idx != jdx and idx < len(fresh) and jdx < len(fresh):
                is_overlap, new_range = overlap(fresh[idx], fresh[jdx])
                if is_overlap:
                    fresh = [rng for i, rng in enumerate(fresh) if i not in [idx, jdx]]
                    fresh.append(new_range)
    if len(fresh) == last_n_fresh:
        break

print(sum([rng[1] - rng[0] + 1 for rng in fresh]))
print(sum([
    any([rng[0] <= ingredient <= rng[1] for rng in fresh])
    for ingredient in ingredients
]))
