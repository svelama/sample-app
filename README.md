# sample-app

 

### Golang
```go
package main

import "C"

import (
	"runtime"
)

func main() {}

var X string = runtime.GOOS

//export GetX
func GetX() *C.char {
	return C.CString(X)
}

//export SetX
func SetX(os *C.char) {
	X = C.GoString(os)
}

```

The dll can be generated using the following command

```
go build --buildmode=c-shared -ldflags="-s -w" -o app.dll app.go
```

### C# 

```c#

using System;
using System.Runtime.InteropServices;
using System.Text;

namespace SampleApp
{
    class Program
    {
        [DllImport("app", EntryPoint = "GetX")]
        public static extern IntPtr GetX();

        [DllImport("app", EntryPoint = "SetX")]
        extern static void SetX(string os);

        static void Main(string[] args)
        {
            // First call returns the default value of X (i.e. OS Name)
            Console.WriteLine("GetX() Returns: " + Marshal.PtrToStringAnsi(GetX()));

            // Amending the value of X
            SetX("S Win");
            Console.WriteLine("GetX() Returns: " + Marshal.PtrToStringAnsi(GetX()));

            // Amending the value of X
            SetX("S Win 2");
            Console.WriteLine("GetX() Returns: " + Marshal.PtrToStringAnsi(GetX()));

            Console.ReadLine();
        }
    }
}
```
Upon executing the C# application, it should display the output as follows,
```txt
GetX() Returns: windows
GetX() Returns: S Win
GetX() Returns: S Win 2
```

### Dev Environment
OS - Windows 11

Go Version - go version go1.19.3 windows/amd64

IDE - Go - VS Code

IDE - C# - Visual Studio Community Edition
