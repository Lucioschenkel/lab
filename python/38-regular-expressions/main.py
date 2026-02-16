# import regular expressions module
import re

pattern = re.compile('this')
string = 'search this inside of this text please!'

a = pattern.search(string) # a match object
b = pattern.findall(string)

search_str = "TV-1234 tv-1234"
c = re.compile(r"[a-z]+-\d+") # the "r" is for "raw" string, escape sequences are ignored, e.g. `\d`

print(c.findall(search_str))


# password checker: 
# - at least 8 characters long
# - contains letters, numbers and $%#@

pattern = re.compile(r"^[\w$%#@]{8,}$")

print(pattern.match("thisisvalid"))
