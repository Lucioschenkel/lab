numbers = [1, 2, 3, 4]

# Add one item to the end of a list
numbers.append(5)

print(numbers)  # [1, 2, 3, 4, 5]

# Add multiple items to the end of the list
numbers.extend([6, 7, 8])

print(numbers)  # [1, 2, 3, 4, 5, 6, 7, 8]

# The assignment operator won't create a new list, since lists are stored by reference
another_list = numbers

# Remove one item from the end `another_list`
another_list.pop()
print(numbers)  # [1, 2, 3, 4, 5, 6, 7]

# If we want a copy, can can use the `.copy` method
yet_another = another_list.copy()

yet_another.append(8)

print(yet_another)  # [1, 2, 3, 4, 5, 6, 7, 8]
print(another_list)  # [1, 2, 3, 4, 5, 6, 7]

# Empty the list
yet_another.clear()

print(yet_another)  # []

# Return the index of a number in the list. Throws if not in the list
print(numbers.index(2))  # 1

# This will return True/False if item is/isn't in the list
print(8 in numbers)  # False

# Count occurrences of elements in a list
print(numbers.count(1))  # 1

# Sorting a list (in place)
abcs = ["z", "a", "r", "c", "p"]

abcs.sort()

print(abcs)  # ['a', 'c', 'p', 'r', 'z']

# Sorting a copy of a list
initial = [1, 8, 3, 9, 4, 2, 5, 7, 6]

copied = sorted(initial)

print(initial)  # [1, 8, 3, 9, 4, 2, 5, 7, 6]
print(copied)  # [1, 2, 3, 4, 5, 6, 7, 8, 9]

# Reverse a list (in place)
copied.reverse()

print(copied)  # [9, 8, 7, 6, 5, 4, 3, 2, 1]

# List unpacking: this is similar to "destructuring" in JavaScript
# the main difference is the "rest" of the list, instead of ..., we use *
a, b, c, *rest, last = [1, 2, 3, 4, 5, 6, 7, 8, 9]

print(a)  # 1
print(b)  # 2
print(c)  # 3
print(rest)  # [4, 5, 6, 7, 8]
print(last)  # 9
