class PlayerCharacter:
    # Class Object Attribute
    membership = True # this is static

    # This is constructor. It is called automatically 
    # when an object is instantiated
    def __init__(self, name, age) -> None:
        if self.membership:
            # Attributes
            self.name = name
            self.age = age

    # Method
    def run(self):
        print('run')
        return 'done'

    def shout(self):
        print(f'My name is {self.name}')

player1 = PlayerCharacter('John', 30)

print(player1.membership)
print(PlayerCharacter.membership)
print(player1.name)

player1.shout()
