public class MergeAlternately {
    public String mergeAlternately(String word1, String word2) {
        final char[] w1 = word1.toCharArray();
        final char[] w2 = word2.toCharArray();
        int min = Math.min(w1.length,w2.length);
        char[] w3 = new char[min*2];
        for(int i = 0;i < min;i++){
            w3[i*2] = w1[i];
            w3[i*2+1] = w2[i];
        }
        String result = new String(w3);
        if(w2.length > w1.length){
            result += word2.substring(min);
        }else{
            result += word1.substring(min);
        }
        return result;
    }
}
