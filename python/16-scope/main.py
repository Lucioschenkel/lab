total = 0

def inc():
    # can't access global variable directly!
    global total
    total += 1 
    return total

print(inc())

def outer():
    x = "local"
    def inner():
        nonlocal x
        # because we used nonlocal, we're modifying the variable in the outer scope
        x = "nonlocal"
        print("inner:", x)

    inner()
    print("outer:",x)

outer()
