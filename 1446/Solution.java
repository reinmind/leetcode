class Solution {
    public static void main(String[] args) {
        Solution s = new Solution();
        int e = s.maxPower("cc");
        System.out.println(e);
    }
    public int maxPower(String s) {
        int e = 1,i = 0;
        char[] ca = s.toCharArray();
        char current = ca[0];
        for(int j = 1;j < ca.length;j++){
            if(ca[j] == current){
                e++;
            }else{
                current = ca[j];
                i = Math.max(i,e);
                e = 1;
            }
        }
        return Math.max(i,e);
    }
}