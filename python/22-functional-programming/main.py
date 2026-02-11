from functools import reduce

def multiply_by_two(item):
    return item * 2

def is_odd(item):
    return item % 2 != 0

def sum_reduce(acc, curr):
    return acc + curr

my_list = [1, 2, 3]
doubled = list(map(multiply_by_two, my_list)) # map returns a map object, must be converted to list

print(doubled)
print(my_list) # [1, 2, 3]

my_list = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
odd = list(filter(is_odd, my_list)) # filter returns a filter object, must be converted to a list

print(odd)

print(list(zip(my_list, odd))) # combines two or more lists combining elements at the same position


print(reduce(sum_reduce,my_list, 0)) # reduces a list down to a single value

# lambda functions in python use the lambda keyword, e.g.:
print(list(map(lambda x: x * 2, [1, 2, 3]))) # will return [2, 4, 6]
