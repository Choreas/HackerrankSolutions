using System.Collections.Generic;
using System;

namespace Practice.Algorithms.Implementation
{
    public static class MigratoryBirds
    {
         /**
         * 
         * Wanna work with the basics? Just count.
         * 
         **/
        public static int migratoryBirds(List<int> arr)
        {
            arr.Sort();
            int n = arr.Count;
            int bestBird = arr[0];
            int currentBird = arr[0];
            int bestBirdCount = 0;
            int currentBirdCount = 0;
            
            for (int i = 0; i < n; i++)
            {                
                if (arr[i] == currentBird)
                {
                    currentBirdCount++;
                    if(i == n - 1 && currentBirdCount > bestBirdCount)
                    {
                        bestBird = currentBird;
                    }
                }
                else if (currentBirdCount > bestBirdCount)
                {
                    bestBird = currentBird;
                    bestBirdCount = currentBirdCount;
                    currentBird = arr[i];
                    currentBirdCount = 1;
                }
                else
                {
                    currentBird = arr[i];
                    currentBirdCount = 1;
                }
            }
            return bestBird;
        }
    }
}
