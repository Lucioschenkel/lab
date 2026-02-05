a = 'helloooooooooo'

if (n := len(a)) > 10:
    # n now has the value of len(a)
    print(f"too long {n} elements")

while (n := len(a)) > 1:
    print(n)
    a = a[:-1]

print(a)
