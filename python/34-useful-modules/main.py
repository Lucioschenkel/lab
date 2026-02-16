from collections import Counter
from array import array

sentence = 'how many times does the letter "e" appear on this sentence?'

dictionary = Counter(sentence)

print(dictionary['e'])

# arrays are better for performance than lists.
# one of the tradeoffs is restricting the type of each item in the array to the typecode
arr = array('i', [1, 2, 3])
print(arr)

print(arr[0])
