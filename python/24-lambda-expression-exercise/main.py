my_list = [5, 4, 3]

# square
print(list(map(lambda n: n * n, my_list)))

# list sorting: sort based on the second item in each tuple
a = [(0, 2), (4, 3), (9, 9), (10, -1)]
a.sort(key=lambda x: x[1])
print(a)
