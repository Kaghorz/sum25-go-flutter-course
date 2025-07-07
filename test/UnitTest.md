# unit testing
How can you ensure that your app continues to work as you add more features or change existing functionality? By writing tests.

Unit tests are handy for verifying the behavior of a single function, method, or class. The test package provides the core framework for writing unit tests, and the flutter_test package provides additional utilities for testing widgets.

## This recipe demonstrates the core features provided by the test package using the following steps:

````
    1.Add the test or flutter_test dependency.
    2.Create a test file.
    3.Create a class to test.
    4.Write a test for our class.
    5.Combine multiple tests in a group.
    Run the tests.
````

### Add the test dependency
````
dev_dependencies:
  test: <latest_version>
````

### Create a test file

In general, test files should reside inside a test folder located at the root of your Flutter application or package. Test files should always end with _test.dart, this is the convention used by the test runner when searching for tests.

### Create a class to test

Next, you need a “unit” to test. Remember: “unit” is another name for a function, method, or class.

### Write a test for our class

Inside test.dart file, write the first unit test. Tests are defined using the top-level test function, and you can check if the results are correct by using the top-level expect function. Both of these functions come from the test package.

### Combine multiple tests in a group

If you have several tests that are related to one another, combine them using the group function provided by the test package.

### Run the tests
Now that you have a class with tests in place, you can run the tests.

#### Run tests using IntelliJ
    The Flutter plugins for IntelliJ and VSCode support running tests. This is often the best option while writing tests because it provides the fastest feedback loop as well as the ability to set breakpoints.
    ``````
    $Open test.dart file
    $Select the Run menu
    $Click the Run 'tests in your dart file' option
    $Alternatively, use the appropriate keyboard shortcut for your platform.
    ``````

#### Run tests in a terminal
    You can also use a terminal to run the tests by executing the following command from the root of the project:

    $ flutter test test/counter_test.dart



