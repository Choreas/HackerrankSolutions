using System.Collections.Generic;
using System.Linq;

public static class GradingStudents
{
    public static List<int> gradingStudents(List<int> grades)
    {
        List<int> roundedGrades = new List<int> { };
        for(int i = 0; i < grades.Count; i++)
        {
            int grade = grades.ElementAt(i);
            int reversedModulo = 5 - (grade % 5);
            if (grade >= 38 && reversedModulo < 3)
            {
                roundedGrades.Add(grade + reversedModulo);
            }
            else
            {
                roundedGrades.Add(grade);
            }
        }
        return roundedGrades;
    }
}
