using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Collections;

public static class HexagonalGrid
{
    private static bool tile(List<int> a, List<int> b, string row, int fieldIdx, int blockedIdx, string blockedRow)
    {
        if (row.Equals("a", StringComparison.OrdinalIgnoreCase))
        {// row a
            if (fieldIdx + 1 == a.Count && a.ElementAt(fieldIdx) == 1)
            {
                if ((b.ElementAt(fieldIdx) == 1 | (blockedIdx == fieldIdx && blockedRow.Equals("b"))))
                {
                    return true;
                }
            }
            else if (a.ElementAt(fieldIdx) == 1)
            {// this tile is blocked, jump to next b or a
                if (b.ElementAt(fieldIdx) != 1 && !(fieldIdx == blockedIdx && blockedRow.Equals("b")))
                {
                    if (tile(a, b, "b", fieldIdx, fieldIdx, "a"))
                    {
                        return true;
                    }
                }
                else
                {
                    if (tile(a, b, "a", fieldIdx + 1, fieldIdx, "a"))
                    {
                        return true;
                    }
                }
            }
            else if (fieldIdx + 1 == a.Count)
            {// last index, so final check for last b
                if (b.ElementAt(fieldIdx) == 0 && !(fieldIdx == blockedIdx && blockedRow.Equals("b")))
                {
                    return true;
                }
            }
            else
            {
                if (b.ElementAt(fieldIdx) == 0 && !(fieldIdx == blockedIdx && blockedRow.Equals("b")))
                {
                    if (tile(a, b, "a", fieldIdx + 1, fieldIdx, "b"))
                    {
                        return true;
                    }
                }
                if (a.ElementAt(fieldIdx + 1) == 0 && !(fieldIdx + 1 == blockedIdx && blockedRow.Equals("a")))
                {
                    if (tile(a, b, "b", fieldIdx, fieldIdx + 1, "a"))
                    {
                        return true;
                    }
                }
            }
        }
        else
        {// row b
            if (fieldIdx + 1 == a.Count && b.ElementAt(fieldIdx) == 1)
            {
                return true;
            }
            else if (b.ElementAt(fieldIdx) == 1)
            {// this tile is blocked, jump to next a or b
                if (!(fieldIdx + 1 == blockedIdx && blockedRow.Equals("a")))
                {
                    if (tile(a, b, "a", fieldIdx + 1, fieldIdx, "b"))
                    {
                        return true;
                    }
                }
                else
                {
                    if (tile(a, b, "b", fieldIdx + 1, fieldIdx, "b"))
                    {
                        return true;
                    }
                }
            }
            else if (fieldIdx + 1 == a.Count)
            {// last index, b cant connect anymore, so failed attempt
                return false;
            }
            else
            {
                if (b.ElementAt(fieldIdx + 1) == 0 && !(fieldIdx + 1 == blockedIdx && blockedRow.Equals("b")))
                {// connect straight to next b
                    if (tile(a, b, "a", fieldIdx + 1, fieldIdx + 1, "b"))
                    {// call a tile with idx + 2
                        return true;
                    }
                }
                if (a.ElementAt(fieldIdx + 1) == 0 && !(fieldIdx + 1 == blockedIdx && blockedRow.Equals("a")))
                {// connect diagonally up to a
                    if (tile(a, b, "b", fieldIdx + 1, fieldIdx + 1, "a"))
                    {
                        return true;
                    }
                }
            }
        }
        return false;
    }

    public static string hexagonalGrid(string a, string b)
    {
        List<int> ArrA = new List<int>();
        List<int> ArrB = new List<int>();
        for (int i = 0; i < a.Length; i++)
        {
            ArrA.Add(Convert.ToInt32(a.Substring(i, 1)));
            ArrB.Add(Convert.ToInt32(b.Substring(i, 1)));
        }
        if (tile(ArrA, ArrB, "a", 0, -1, ""))
        {
            return "YES";
        }
        else
        {
            return "NO";
        }
    }
}