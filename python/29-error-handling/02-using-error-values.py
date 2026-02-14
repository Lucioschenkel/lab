def div(num1, num2):
    try:
        return num1 / num2
    # use the error value in the program, combine multiple exception types
    except (TypeError, ZeroDivisionError) as err:
        print(err)

print(sum('1', 2))
