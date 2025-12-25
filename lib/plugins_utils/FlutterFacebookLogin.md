#flutter_facebook_login

A Flutter plugin for using the native Facebook Login SDKs on Android and iOS.

Add this to your package's pubspec.yaml file:

````
dependencies:
  flutter_facebook_login:
````
2. Install it
You can install packages from the command line:

with Flutter:

````
$ flutter pub get
````

AndroidX support

if you want to avoid AndroidX, use version 1.2.0.
for AndroidX Flutter projects, use versions 2.0.0 and up.
Installation

To get things up and running, you'll have to declare a pubspec dependency in your Flutter project. Also some minimal Android & iOS specific configuration must be done, otherise your app will crash.

###Android
This assumes that you've done the "Associate Your Package Name and Default Class with Your App" and "Provide the Development and Release Key Hashes for Your App" in the the Facebook Login documentation for Android site.

After you've done that, find out what your Facebook App ID is. You can find your Facebook App ID in your Facebook App's dashboard in the Facebook developer console.

Once you have the Facebook App ID figured out, youll have to do two things.

First, copy-paste the following to your strings resource file. If you don't have one, just create it.

<your project root>/android/app/src/main/res/values/strings.xml

````
<?xml version="1.0" encoding="utf-8"?>
<resources>
    <string name="app_name">Your App Name here.</string>

    <!-- Replace "000000000000" with your Facebook App ID here. -->
    <string name="facebook_app_id">000000000000</string>

    <!--
      Replace "000000000000" with your Facebook App ID here.
      **NOTE**: The scheme needs to start with `fb` and then your ID.
    -->
    <string name="fb_login_protocol_scheme">fb000000000000</string>
</resources>
````
Then you'll just have to copy-paste the following to your Android Manifest:

<your project root>/android/app/src/main/AndroidManifest.xml

````
<meta-data android:name="com.facebook.sdk.ApplicationId"
    android:value="@string/facebook_app_id"/>

<activity android:name="com.facebook.FacebookActivity"
    android:configChanges=
            "keyboard|keyboardHidden|screenLayout|screenSize|orientation"
    android:label="@string/app_name" />

<activity
    android:name="com.facebook.CustomTabActivity"
    android:exported="true">
    <intent-filter>
        <action android:name="android.intent.action.VIEW" />
        <category android:name="android.intent.category.DEFAULT" />
        <category android:name="android.intent.category.BROWSABLE" />
        <data android:scheme="@string/fb_login_protocol_scheme" />
    </intent-filter>
</activity>
````
Done!

###iOS
This assumes that you've done the "Register and Configure Your App with Facebook" step in the the Facebook Login documentation for iOS site. (Note: you can skip "Step 2: Set up Your Development Environment" and "Step 5: Connect Your App Delegate").

After you've done that, find out what your Facebook App ID is. You can find your Facebook App ID in your Facebook App's dashboard in the Facebook developer console.

Once you have the Facebook App ID figured out, then you'll just have to copy-paste the following to your Info.plist file, before the ending </dict></plist> tags. (NOTE: If you are using this plugin in conjunction with for example google_sign_in plugin, which also requires you to add CFBundleURLTypes key into Info.plist file, you need to merge them together).

<your project root>/ios/Runner/Info.plist
````
<key>CFBundleURLTypes</key>
<array>
    <!--
    <dict>
    ... Some other CFBundleURLTypes definition.
    </dict>
    -->
    <dict>
        <key>CFBundleURLSchemes</key>
        <array>
            <!--
              Replace "000000000000" with your Facebook App ID here.
              **NOTE**: The scheme needs to start with `fb` and then your ID.
            -->
            <string>fb000000000000</string>
        </array>
    </dict>
</array>

<key>FacebookAppID</key>

<!-- Replace "000000000000" with your Facebook App ID here. -->
<string>000000000000</string>
<key>FacebookDisplayName</key>

<!-- Replace "YOUR_APP_NAME" with your Facebook App name. -->
<string>YOUR_APP_NAME</string>

<key>LSApplicationQueriesSchemes</key>
<array>
    <string>fbapi</string>
    <string>fb-messenger-share-api</string>
    <string>fbauth2</string>
    <string>fbshareextension</string>
</array>
````
A sample of a complete Info.plist file can be found here.

Done!