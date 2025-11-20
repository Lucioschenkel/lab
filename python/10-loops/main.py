# a list is an iterable
for item in [1, 2, 3, 4, 5]:
    print(item)

# same with a tuple
for item in (1, 2, 3, 4, 5):
    print(item)

# with dictionaries
my_dict = {"foo": "bar", "bar": "baz"}

for key, val in my_dict.items():
    print(key, val)

for key in my_dict.keys():
    print(key)

for i, char in enumerate("ztm"):
    # print the index and the character
    print(i, char)

for i, item in enumerate(list(range(100))):
    if item == 50:
        print(f"the index of 50 is {i}")
