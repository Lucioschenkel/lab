def generator_function(num):
    for i in range(num):
        # yield "pauses" execution, until next is invoked for this generator
        yield i*2

g = generator_function(100)
# this tells the generator to resume execution
next(g)
next(g)
print(next(g))

# Custom range func
class MyGen:
    current = 0
    def __init__(self, first, last):
        self.first = first
        self.last = last

    def __iter__(self):
        return self

    def __next__(self):
        if MyGen.current < self.last:
            num = MyGen.current
            MyGen.current += 1
            return num
        raise StopIteration

gen = MyGen(0, 100)
for i in gen:
    print(i)
