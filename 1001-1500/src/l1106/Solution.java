package l1106;

import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

public class Solution {
    public boolean parseBoolExpr(String expression) {

        return eval(expression);
    }

    boolean eval(String expression) {
        final char c = expression.charAt(0);
        if (c == 'f') {
            return false;
        }
        if (c == 't') {
            return true;
        }
        final String subExpression = expression.substring(2, expression.length() - 1);
        Stack<Integer> op = new Stack<>();
        List<Boolean> results = new ArrayList<>();

        for (int i = 0; i < subExpression.length(); i++) {
            if (subExpression.charAt(i) == '(') {
                op.push(i);
            }
            if (subExpression.charAt(i) == ')') {
                final Integer pop = op.pop();
                if (op.empty()) {
                    results.add(eval(subExpression.substring(pop - 1, i + 1)));
                }
            }
            if(subExpression.charAt(i) == 't' && op.empty()){
                results.add(true);
            }
            if(subExpression.charAt(i) == 'f' && op.empty()){
                results.add(false);
            }
        }
        boolean init = true;
        if (c == '&') {
            init = true;
            for(Boolean b:results){
                init = init & b;
            }
        }
        if (c == '|') {
            init = false;
            for(Boolean b:results){
                init = init || b;
            }
        }
        if(c == '!'){
            for(Boolean b:results){
                init = !b;
            }
        }
        return init;
    }

    // !(&(f,t))
    public static void main(String[] args) {
        final Solution solution = new Solution();
        final boolean b = solution.parseBoolExpr("!(&(f,t))");
        System.out.println(b);
    }
}
