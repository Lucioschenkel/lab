from datetime import date

today = date.today()

current_year = today.year

while True:
    birth_year = input('what year were you born? ')


    # int(birth_year) throws an error if birth_year is not a number
    try:
        age = current_year - int(birth_year)

        print(f'Your age is {age}')
        break
    except Exception as e:
        print('Please enter a valid year.') 
