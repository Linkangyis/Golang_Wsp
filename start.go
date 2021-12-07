package main

import (
        "fmt"
        "strings"
        "io/ioutil"
        "regexp"
        "encoding/base64"
        "strconv"
        "math/rand"
        "time"
        "os"
        "flag"
)

var (
	file string
)
func init() {
	flag.StringVar(&file, "f", "./start.wsp", "源代码目录")
	flag.Parse()
}

func GetFileContentAsStringLines(filePath string) ([]string) {
	result := []string{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return result
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	return result
}

func RandNum(min , max int) int {
    var Rander = rand.New(rand.NewSource(time.Now().UnixNano()))

    const letterString = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    const numLetterString = "0123456789"

	return Rander.Intn(max - min + 1) + min
}


var variable_table = map[string]string{}
var function_code = map[string]string{}
var function_variable = map[string]string{}
var return_code_eval string



func var_array_s(variable_code string){
if(len(variable_table[variable_code])!=0){
if string(variable_table[variable_code][len(variable_table[variable_code])-1])==string(";"){
    re := regexp.MustCompile(`\$(.*?)\&`)
    code_tmp_s_r := re.FindAllString(variable_table[variable_code],-1)
    code_for:=strings.Replace(string(code_tmp_s_r[0]),"&","",1)
                
    re = regexp.MustCompile(`\&(.*?)\;`)
    code_tmp_s := re.FindAllString(variable_table[variable_code],-1)
    for i:=0;i<=len(code_tmp_s)-1;i++{
        code_for_num:=strings.Replace(string(code_tmp_s[i]),"&","",1)
        code_for_num=strings.Replace(string(code_for_num),";","",1)
        variable_table[variable_code+"["+code_for_num+"]"]=variable_so(code_for+"["+code_for_num+"]")
        var_array_s(variable_code+"["+code_for_num+"]")
        }
        
    }
}
}

func variable_code(variable []string){    //变量存储
variable_len:=len(variable)-1
variable_name := "$"
    if variable_len != 0{
    //fmt.Println(variable[0])
        for i := variable_len; i >= 0; i-- {
            array_len := len(variable[i])-1
            if string(variable[i][0])==string(variable_name) && string(variable[i][array_len])!=string("]"){
                variable_table[variable[i]]=variable_so(variable[variable_len])
                
            }else if string(variable[i][array_len])==string("]"){
            
                code_name_tmp:=strings.Split(variable[i],"[")
                code_name :=code_name_tmp[0]
                re := regexp.MustCompile(`\[(.*?)\]`)
                code_tmp_s := re.FindAllString(variable[i],-1)
                code_s:=code_tmp_s[0]
                code_s=strings.Replace(code_s,"[","",1)
                code_s=strings.Replace(ReverseString(code_s),"]","",1)
                code_s=ReverseString(code_s)
                
                
              re = regexp.MustCompile(`\$(.*?)\[`)
              code_tmp_s = re.FindAllString(variable[i],-1)
              code_tmp_s_r := strings.Replace(code_tmp_s[0],"[","",1)
              //fmt.Println(code_tmp_s_r)
              
              add_s := "yes"
              re = regexp.MustCompile(`\&(.*?)\;`)
              code_tmp_s = re.FindAllString(variable_table[code_name],-1)
              if len(code_tmp_s)!=0{
               for z:=0;z<=len(code_tmp_s)-1;z++{
                  if code_tmp_s[z] =="&"+variable_so(code_s)+";"{
                      add_s = "no"
                  }
               }
              }
                
                if string(add_s)==string("yes"){
                    variable_table[code_name] += code_tmp_s_r
                    variable_table[code_name] += "&"
                    variable_table[code_name] += variable_so(code_s)
                    variable_table[code_name] += ";"
                }
                
                if string(variable[i][len(variable[i])-1])==string("]"){
                variable[i]=code_name+"["+variable_so(code_s)+"]"
                
                }
                
                variable_table[variable[i]]=variable_so(variable[variable_len])
                
                var_array_s(variable[i])
                

            }
	    }
    }
}

func variable_so(name string)(string){    //变量指针
    //mt.Println(name)
    lens := len(name)-1
    if string(name[0])==string("$")&&string(name[lens])!=string("]"){
    name = strings.Replace(name, " ", "", -1 )
        return strings.Replace(variable_table[name], "\"", "", -1)
    }else if string(name[0])==string("$")&&string(name[lens])==string("]"){
        code_name_tmp:=strings.Split(name,"[")
        code_name :=code_name_tmp[0]
        //re := regexp.MustCompile(`\[(.*?)\]`)
        re := regexp.MustCompile(`\[(.*)\]`)
        code_tmp_s := re.FindAllString(name,-1)
        code_s:=code_tmp_s[0]
        code_s=strings.Replace(code_s,"[","",1)
        code_s=strings.Replace(ReverseString(code_s),"]","",1)
        code_s=ReverseString(code_s)
        
        return strings.Replace(variable_table[code_name+"["+variable_so(code_s)+"]"], "\"", "", -1)
       
    }else if string(name[0]) != string("\""){
        evals_code(name)
        return return_code_eval
    }else{
        return strings.Replace(name, "\"", "", -1)
    }
    
}

func ReverseString(str string) string {
	strArr := []rune(str)
	for i := 0; i < len(strArr)/2; i++ {
		strArr[len(strArr)-1-i], strArr[i] = strArr[i], strArr[len(strArr)-1-i]
	}
	return string(strArr)
}


func array_so(code string)(string){
    if string(code[len(code)-1]) == string(";"){
        re := regexp.MustCompile(`\&(.*?)\;`)
        code_tmp_s := re.FindAllString(code,-1)
     //fmt.Println(code_tmp_s)
    
        re = regexp.MustCompile(`\$(.*?)\&`)
        code_tmp_s_r := re.FindAllString(code,-1)
        code_for:=strings.Replace(string(code_tmp_s_r[0]),"&","",1)
        code_nn:=","
         result := "array("
        for i:=0;i<=len(code_tmp_s)-1;i++{
            code_for_num:=strings.Replace(string(code_tmp_s[i]),"&","",1)
            code_for_num=strings.Replace(string(code_for_num),";","",1)
            if i==len(code_tmp_s)-1{
                code_nn = ""
            }
            result += "["+code_for_num+"]"+"=>"+array_so(variable_so(code_for+"["+code_for_num+"]"))+code_nn
        }
        result += ")"
        return result
    }else{
        return code
    }
    
}

func evals_code(code string)(string){    //定义函数
    code_type := strings.Split(code,"(")
    if(string(code_type[0]) == string("print")){  //输出函数
       re := regexp.MustCompile(`(?s)\((.*)\)`)
       code_tmp := re.FindAllString(code,-1)
       code=strings.Replace(string(code_tmp[0]),"(","",1)
       code=strings.Replace(ReverseString(string(code)),")","",1)
       code=ReverseString(code)
       
       //fmt.Println(code)
       
       code_for_num:=""
       re = regexp.MustCompile(`\&(.*?)\;`)
       code_tmp_s := re.FindAllString(variable_so(code),-1)
       if len(code_tmp_s)==0{
           fmt.Println(variable_so(code))
       }else{
       fmt.Printf("array(")
       code_nn := ","
           for i:=0;i<=len(code_tmp_s)-1;i++{
           if i==len(code_tmp_s)-1{
                code_nn = ""
            }
               code_for_num=strings.Replace(string(code_tmp_s[i]),"&","",1)
               code_for_num=strings.Replace(string(code_for_num),";","",1)
               code_for_num="["+code_for_num+"]"
               arr_a:=variable_so(code+code_for_num)
               //fmt.Println(code,code_for_num)
               if len(arr_a)!=0{
                    if string(arr_a[len(arr_a)-1])==string(";"){
                        fmt.Printf(code_for_num+"=>"+array_so(arr_a)+code_nn)
                    }else{
                        fmt.Printf(code_for_num+"=>"+arr_a+code_nn)
                    }
                }
           }
       fmt.Println(")")
       
       }
       
      return_code_eval="true"
    }else if (string(code_type[0]) == string("base64Encode")){  //base64
       re := regexp.MustCompile(`(?s)\((.*)\)`)
       code_tmp := re.FindAllString(code,-1)
       code=strings.Replace(string(code_tmp[0]),"(","",1)
       code=strings.Replace(ReverseString(string(code)),")","",1)
       code=ReverseString(code)
       
       var msg =[]byte(variable_so(code))
	   encoding := base64.StdEncoding.EncodeToString(msg)
       return_code_eval=encoding
    }else if (string(code_type[0]) == string("base64Decode")){    //base64
    
       re := regexp.MustCompile(`(?s)\((.*)\)`)
       code_tmp := re.FindAllString(code,-1)
       code=strings.Replace(string(code_tmp[0]),"(","",1)
       code=strings.Replace(ReverseString(string(code)),")","",1)
       code=ReverseString(code)
       
       bytes, e := base64.StdEncoding.DecodeString(variable_so(code))
	   if e!=nil{
		   return_code_eval="FALSE"
       }
       return_code_eval=string(bytes)
    }else if (string(code_type[0]) == string("len")){   //len函数
       re := regexp.MustCompile(`(?s)\((.*)\)`)
       code_tmp := re.FindAllString(code,-1)
       code=strings.Replace(string(code_tmp[0]),"(","",1)
       code=strings.Replace(ReverseString(string(code)),")","",1)
       code=ReverseString(code)
       return_code_eval=strconv.Itoa(len(variable_so(code)))
    }else if(string(code_type[0]) == string("for ")){   //for循环
       //re := regexp.MustCompile(`(?s)\((.*)\):`)
       re := regexp.MustCompile(`([^)]+)`)
       code_tmp := re.FindAllString(code,-1)
       code_for:=strings.Replace(string(code_tmp[0]),"(","",1)
       code_for=strings.Replace(ReverseString(string(code_for)),")","",1)
       code_for=strings.Replace(string(code_for),":","",1)
       code_for=ReverseString(code_for)
             
       code_for_num:=strings.Replace(string(code_for),"[","",1)
       code_for_num=strings.Replace(string(code_for_num),"]","",1)
       //code_for_num =ReverseString(code_for_num)
       code_for_num_array := strings.Split(code_for_num,",")
       
       
       re = regexp.MustCompile(`(?s)\{(.*)\}`)
       code_tmp = re.FindAllString(code,-1)
       code=strings.Replace(string(code_tmp[0]),"{","",1)
       code=strings.Replace(ReverseString(string(code)),"}","",1)
       code=ReverseString(code)
       code_array := strings.Split(code,"#"+code_for_num_array[5])
       code_array_len:=len(code_array)-2
       
       
       if string(code_for_num_array[4])==string("+"){  //i++
       variable_table[code_for_num_array[3]]=code_for_num_array[0];
       i_1, _ := strconv.Atoi(code_for_num_array[0])
       i_2, _ := strconv.Atoi(code_for_num_array[1])
       i_3, _ := strconv.Atoi(code_for_num_array[2])
       variable_table[code_for_num_array[3]]=strconv.Itoa(0)
       
        for i := i_1; i <= i_2; i=i+i_3 {
           for z:=0;z<=code_array_len;z++{
               eval_code_whp(code_array[z])
           }
           variable_table_for_1,_:=strconv.Atoi(variable_table[code_for_num_array[3]])
           variable_table_for_2:=i_3
           vars := variable_table_for_1+variable_table_for_2
	 	   variable_table[code_for_num_array[3]]=strconv.Itoa(vars)
        }
       
       }else{   //i--
       i_1, _ := strconv.Atoi(code_for_num_array[0])
       i_2, _ := strconv.Atoi(code_for_num_array[1])
       i_3, _ := strconv.Atoi(code_for_num_array[2])
       variable_table[code_for_num_array[3]]=strconv.Itoa(0)
        for i := i_1; i >= i_2; i=i-i_3 {
        variable_table[code_for_num_array[3]]=strconv.Itoa(i)
           for z:=0;z<=code_array_len;z++{
               eval_code_whp(code_array[z])
           }
        }
       }
       delete(variable_table,code_for_num_array[3])
    }else if(string(code_type[0]) == string("if ")){    //判断
    
    
           res := regexp.MustCompile(`(\[)[^]]*(\])`)
           code_tmp_bb_s := res.FindAllString(code,-1)
           code_bb_s:=strings.Replace(string(code_tmp_bb_s[0]),"[-","",-1)
           code_bb_s=strings.Replace(ReverseString(string(code_bb_s)),"-]","",-1)
           code_bb_s=ReverseString(code_bb_s)
           
           code_bb_s=strings.Replace(ReverseString(string(code_bb_s)),"-","",-1)
           code_bb_s=strings.Replace(ReverseString(string(code_bb_s)),"]","",-1)
    
       
    
       code_array := strings.Split(code,"else#e"+code_bb_s)
       code_array_len:=len(code_array)-1
       
       
       
       for i:=0;i<=code_array_len;i++{
           //re := regexp.MustCompile(`(?s)\{(.*)\}`)
           re := regexp.MustCompile(`\{(.*)\}`)
           code_tmp := re.FindAllString(code_array[i],-1)
           
           //fmt.Println(code_tmp)
           
           code_code:=strings.Replace(string(code_tmp[0]),"{","",1)
           if(string(code_code[:2]) != string("if") && string(code_code[:3]) != string("fun")) && string(code_code[:3]) != string("for"){
               code_code=strings.Replace(ReverseString(string(code_code)),"}","",1)
               code_code=ReverseString(code_code)
           }
           
             
           re = regexp.MustCompile(`(\[)[^]]*(\])`)
           code_tmp_bb := re.FindAllString(code,-1)
           code_bb:=strings.Replace(string(code_tmp_bb[0]),"[-","",-1)
           code_bb=strings.Replace(ReverseString(string(code_bb)),"-]","",-1)
           code_bb=ReverseString(code_bb)
           
           code_bb=strings.Replace(ReverseString(string(code_bb)),"-","",-1)
           code_bb=strings.Replace(ReverseString(string(code_bb)),"]","",-1)
           //fmt.Println(code_bb)
           //fmt.Println(code)
           
           code_array_s := strings.Split(code_code,"#i"+code_bb)
           code_array_s_len:=len(code_array_s)-1
           
           //re = regexp.MustCompile(`(?s)\((.*)\)`)
           code_tmp_s := strings.Replace(code_array[i], " ", "", -1 )
           code_tmp_s = strings.Replace(code_tmp_s, "if", "", -1 )
           re = regexp.MustCompile(`([^)]+)`)
           code_tmp = re.FindAllString(code_tmp_s,-1)

           code_if:=strings.Replace(string(code_tmp[0]),"(","",1)
           code_if=strings.Replace(ReverseString(string(code_if)),")","",1)
           code_if=ReverseString(code_if)
           
           code_if_array := strings.Split(code_if,"=")
           code_if_0:=code_if_array[0]
           code_if_1:=""
           
           
           
           if (code_if_0=="true") {
               code_if_1=code_if_array[0]
           }else{
               code_if_1=code_if_array[1]
           }
           
           
           
           if code_if_0[len(code_if_0)-1:]==">" {
               if variable_so(code_if_0)>=variable_so(strings.Replace(code_if_1, " ", "", -1 )){              
                   for z:=0;z<=code_array_s_len;z++{
                       //variable_so(code_array_s[z])
                   }
                   break
                   
               }
           }else if code_if_0[len(code_if_0)-1:]=="<" {
               if variable_so(code_if_0)<=variable_so(strings.Replace(code_if_1, " ", "", -1 )){
                   for z:=0;z<=code_array_s_len;z++{
                       eval_code_whp(code_array_s[z])
                   }
                   break
                   
               }
           }else if code_if_0[len(code_if_0)-1:]=="!" {
               if variable_so(code_if_0)!=variable_so(strings.Replace(code_if_1, " ", "", -1 )){
                   for z:=0;z<=code_array_s_len;z++{
                       eval_code_whp(code_array_s[z])
                   }
                   break
                   
               }
           }else if code_if_0!="true"{
           
               if variable_so(code_if_0)==variable_so(strings.Replace(code_if_1, " ", "", -1 )){        
                   for z:=0;z<=code_array_s_len;z++{
                       eval_code_whp(code_array_s[z])
                   }
                   break
                   
               }
           }else if code_if_0=="true"{
                   for z:=0;z<=code_array_s_len;z++{
                       eval_code_whp(code_array_s[z])
                   }
                   break
           }else{
           
           }
       }
       
    }else if(string(code_type[0]) == string("function ")){   //function
       re := regexp.MustCompile(`([^)]+)`)
       code_tmp := re.FindAllString(code,-1)
       code_for:=strings.Replace(string(code_tmp[0]),"(","",1)
       code_for=strings.Replace(ReverseString(string(code_for)),")","",1)
       code_for=strings.Replace(string(code_for),":","",1)
       code_for=ReverseString(code_for)
       code_for=strings.Replace(code_for, " ", "", -1 )
       code_for=strings.Replace(code_for, "function", "", -1 )
       
       function_name_tmp := strings.Split(code_for,"[")
       function_name :=function_name_tmp[0]
       

       code_for=strings.Replace(string(code_for),function_name+"[","",1)
       code_for=strings.Replace(ReverseString(string(code_for)),"]","",1)
       code_for=strings.Replace(string(code_for),":","",1)
       code_for=ReverseString(code_for)
       code_for=strings.Replace(code_for, " ", "", -1 )
       code_for=strings.Replace(code_for, "function", "", -1 )
       
       function_var_tmp:=code_for
       //function_var := strings.Split(function_var_tmp,",")
       
       //re = regexp.MustCompile(`(?s)\{(.*)\}`)
       re = regexp.MustCompile(`\{(.*)\}`)
       code_tmp = re.FindAllString(code,-1)

       code_code:=strings.Replace(string(code_tmp[0]),"{","",1)
       code_code=strings.Replace(ReverseString(string(code_code)),"}","",1)
       code_code=ReverseString(code_code)

       function_code[function_name]=code_code
       function_variable[function_name]=function_var_tmp
       
       
       
    }else if(string(code_type[0]) == string("rand")){    //rand随机数
    
       re := regexp.MustCompile(`(?s)\((.*)\)`)
       code_tmp := re.FindAllString(code,-1)
       code=strings.Replace(string(code_tmp[0]),"(","",1)
       code=strings.Replace(ReverseString(string(code)),")","",1)
       code=ReverseString(code)
       
       
       code_array:=strings.Split(code,",")
       code_array_0,_ := strconv.Atoi(code_array[0])
       code_array_1,_ := strconv.Atoi(code_array[1])
       
    return_code_eval=strconv.Itoa(RandNum(code_array_0,code_array_1));
    
    
    }else if(string(code_type[0]) == string("time")){  //time时间戳
       
       return_code_eval=strconv.FormatInt((time.Now().UnixNano()),10)
       
    }else{
       if _, ok := function_code[string(code_type[0])]; ok {
           function_name := string(code_type[0])
	       variable_array := strings.Split(function_variable[function_name],",")
	       
           re := regexp.MustCompile(`(?s)\((.*)\)`)
           code_tmp := re.FindAllString(code,-1)
           code=strings.Replace(string(code_tmp[0]),"(","",1)
           code=strings.Replace(ReverseString(string(code)),")","",1)
           code=ReverseString(code)
       
           function_var := strings.Split(code,",")
	       

	       for i :=0;i<=len(variable_array)-1;i++{
	           variable_table[variable_array[i]]=variable_so(function_var[i])
	       }
	       
	       function_code_code:=strings.Split(function_code[function_name],"#F")
	       for i := 0; i<=len(function_code_code)-1;i++{
               eval_code_whp(function_code_code[i])
           }
           for i :=0;i<=len(variable_array)-1;i++{
	           delete(variable_table,variable_array[i])
	       }
           
       }else if code_type[0] != code && code[:2]!="//"{
           fmt.Println("运行异常,在运行",code,"时出错")
           os.Exit(3);
       }
        return_code_eval=code
    }
    return return_code_eval
}

func eval_code_whp(code string){   //代码运行
if(len(code)>2){
if(string(code[:2]) != string("if") && string(code[:3]) != string("fun")){
    variable := strings.Split(code,"=")
    variable_len:=len(variable)-1
    evals_code(variable[variable_len]);
    variable[variable_len]=return_code_eval
    variable_code(variable)
}else{
    evals_code(code);
}
}
}

func main(){   //结构代码
    code := GetFileContentAsStringLines(file)
    code_len := len(code)-1
    var code_new string
    for i := 0; i<=code_len;i++{
        code_new += code[i]
    }
    
    code = strings.Split(code_new,";")
    code_len=len(code)-1
    for i := 0; i<=code_len;i++{
        eval_code_whp(code[i])
    }
    //fmt.Println(variable_table)
    
}