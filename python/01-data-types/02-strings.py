a = 'this is a string'

b = "this is also a string"

long_string = '''
WOW
0 0
---
'''

print(long_string)

first_name = 'john'

last_name = 'doe'

# String concatenation
full_name = first_name + ' ' + last_name

# Type conversion
print(str(100)) # "100" as a string

# Escape sequences
print("It's \"kind of\" sunny.")

# Formatted strings
name = 'Johnny'
age = 50

# On Python 3
print(f'Hi, {name}. You are {age} years old.')

# On Python 2
print('Hi, {1}. You are {0} years old.'.format(age, name))

print('Hi, {new_name} You are {new_age} years old.'.format(new_name='Sally', new_age=100))

# Indexing - zero-based
selfish = "me me me"

print(selfish[0]) # m

# We can specify start:end to get a slice of the string. 'end' is exclusive
print(selfish[0:4]) # me m

# We can also specify a "stepover", the default is one
print(selfish[0:8:3]) # mmm

# Grap everythin from 1 till the end of the string
print(selfish[1:])

my_numbers = '01234567'
# Neat little trick to reverse the string
print(my_numbers[::-1])

# my_numbers[0] = '8' Throws an exception, as strings are immutable. You can, however, assign it a new string e.g.: my_numbers = '8'
