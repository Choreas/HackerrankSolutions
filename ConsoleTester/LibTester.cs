using System;
using System.Collections.Generic;
using System.Linq;

using HackerrankSolutions;

namespace ConsoleTester
{
    class LibTester
    {      
        static void Main(string[] args)
        {
            divisibleSumPairs();
        }

        static void hexagonalGrid()
        {
            Console.WriteLine("Hexagonal Grid");
            Console.WriteLine(HexagonalGrid.hexagonalGrid("110", "110"));
        }

        static void gradingStudents()
        {
            Console.WriteLine("gradingStudents");
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

            Console.WriteLine("Apple and Orange");
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

            Console.WriteLine("Kangaroo");
            Console.WriteLine(result);
        }

        public static void betweenTwoSets()
        {
            List<int> arr = new List<int> {2,4 };
            List<int> brr = new List<int> { 16, 32, 96 };

            int total = BetweenTwoSets.getTotalX(arr, brr);

            Console.WriteLine("Between Two Sets");
            Console.WriteLine(total);
        }

        static void breakingRecords()
        {
            int[] scores = new int[] { 10, 5, 20, 20, 4, 5, 2, 25, 1 };
            int[] result = BreakingRecords.breakingRecords(scores);
            Console.WriteLine("Breaking Records");
            Console.WriteLine(result[0] + " " + result[1]);
        }

        static void birthdayChocolate()
        {            
            List<int> s = new List<int> { 1, 2, 1, 3, 2 };
            int d = 3;
            int m = 2;
           
            int result = BirthdayChocolate.birthday(s, d, m);
            Console.WriteLine("Birthday Chocolate");
            Console.WriteLine(result);
        }

        static void divisibleSumPairs()
        {           
            int n = 6;
            int k = 3;
            int[] ar = new int[] { 1, 3, 2, 6, 1, 2 };

            int result = DivisibleSumPairs.divisibleSumPairs(n, k, ar);

            Console.WriteLine("Divisible Sum Pairs");
            Console.WriteLine(result);
        }
    }
}
