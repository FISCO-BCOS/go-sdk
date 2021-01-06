## Build iOS SDK
Please make sure you run this on a Mac.
### Step 1: Prepare environment
Copy the following files
* ``libproc.h``
* ``net/route.h``
* ``sys/kern_control.h``
* ``sys/proc_info.h``


from folder ``/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/usr/include`` 

to folders

``/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS.sdk/usr/include`` 

and ``/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulatorOS.sdk/usr/include``

### Step 2: Compile
```
# under the folder of go-sdk
$ export CGO_LDFLAGS_ALLOW=".*"
$ gomobile bind -target=ios ./fiscoBcosMobile
```