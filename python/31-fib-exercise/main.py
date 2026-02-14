def fib(num):
    a, b =  0, 1
    for _ in range(num):
        yield a
        temp = a
        a = b
        b = temp + b

for x in fib(21):
    print(x)

