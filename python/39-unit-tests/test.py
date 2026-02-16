import unittest
import main

class TestMain(unittest.TestCase):
    def setUp(self) -> None:
        # runs before each test, like "beforeEach" in Jest and Vitest
        print('about to test a function')
        return super().setUp()
    def test_do_stuff(self):
        num = 10
        result = main.do_stuff(num)
        self.assertEqual(result, 15)

if __name__ == "__main__":
    unittest.main()
