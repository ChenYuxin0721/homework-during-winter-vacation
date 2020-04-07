# Status
1：成功  
0：失败
# Code
0：没有内容  
1：达到最后一页  
2：成功但不是最后一页  
# api/login
1. method:POST

2. Request Headers

    "Content-Type": "application/json"

3. Request Body


    {  
        "username": "abc",  
    
        "password": "1234"    
    }

4. Response Body

    //成功  
    {  
        "message": "",  
  
        "status": 1    
    }


    //失败  
    {  
        "message": "", 
  
        "status": 0,  
  
        "token": ""  
    }



