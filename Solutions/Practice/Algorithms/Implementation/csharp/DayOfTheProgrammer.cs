namespace Practice.Algorithms.Implementation
{
    public static class DayOfTheProgrammer
    {
        public static string dayOfProgrammer(int year)
        {
            int month = 9;
            int day = 13;

            if(year <= 1917)
            {
                if(year % 4 == 0)
                {
                    day--;
                }                
            }
            else if(year > 1918)
            {
                if((year % 4 == 0 && year % 100 != 0) | (year % 400 == 0))
                {//This is not optimal. Better to test for % 4 first, to speed things up.
                    day--;
                }
            }
            else
            {//Year 1918
                day += 13;
            }
            return day + "." + "0" + month + "." + year;
        }
    }
}
