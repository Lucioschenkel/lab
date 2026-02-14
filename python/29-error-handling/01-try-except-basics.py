# you can raise your own exceptions using the raise keyword
# e.g. raise ValueError('that does not look right')

while True:
    try:
        age = int(input('what is your age? '))
        10/age
    except ValueError:
        print('please enter a number')
    except ZeroDivisionError:
        print('please enter age higher than zero')
    else:
        print('thank you!')
        break
    # runs after the else block, always executed regardless of success or failure
    finally:
        print('finally done')
