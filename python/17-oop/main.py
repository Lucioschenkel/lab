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
            # Python does not support access modifiers
            # By convention, users should treat attributes and methods 
            # prefixed by an underscore as private
            self._address = "Secret hideout"

    # Method
    def run(self):
        print('run')
        return 'done'

    def shout(self):
        print(f'My name is {self.name}')

    @classmethod
    def adding_things(cls, num1, num2): # cls is the class, in this case 'PlayerCharacter'
        return num1 + num2

    @staticmethod
    def adding_things2(num1, num2): # Unlike @classmethod, we don't have access to the class
        return num1 + num2

player1 = PlayerCharacter('John', 30)

print(player1.membership)
print(PlayerCharacter.membership)
print(player1.name)

player1.shout()

# print(player1.adding_things(2,3))
print(PlayerCharacter.adding_things(2,3))
print(PlayerCharacter.adding_things2(2, 3))
