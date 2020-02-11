using System;
using System.Collections.Generic;
using System.Linq;

namespace HackerrankSolutions
{
    public static class Bon_Appétit
    {
        public static void bonAppetit(List<int> bill, int k, int b)
        {           
            bill.RemoveAt(k);
            int actual = (bill.Sum()) / 2;

            if (actual == b)
            {
                Console.WriteLine("Bon Appetit");
            }
            else
            {
                Console.WriteLine(b - actual);
            }
        }
    }
}
