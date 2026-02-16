# my_file = open('test.txt')

# readlines: array of lines, readline: returns lines one by one (cursor)
# read: reads the entire file contents
# print(my_file.readlines())

# my_file.close()

# Better: using "with"
with open('test.txt', mode='a') as f:
    text = f.write(':)')
    # if mode='w', this throws
    # print(f.readlines())
    print(text)
