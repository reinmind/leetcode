package lc227;

import java.util.Stack;

/**
 * @author xiang.zhang
 * @since 2024/2/7
 */
public class Solution {
    /**
     * 11 + (2 + (3 + 5 * 2 - 1) / 3 ) * 2 + 100 = 123
     * @param s
     * @return
     */
    public int calculate(String s) {
        Stack<Character> op = new Stack<>();
        Stack<Integer> num = new Stack<>();
        final char[] exp = s.toCharArray();
        for(int i = 0;i<exp.length;i++){
            if(exp[i] == ' '){
                continue;
            }
            if(exp[i] == '('){
                op.push(exp[i]);
            }
            if(exp[i] >= '0' && exp[i] <= '9' ){
                int t = i;
                while(t < s.length() && exp[t] >= '0' && exp[t] <= '9'){
                    t++;
                }
                int tmp = Integer.parseInt(s.substring(i,t));
                num.push(tmp);

                if(!op.isEmpty() && op.peek() == '*'){
                    op.pop();
                    final Integer pop = num.pop();
                    final Integer pop2 = num.pop();
                    num.push(pop*pop2);
                }
                if(!op.isEmpty() && op.peek() == '/'){
                    op.pop();
                    final Integer pop = num.pop();
                    final Integer pop2 = num.pop();
                    num.push(pop2/pop);
                }
                i = t - 1;
                continue;
            }
            if(exp[i] == '*' || exp[i] == '-' || exp[i] == '+' || exp[i] == '/'){
                if(!op.isEmpty() && (op.peek() == '+' || op.peek() == '-') && (exp[i] == '-' || exp[i] == '+')){
                    char opc = op.pop();
                    final Integer pop1 = num.pop();
                    final Integer pop2 = num.pop();
                    num.push(simpleCalc(opc,pop2,pop1));
                }
                op.push(exp[i]);

            }
            if(exp[i] == ')'){
                Character opc = op.pop();
                while(opc != '('){
                    final Integer pop1 = num.pop();
                    final Integer pop2 = num.pop();
                    num.push(simpleCalc(opc,pop2,pop1));
                    opc = op.pop();
                }
            }
        }
        while(!op.isEmpty()){
            char opc = op.pop();
            final Integer pop1 = num.pop();
            final Integer pop2 = num.pop();
            num.push(simpleCalc(opc,pop2,pop1));
        }
        return num.peek();
    }

    int simpleCalc(char op,int a,int b){
        if(op == '-'){
            return a - b;
        }
        if(op == '+'){
            return a + b;
        }
        if (op == '/'){
            return a/b;
        }
        if(op == '*'){
            return a*b;
        }
        return 0;
    }

    public static void main(String[] args) {
        final Solution solution = new Solution();
        final int calculate = solution.calculate("11 + (2 + (3 + 5 * 2 - 1) / 3 ) * 2 - 3");
        final int calculate2 = solution.calculate("1-1+1");
        System.out.println(calculate);
        System.out.println(calculate2);
    }
}
