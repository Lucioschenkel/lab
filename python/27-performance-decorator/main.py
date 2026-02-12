from time import time

def performance(fn):
    def wrapper(*args, **kwargs):
        t1 = time()
        result = fn(*args, **kwargs)
        t2 = time()
        print(f'it took {t2-t1}s')
        return result

    return wrapper

@performance
def long_time():
    total = 0
    for i in range(100_000_000):
        total += i*5
    return total

long_time()
