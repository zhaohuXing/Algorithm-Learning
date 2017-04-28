# 题目描述

定义栈的数据结构，请在该类型中实现一个能够得到栈最小元素的min函数。<br />
import java.util.Stack;
import java.util.Iterator;
public class Solution {

   
    Stack<Integer> stack = new Stack<>();
    
    public void push(int node) {
        
       	stack.push(node);
    }
    
    public void pop() {
        stack.pop();
    }
    
    public int top() {
        return stack.peek();
    }
    
    public int min() {
        int min = stack.peek();
       
        Iterator<Integer> iterator = stack.iterator();
        while (iterator.hasNext()){
            int data = iterator.next();
            if (min > data){
                min = data;
            }
        }
        return min;
    }
}
