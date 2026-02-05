# probably most performant
def highest_even(l):
    highest = -1
    for item in l:
        if item % 2 == 0 and item > highest:
            highest = item
    return highest

# using list comprehension
def highest_even2(l):
    # creates a list of all even numbers in l
    # then, uses the builtin max function to return the highest
    return max([x for x in l if x % 2 == 0])

print(highest_even([10, 2, 3, 4, 8, 11]))
print(highest_even2([10, 2, 3, 4, 8, 11]))
