
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
            // First call return the default value of X (i.e. OS Name)
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
