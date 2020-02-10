using System;
using System.Collections.Generic;
using System.Linq;

namespace ConsoleTester
{
    class LibTester
    {
        
        static void Main(string[] args)
        {
            gradingStudents();
        }

        static void hexagonalGrid()
        {
            Console.WriteLine(HexagonalGrid.hexagonalGrid("110", "110"));
        }

        static void gradingStudents()
        {
            List<int> i = new List<int> {73,67,38,33 };
            i = GradingStudents.gradingStudents(i);
            for(int ii = 0; ii < i.Count; ii++)
            {
                Console.WriteLine(i.ElementAt(ii));
            }
        }
    }
}
