class User:
    def __init__(self, email):
        self.email = email 

    def sign_in(self):
        print('logged in')

# This indicates that Wizard inherits from User
class Wizard(User):
    def __init__(self, email, name, power):
        # This call also be User.__init__(self,email)
        super().__init__(email)
        self.name = name
        self.power = power

    def attack(self):
        print(f'Attacking with power of {self.power}')

class Archer(User):
    def __init__(self, name, num_arrows) -> None:
        self.name = name
        self.num_arrows = num_arrows

    def attack(self):
        self.num_arrows -= 1
        print(f'Attacking with arrows: arrows left - {self.num_arrows}')

wizard1 = Wizard('wizard@example.com', 'Merling', 50)
print(wizard1.email)
print(dir(wizard1)) # prints the list of available methods and attributes
# wizard1.sign_in()
# wizard1.attack()

archer1 = Archer('Robin', 100)
# archer1.sign_in()
# archer1.attack()

# We can check whether an object is an instance of a class or not using isinstance
# print(isinstance(wizard1, Wizard)) # True
# print(isinstance(wizard1, User)) # True
# print(isinstance(wizard1, object)) # True! In Python, everything inherits from object!
