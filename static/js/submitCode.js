var num = "";

function line(n) {
    var lineobj = $("#leftNum");
    for (var i = 1; i <= n; i++) {

        num += i + "\n";
    }
    lineobj.val(num);
    num = "";
}

function keyUp() {
    var str = $("#codeArea").val();
    strNum = str.split("\n");
    n = strNum.length;
    line(n);
}

function submitCode() {
    let problemID = $.getUrlParam("ProblemID");
    var strCode = $("#codeArea").val();
    strCode = "#include <cstdio>\n" + strCode;
    strArr = strCode.split("\n");
    len = strArr.length;
    let flag = false;
    for (i = 0; i < len; i++) {
        if (strArr[i].indexOf("int main()") != -1) {
            let pos = strArr[i].indexOf("int main()");
            pos += 4;
            strArr[i] = strArr[i].slice(0, pos) + 'M' + strArr[i].slice(pos + 1);
            flag = true;
            break;
        }
    }

    let template = "\nint main(){\n\tMain();\n\tchar STOP_CONTROLLER;\n\twhile(STOP_CONTROLLER!='~'){\n\t\tSTOP_CONTROLLER=getchar();\n\t}\n}";

    if (flag) {
        let code = strArr.join("\n");
        code += template;
        console.log(code);
        let request = {
            "ProblemID": problemID,
            "CodeContext": code,
        };
        let dataJson = JSON.stringify(request);
        // ajax
        $.ajax({
            headers:{
                "token": localStorage.getItem("token"),
            },
            url: 'http://localhost:8000/api_1_0/submit/submit',
            type: 'POST',
            data: dataJson, //form-data or json
            contentType: 'application/json', //application/json or form-data
            async: true,
            success: function (res) {
                console.log(res);
                if (res.code != "2000") {
                    alert("代码提交失败，错误信息为：\n" + res['message']);
                } else {
                    let SubmitRecordID = res['data'];
                    $(location).attr("href", "../html/submit-result.html?SubmitID=" + SubmitRecordID);
                }
            },
            error: function (param) {
                console.log(param);
            },
        })
    } else {
        alert("没有main函数入口");
    }
}

function gotoSubmitRecord() {
    let problemID = $.getUrlParam("ProblemID");
    console.log(problemID);
    $(location).attr("href", "../html/submit-record.html?ProblemID=" + problemID);
}

$(function () {
    let problemID = $.getUrlParam("ProblemID");
    $.ajax({
        url: 'http://localhost:8000/api_1_0/problems/description?ProblemID=' + problemID,
        type: 'GET',
        dataType: 'json',
        async: true,
        success: function (res) {
            if (res.code != "2000") {
                console.log(res);
                alert("获取题目信息失败，错误信息为：\n" + res.message);
            } else {
                let info = res['data'];
                $("#problem-title").text(info.Title);
                var tle = $("<p></p>", {
                    text: "Time limit:" + info.TimeLimit + "ms",
                })
                var mle = $("<p></p>", {
                    text: "Memory limit: " + info.MemoryLimit + "Mb",
                })
                $("#problem-description").append(tle, mle);
                let content = info.Content;
                content = content.split("\n");
                for (var i = 0; i < content.length; i++) {
                    var des = $("<p></p>", {
                        text: content[i],
                    })
                    $("#problem-description").append(des);
                }
                let inputLayout = info.InputLayout;
                inputLayout = inputLayout.split('\n');
                for (var i = 0; i < inputLayout.length; i++) {
                    var inputRequire = $("<p></p>", {
                        text: inputLayout[i],
                    })
                    $("#input-description").append(inputRequire);
                }

                let outputLayout = info.OutputLayout;
                outputLayout = outputLayout.split('\n');
                for (var i = 0; i < outputLayout.length; i++) {
                    var outputRequire = $("<p></p>", {
                        text: outputLayout[i],
                    })
                    $("#output-description").append(outputRequire);
                }

                let inputExample = info.ExampleIn;
                inputExample = inputExample.split('\n');
                for (var i = 0; i < inputExample.length; i++) {
                    var inputRequire = $("<p></p>", {
                        text: inputExample[i],
                    })
                    $("#input-example").append(inputRequire);
                }

                let outputExample = info.ExampleOut;
                outputExample = outputExample.split('\n');
                for (var i = 0; i < outputExample.length; i++) {
                    var outputRequire = $("<p></p>", {
                        text: outputExample[i],
                    })
                    $("#output-example").append(outputRequire);
                }
            }
        },
        error: function (param) {
            console.log(param);
        },
    })


})