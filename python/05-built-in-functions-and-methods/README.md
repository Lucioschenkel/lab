# Built-in functions

Available globally, not scoped to a particular object.

**Examples:**

- `abs` - Return the absolute value of a number
- `len` - Returns the length of a string or list

# Built-in methods

Scoped to a particular object, invoked with `.` notation.

**Examples**:

- `'Hello {}'.format('john')` - `.format` is a method on the `str` class
- `'to be or not to be'.replace('be', 'me')` - Replaces all instances of `be` with `me`. This does not modify the original string, since strings are immutable.
