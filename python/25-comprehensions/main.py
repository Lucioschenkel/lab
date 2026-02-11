# lists, sets and dictionaries

my_list = [char for char in 'hello']
my_nums = [n * 2 for n in range(0, 100)]
only_evens = [n ** 2 for n in range(0, 100) if n ** 2 % 2 == 0]

print(my_list) # ['h', 'e', 'l', 'l', 'o']
print(my_nums) # [0, 2, 4, 6, ..., 198]
print(only_evens) # [0, 4, 6, 8, ..., 10000]

# sets
my_set = {n for n in [1, 2, 3, 5, 2, 7, 8, 8]}

print(my_set) # {1, 2, 3, 5, 6, 8}

# dictionaries
simple_dictionary = {
    'a': 1,
    'b': 2
}
my_dict = {key:value**2 for key, value in simple_dictionary.items()}

print(my_dict)
