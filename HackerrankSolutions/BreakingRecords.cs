
public static class BreakingRecords
{
    /**
    * 
    * Super simple. Set highscore and lowscore in a loop after each new record.
    * Start both records as first game's score. Count the times u had to set them.
    * 
    **/
    public static int[] breakingRecords(int[] scores)
    {
        int[] records = new int[] { 0, 0 };
        int highscore = scores[0];
        int lowScore = scores[0];
        for(int i = 0; i < scores.Length; i++)
        {
            int score = scores[i];
            if(score < lowScore)
            {
                lowScore = score;
                records[1]++;
            }
            else if(score > highscore)
            {
                highscore = score;
                records[0]++;
            }
        }
        return records;
    }
}
