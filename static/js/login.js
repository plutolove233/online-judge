function login() {
    let data = {};
    let value = $("#login-form").serializeArray();
    $.each(value, function (index, item) {
        data[item.name] = item.value;
    });
    let dataJson = JSON.stringify(data);

    $.ajax({
        url: "http://localhost:8000/api_1_0/login",
        type: "POST",
        async: true,
        contentType: "application/json",
        data: dataJson,
        success: function (res) {
            console.log(res);
            if (res['code'] != "2000") {
                alert("登录失败，错误信息为：\n" + res['message']);
            } else {
                localStorage.setItem('token', res.data.token);
                $(location).attr("href", "../html/problemList.html");
            }
        },
        error: function (param) {
            console.log(param);
        }
    })
}


function register() {
    let data = {};
    let value = $("#register-form").serializeArray();
    $.each(value, function (index, item) {
        data[item.name] = item.value;
    });
    data["IsAdmin"] = false;
    if (data['UserName'] === "" || data['Password'] === "" || data['RePassword'] === "") {
        alert("请填写完表格");
    } else {
        if (data['Password'] === data['RePassword']) {
            let dataJson = JSON.stringify(data);
            $.ajax({
                url: "http://localhost:8000/api_1_0/users/register",
                type: "POST",
                contentType: "application/json",
                data: dataJson,
                async: true,
                success: function (res) {
                    console.log(res);
                    if (res.code != "2000") {
                        alert("注册失败，错误信息为：\n" + res.message);
                    } else {
                        $(location).attr("href", "../html/login.html");
                    }
                },
                error: function (param) {
                    console.log(param);
                }
            })

        } else {
            alert("两次密码不相同");
        }
    }
}

function registerAdmin(){
    let data = {};
    let value = $("#register-form").serializeArray();
    $.each(value, function (index, item) {
        data[item.name] = item.value;
    });
    data["IsAdmin"] = true;
    if (data['UserName'] === "" || data['Password'] === "" || data['RePassword'] === "") {
        alert("请填写完表格");
    } else {
        if (data['Password'] === data['RePassword']) {
            let dataJson = JSON.stringify(data);
            $.ajax({
                url: "http://localhost:8000/api_1_0/users/register",
                type: "POST",
                contentType: "application/json",
                data: dataJson,
                async: true,
                success: function (res) {
                    console.log(res);
                    if (res.code != "2000") {
                        alert("注册失败，错误信息为：\n" + res.message);
                    } else {
                        $(location).attr("href", "../html/login.html");
                    }
                },
                error: function (param) {
                    console.log(param);
                }
            })

        } else {
            alert("两次密码不相同");
        }
    }
}