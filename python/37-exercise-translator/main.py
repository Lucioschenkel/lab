from translate import Translator

translator = Translator('ja')
with open('test.txt') as f:
    result = translator.translate(f.read())
    print(result)
