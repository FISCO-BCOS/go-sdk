## Develop a FISCO BCOS application for iOS
Please make sure you run this on a Mac.
### Step 1: Compile and get the sdk library
You can download the library ``FiscoBcosIosSdk.framework`` directly from the release page, or you can compile it by yourselves follow step 1.
#### Prepare environment
Copy the following files
* ``libproc.h``
* ``net/route.h``
* ``sys/kern_control.h``
* ``sys/proc_info.h``


from folder ``/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/usr/include`` 

to folders

``/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS.sdk/usr/include`` 

and ``/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulatorOS.sdk/usr/include``

#### Compile
```bash
# under the folder of go-sdk
$ export CGO_LDFLAGS_ALLOW=".*"
$ gomobile bind -target=ios -o FiscoBcosIosSdk.framework ./mobile/iOS 
```
Now you get the iOS SDK ``FiscoBcosIosSdk.framework`` of FISCO BCOS blockchain.

### Step 2: Generate objc wrapper contract
#### Prepare solidity contract
```bash
// Under go-sdk folder
$ mkdir typetest && cd typetest
```
DataTypeTest.sol
```javascript
pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;

contract DataTypeTest {
    
    // Basic type
    
    bool b;
    
    int i;
    int8 i8;
    int16 i16;
    int32 i32;
    int64 i64;
    int256 i256;
    
    
    uint u;
    uint8 u8;
    uint16 u16;
    uint32 u32;
    uint64 u64;
    uint256 u256;
    
    string s;
    bytes1 b1;
    bytes5 b5;
    bytes32 b32;
    
    bytes bs;


    enum TestEnum { First, Second, Third}
    TestEnum te;
    
    address ad;
    
    // Array type
   int[] iar;
   int64[] i64ar;
   int256[] i256ar;
   bytes32[] b32ar;
   
   struct Donation {  
          int32 value;
          int32 date;
    }  
    
    Donation d;
    
    
    function storeBool(bool boolArg) public {
        b = boolArg;
    }
    
    function storeInt(int intArg, int8 int8Arg, int16 int16Arg, int32 int32Arg, int64 int64Arg) public {
        i = intArg;
        i8 = int8Arg;
        i16 = int16Arg;
        i32 = int32Arg;
        i64 = int64Arg;
    }
    
    function storeBigInt(int256 int256Arg) public{
        i256 = int256Arg;
    }
    
    function storeUint(uint uintArg, uint8 uint8Arg, uint16 uint16Arg, uint32 uint32Arg, uint64 uint64Arg) public {
        u = uintArg;
        u8 = uint8Arg;
        u16 = uint16Arg;
        u32 = uint32Arg;
        u64 = uint64Arg;
    }
    
    function storeUBigInt(uint256 uint256Arg) public{
        u256 = uint256Arg;
    }
    
    function storeStr(string strArg) public{
        s = strArg;
    }
    
    function storeFixedBytes(bytes1 byte1Arg,bytes5 byte5Arg,bytes32 byte32Arg) public{
        b1 = byte1Arg;
        b5 = byte5Arg;
        b32 = byte32Arg;
    }
    
    function storeBytes(bytes bytesArg) public{
        bs = bytesArg;
    }
    
    function storeEmum(TestEnum teArg) public{
        te = teArg;
    }

    function storeAddress(address adArg) public{
        ad = adArg;
    }
    
    function storeIntArray(int[] iarArg, int64[] i64arArg, int256[] i256arArg) public{
        iar = iarArg;
        i64ar = i64arArg;
        i256ar = i256arArg;
    }
    
    function storeByteArray(bytes32[] b32arArg) public{
        b32ar = b32arArg;
    }

    function retrieve() public view returns (uint256){
        return u256;
    }
    
    function retrieveArray() public view returns (bytes32[],int64[]){
        return (b32ar,i64ar);
    }
    
    function storeStruct(Donation structArg) public{
        d = structArg;
    }
}
```
#### Compile contract
```bash
# Under go-sdk/typetest folder
# Download compiler
bash ../tools/download_solc.sh -v 0.4.25
# Compile contract
./solc-0.4.25 --bin --abi -o ./ ./DataTypeTest.sol
```

#### Generate objc wrapper contract
```bash
# Build code generate tool
$ go build ../cmd/abigen
# Generate objc wrapper contract
$ ./abigen --bin DataTypeTest.bin --abi DataTypeTest.abi --pkg datatypetest --type DataTypeTest --out ./DataTypeTest.m --lang objc
```

### Step 4. Build an iOS application
#### Import iOS SDK 
Create an iOS application in Xcode. Copy and import the FiscoBcosIosSdk.framework as well as your generated objc wrapper code ``DataTypeTest.h`` and ``DataTypeTest.m`` into your project.

#### Use the SDK and objc wrapper contract
Import ``DataTypeTest.h`` in your project.
```objc
NSString *path = [NSBundle mainBundle].bundlePath;
NSString *endpoint = @"localhost:8170";
NSString *keyFile = [NSString stringWithFormat:@"%@/%@", path, @"key.pem" ];
// Build sdk
MobileBuildSDKResult *result = MobileBuildSDKWithParam(path,keyFile, 1, endpoint, false,1,false );
NSLog(@"Connect result: %@",result.information);
if (result.isSuccess == false){
    return;
}

// Deploy contract
DataTypeTest *contract = [DataTypeTest new];
MobileDeployContractResult *deployResult = [contract deploy];
if (deployResult.errorInfo.length != 0){
    NSLog(@"Deploy failed: %@",deployResult.errorInfo);
}else{
    NSLog(@"Deploy succeed, contract address: %@",deployResult.address);
}
```