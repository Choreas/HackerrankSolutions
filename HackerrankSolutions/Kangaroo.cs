using System;

public static class Kangaroo
{
    /**
    * Recursive solution.
    * 
    **/
    public static string kangaroo(int x1, int v1, int x2, int v2)
    {
        if(x1 == x2)
        {
            return "YES";
        }
        else if((v1 >= v2 && x1 > x2) | (v2 >= v1 && x2 > x1))
        {
            return "NO";
        }
        else
        {
            if (kangaroo(x1 + v1, v1, x2 + v2, v2).Equals("YES", StringComparison.OrdinalIgnoreCase))
            {
                return "YES";
            }
        }
   
        return "NO";
    }
}
