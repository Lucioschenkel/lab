# Set is an unordered collection of unique items
my_set = {1, 2, 3, 4, 5, 5}

print(my_set)  # {1, 2, 3, 4, 5}

my_set.add(2)

print(my_set)  # {1, 2, 3, 4, 5}

# We cannot access indexes, it is unordered!
# print(my_set[0]) # this would fail

# Check if element is in the set
print(1 in my_set)  # True

# Regular mathematical operations in sets are available
print(my_set.difference({1, 2, 3}))  # {4, 5}

print(my_set.intersection({4, 5}))  # {4, 5}

print(my_set.union({6, 7}))  # {1, 2, 3, 4, 5, 6, 7}

print({1, 2}.issubset(my_set))  # True

print(my_set.issuperset({1, 2}))  # True

my_set.discard(1)  # Remove an element if it is a member

print(my_set)  # {2, 3, 4, 5}

# Another way to do intersection
print(my_set & {4, 5})  # {4, 5}

# Another way to do union
print(my_set | {1})  # {1, 2, 3, 4, 5}
