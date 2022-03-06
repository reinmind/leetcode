class Solution {
    int nc,oc,width,height;
    public int[][] floodFill(int[][] image, int sr, int sc, int newColor) {
        nc = newColor;
        oc = image[sr][sc];
        if(nc == oc) {
            return image;
        }
        width = image[0].length;
        height = image.length;
        dfs(image,sr,sc);
        return image;
    }
    public void dfs(int [][]image,int sr,int sc){
        if(sr < 0 || sr >= height || sc < 0 || sc >= width){
            return ;
        }
        if(image[sr][sc] == oc){
            image[sr][sc] = nc;
        }else{
            return;
        }
        // move left
        dfs(image,sr,sc-1);
        // move right
        dfs(image,sr,sc+1);
        // move up
        dfs(image,sr-1,sc);
        // move down
        dfs(image,sr+1,sc);
    } 
    public static void main(String... args){
        int[][] image = new int[][]{{0,0,0},{0,1,1}};
        Solution s = new Solution();
        s.floodFill(image,1,1,1);
        for(int[] x : image){
            for(int y: x){
                System.out.print(y + " ");
            }
            System.out.println();
        }
    }
}