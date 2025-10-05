### Map 
    
1. Map : It is a collection of indexable kays and values pairs.  
2. To reach value in map by it's key.   
3. The key should unique.  
4. That can be duplicate value.     
5. map is literal there for it create, initialize use and return.   

### why Map
1. compare two slice it use O(n) algorithm = inefficient
It need to loop over the entire slice to find matching word.   
2. Map is use o(1) algorithm = efficient.  
on average a map spend constant amount of time when looking key. for any length.  
   
### define map 
1. `var dict map[string]string`    
Explain:   
   key is string and value is string.    
2. Define empty map. because it can not have any key and value.      
`var dict[string]string{}`    
 
* default value of map is nil. initial value of map is nil.      
* if an element does not exist in the map the map return zero value depending on the element type of map.   
example1 int data type value not found return zero. 
example2 string data type value not found return "".  
example3 bool data type value not found return false.  
* Without initial map you can not assign key and value.  

### Please careful for define kay data type  
1. Data type of key is comparable means it support equality operator.  
2. Because of optimization purpose.  
3. example int.string,bool etc.  
4. slice, map, and function are not comparable.    
5. There for not use slice, map, and function as key type.  
```  
Not use.  
var broken map[[]int]int  // slice
var broken map[map[int]string]int  // map   
```
6. You can not compare map to another map even itself.  
7. you can compare map only by nil.  


### difference between slice, array, map  
1. For array you can retrive value only by using index or key.      
2. For retrive value in map by int, string, etc. but data type of key is comparable.   


### escape sequence  
1. Data type with initial value.  
`fmt.Printf("Zero Value: %#v\n", dict)`  
2. data type 
`fmt.Printf("Zero Value: %T\n", dict)`   
   