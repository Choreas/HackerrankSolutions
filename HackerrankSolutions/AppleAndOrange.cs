using System.Linq;
using System;

public static class AppleAndOrange
{
    public static void countApplesAndOranges(int s, int t, int a, int b, int[] apples, int[] oranges)
    {
        int applesOnMyLawn = 0;
        int orangesOnMyLawn = 0;

        for(int i = 0; i < apples.Length; i++)
        {
            int applePosition = a + apples.ElementAt(i);
            if (applePosition >= s && applePosition <= t)
            {
                applesOnMyLawn++;
            }
        }
        for (int i = 0; i < oranges.Length; i++)
        {
            int orangePosition = b + oranges.ElementAt(i);
            if (orangePosition >= s && orangePosition <= t)
            {
                orangesOnMyLawn++;
            }
        }

        Console.WriteLine(applesOnMyLawn);
        Console.WriteLine(orangesOnMyLawn);
    }
}
