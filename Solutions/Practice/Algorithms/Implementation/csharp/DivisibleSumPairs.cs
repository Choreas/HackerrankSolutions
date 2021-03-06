﻿namespace Practice.Algorithms.Implementation
{
    public static class DivisibleSumPairs
    {
        public static int divisibleSumPairs(int n, int k, int[] ar)
        {
            int pairs = 0;
            for(int i = 0; i < n; i++)
            {
                for(int j = n - 1; j > i; j--)
                {
                    if((ar[i] + ar[j]) % k == 0)
                    {
                        pairs++;
                    }
                }
            }
            return pairs;
        }
    }
}
