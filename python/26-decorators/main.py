# decorators use the '@' sign, followed by a name, e.g: @classmethod and @staticmethod

def hello():
    print('helloooo')

greet = hello
del hello

# greet still points to the memory location of the function, so this works
greet()

# this wouldn't
# hello()

# passing functions as arguments is possible
def func1(func):
    func()

def func2():
    print('hellooo')

func1(func2)

# higher order functino (HOF) -> a function that returns another function
def outer():
    def inner():
        print('hello again')

    return inner

# I can "invoke it twice", to run the inner function
outer()()

# decorators build on top of these language features
def my_decorator(func):
    def wrap_func(*args, **kwargs):
        print('***')
        func(*args, **kwargs)
        print('***')

    return wrap_func

@my_decorator
def hello2(x):
    print(f'hello {x}')

hello2('john doe')
