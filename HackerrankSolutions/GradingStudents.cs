using System.CodeDom.Compiler;
using System.Collections.Generic;
using System.Collections;
using System.ComponentModel;
using System.Diagnostics.CodeAnalysis;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Runtime.Serialization;
using System.Text.RegularExpressions;
using System.Text;
using System;

public static class GradingStudents
{

    /*
     * Complete the 'gradingStudents' function below.
     *
     * The function is expected to return an INTEGER_ARRAY.
     * The function accepts INTEGER_ARRAY grades as parameter.
     */

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
