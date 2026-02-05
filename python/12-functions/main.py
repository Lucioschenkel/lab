# def is used to define a function
# parameters
def say_hello(name='Darth Vader', emoji = '🚀'):
    print(f'hello {name} {emoji}')

# once defined, we can call a function
# (positional) arguments
say_hello('john', '🚀')
say_hello('mary', '🚀')
say_hello('joseph', '🚀')

# keyword arguments
say_hello(emoji='🚀', name='Bibi')

# uses default parameters
say_hello()

# functions with return values
def sum2(num1, num2):
    return num1 + num2

total = sum2(10, 5)
print(total)

# doc strings document a function
def test(a):
    '''
    Info: this function tests and prints param a
    '''
    print(a)

# use mouse hover, or whatever the LSP is hooked up to, to read the docstring
test('!!!')

# Or, use the 'help' builtin function. This opens as a "modal" window in a tty terminal
# help(test)

# *args and **kwargs
def super_func(*args, **kwargs):
    # prints space-separated arguments
    print(*args)
    print(kwargs)
    total = 0
    for item in kwargs.values():
        total += item
    return sum(args) + total

# accepts any number of positional arguments
# because of **kwargs, also accepts any number of keyword arguments
print(super_func(1, 2, 3, 4, 5, num1=5, num2=10))

# Rule: named paramenters, *args, default parameters, **kwargs
