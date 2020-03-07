using System.Collections.Generic;
using System.Linq;

namespace Practice.Algorithms.Implementation
{
    public static class BetweenTwoSets
    {
        /**
        * 
        * Simple. Increment candidate by max(a) and test sequentially. Break at first invalid result.
        * Continue until candidate becomes bigger than max(b).
        * 
        **/
        public static int getTotalX(List<int> a, List<int> b)
        {
            int maxA = a.Max();
            int maxB = b.Max();
            int betweenNumbers = 0;

            for (int candidate = maxA; candidate <= maxB; candidate += maxA)
            {
                bool stillPossible = true;
                for (int i = 0; i < a.Count; i++)
                {
                    if (candidate % a.ElementAt(i) != 0)
                    {
                        stillPossible = false;
                        break;
                    }
                }
                if (stillPossible)
                {
                    for (int i = 0; i < b.Count; i++)
                    {
                        if (b.ElementAt(i) % candidate != 0)
                        {
                            stillPossible = false;
                            break;
                        }
                    }
                }
                if (stillPossible)
                {
                    betweenNumbers++;
                }
            }
            return betweenNumbers;
        }
    }
}