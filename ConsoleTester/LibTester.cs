using System;
using System.Collections.Generic;
using System.Linq;

namespace ConsoleTester
{
    class LibTester
    {      
        static void Main(string[] args)
        {
            kangaroo();
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

        static void appleAndOrange()
        {
            int s = 7;
            int t = 11;
            int a = 5;
            int b = 15;
            int[] apples = new int[] { -2, 2, 1 };
            int[] oranges = new int[] { 5, -6 };            

            AppleAndOrange.countApplesAndOranges(s, t, a, b, apples, oranges);
        }

        static void kangaroo()
        {            
            string[] x1V1X2V2 = new string[] {"0","3","4","2" };

            int x1 = Convert.ToInt32(x1V1X2V2[0]);
            int v1 = Convert.ToInt32(x1V1X2V2[1]);
            int x2 = Convert.ToInt32(x1V1X2V2[2]);
            int v2 = Convert.ToInt32(x1V1X2V2[3]);

            string result = Kangaroo.kangaroo(x1, v1, x2, v2);

            Console.WriteLine(result);
        }
    }
}
