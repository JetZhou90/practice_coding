#pragma once
#ifdef EXPORT_WIN32POJ
#define EXPORTS_DEMO _declspec(dllexport)
#else
#define EXPORTS_DEMO _declspec(dllimport)
#endif
namespace MathFuncs
{
    // This class is exported from the testdll.dll  
    class MyMathFuncs
    {
    public:
        // Returns a + b  
        static EXPORTS_DEMO double Add(double a, double b);

        // Returns a - b  
        static EXPORTS_DEMO double Subtract(double a, double b);

        // Returns a * b  
        static EXPORTS_DEMO double Multiply(double a, double b);

        // Returns a / b  
        // Throws const std::invalid_argument& if b is 0  
        static EXPORTS_DEMO double Divide(double a, double b);
    };
}