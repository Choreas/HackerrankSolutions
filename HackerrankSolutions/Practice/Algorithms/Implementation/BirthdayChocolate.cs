using System;
using System.Collections.Generic;
using System.Text;

namespace Practice.Algorithms.Implementation
{
    /**
    * Loop through s, for each Element[s] add the next m Elements[s] and check if result == d.
    **/
    public static class BirthdayChocolate
    {
        public static int birthday(List<int> s, int d, int m)
        {
            int numberOfBars = 0;

            for(int i = 0; i < s.Count; i++)
            {
                int tempResult = 0;
                for(int ii = i; ii <= m + i - 1; ii++)
                {
                    if(ii >= s.Count)
                    {
                        break;
                    }
                    tempResult += s[ii];
                }
                if(tempResult == d)
                {
                    numberOfBars++;
                }
            }
            return numberOfBars;
        }
    }
}
