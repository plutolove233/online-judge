function login() {
    let data = {};
    let value = $("#login-form").serializeArray();
    $.each(value, function (index, item) {
        data[item.name] = item.value;
    });
    let json = JSON.stringify(data);
    
    if (data["UserName"] === "123456" && data["Password"]==="456"){
        console.log(json);
        $(location).attr("href", "../html/problemList.html");
    } 
    else {
        alert("账号密码错误");
        console.log(data);
    }
}


function register(){
    let data = {};
    let value = $("#register-form").serializeArray();
    $.each(value, function (index, item) {
        data[item.name] = item.value;
    });
    if (data['UserName']==="" || data['Password']==="" || data['RePassword']===""){
        alert("请填写完表格");
    }
    else {
        if (data['Password'] === data['RePassword']){
            // do ajax
            console.log("register success");
            $(location).attr("href", "../html/login.html");
        }
        else {
            alert("两次密码不相同");
        }
    }
}