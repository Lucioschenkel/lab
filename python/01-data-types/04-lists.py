# Ordered collection
cart = ['apples', 'rice', 'eggs']

print(cart)

# Similar to strings, we can access an element using it's index
print(cart[0])

# Return a slice of the list, starting from index 1 til the end
print(cart[1:])

# Start on 1, step = 2
print(cart[1::2]) # ['rice']

# Unlike strings, we can mutate an item at a specific index
cart[0] = 'pineapple'

print(cart)

# Because lists are stored by reference, we're basically poiting new_cart to the same
# memory address as cart. Therefore, if we mutate new_cart, cart will also be poiting to 
# the same (modified) list.
new_cart = cart
new_cart[0] = 'apples'
print(cart) # ['apples', 'rice', 'eggs']

# if we don't want that, then we can copy using list slicing
new_cart_2 = cart[:]

new_cart_2[0] = 'pineapple'

print(cart) # This will still be ['apples', 'rice', 'eggs']
