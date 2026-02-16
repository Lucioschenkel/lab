import sys
import random

start = int(sys.argv[1])
end = int(sys.argv[2])

actual = random.randint(start, end)

guess = int(input(f'guess a number between {start} and {end}: '))
if guess == actual:
    print("You're a genius!")
else:
    print('Better luck next time!')

