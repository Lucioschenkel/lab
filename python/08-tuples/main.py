# A Tuple is like an immutable list
my_tuple = (1, 2, 3)

# my_tuple[1] = 10 # This would fail

print(my_tuple)  # (1, 2, 3)

# Items can be unpacked
x, y, z = my_tuple

print(x)  # 1
print(y)  # 2
print(z)  # 3

# Tuples can be sliced
print(my_tuple[1:])  # (2, 3)

# To create a Tuple with a single item, use a trailing comma
another_tuple = (1,)

print(another_tuple)  # (1,)

# Count occurrences of an item
print(my_tuple.count(1))  # 1

# Get the (first) index of an element
print(my_tuple.index(3))  # 2
