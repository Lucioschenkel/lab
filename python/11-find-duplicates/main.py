some_list = ['a', 'b', 'c', 'b', 'd', 'm', 'n', 'n']

duplicates = []

for x in some_list:
    if some_list.count(x) > 1:
        if x not in duplicates:
            duplicates.append(x)

print(duplicates)
