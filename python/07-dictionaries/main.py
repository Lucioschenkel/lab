# Dictionary
dictionary = {123: [1, 2, 3], True: "hello", "something": True}
# All of these are fine as keys, but
# dictionary = { [100]: 'something' } would create an exception, because it is an unhashable type

# Keys must be unique, like any hash map like implementation, given its underlying implementation

# You can access a key using square brackets
dictionary = {"foo": "bar"}

print(dictionary["foo"])

# If you try to access a key that doesn't exist using this notation, you will get a KeyError
# print(dictionary["bar"]) # won't work

# Instead, you can use `.get`
print(dictionary.get("bar"))  # None

# `.get` accepts a "default value" as the second parameter
print(dictionary.get("bar", 55))  # 55

# Dictionaries can alternatively be created with the `dict` function
user = dict(name="John")
print(user)  # {'name': 'John'}

# The 'in' keyword can be used with dictionaries to check for a key
print("size" in user)  # False

# We can iterate of items, keys, or values using the following methods
for key in user.keys():
    print(key)  # name

for val in user.values():
    print(val)  # John

for key, val in user.items():
    print(f"{key}: {val}")  # name: John

# .copy creates a copy - not the same reference
user2 = user.copy()

user2["name"] = "Mary"

print(user)  # {'name': 'John'}

# .clear empties the dictionary
user2.clear()

print(user2)  # {}

user = {"name": "John", "age": 20}

print(user.pop("age"))  # 20: Remove age and return its value

print(user)  # {'name': 'John'}

# Similarly, `.popitem` removes an element from the dictionary (the last inserted key)
