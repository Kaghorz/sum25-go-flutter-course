# An introduction to widget testing

***To test widget classes, you need a few additional tools provided by the flutter_test package, which ships with the Flutter SDK.***
following tools for testing widgets:

## 1. Add the flutter_test dependency
Before writing tests, include the flutter_test dependency in the dev_dependencies section of the pubspec.yaml file.
````
dev_dependencies:
  flutter_test:
    sdk: flutter
````

## 2. Create a widget to test
Next, create a widget for testing. For this recipe, create a widget that displays a title and message.

## 3. Create a testWidgets test
With a widget to test, begin by writing your first test. Use the testWidgets() function provided by the flutter_test package to define a test. The testWidgets function allows you to define a widget test and creates a WidgetTester to work with.

## 4. Build the widget using the WidgetTester
Next, build your widget inside the test environment by using the pumpWidget() method provided by WidgetTester. The pumpWidget method builds and renders the provided widget.

***Notes about the pump() methods***
After the initial call to pumpWidget(), the WidgetTester provides additional ways to rebuild the same widget. This is useful if you’re working with a StatefulWidget or animations.

## 5. Search for our widget using a Finder
Use the top-level find() method provided by the flutter_test package to create the Finders. Since you know you’re looking for Text widgets, use the find.text() method.

## 6. Verify the widget using a Matcher
Finally, verify widgets which is appear on screen using the Matcher constants provided by flutter_test.
Matcher classes are a core part of the test package, and provide a common way to verify a given value meets expectations.

## 7. View complete example in test/widget_test.dart.
