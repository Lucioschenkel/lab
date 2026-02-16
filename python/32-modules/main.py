# utils is a model
import utils
# shopping is a package, which contains the shopping_cart model
import shopping.shopping_cart

# alternatively, we can import a package using the from,import syntax:
from shopping.shopping_cart import buy

# __main__ is the name of the file used as the "entrypoint", 
# i.e the one passed as an argument to the python interpreter
if __name__ == '__main__':
    print(utils)
    print(shopping.shopping_cart)

    print(buy('Rice'))

