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
    var strCode = $("#codeArea").val();
    strArr = strCode.split("\n");
    len = strArr.length;
    var flag = false;
    for (i = 0; i < len; i++)
        if (strArr[i].indexOf("int main()") != -1) {
            for (j = i; j < len; j++)
                if (strArr[j].indexOf("return 0") != -1) {
                    strArr.splice(j, 0, "\tgetchar();");
                    flag = true;
                    break;
                }
            if (flag) break;
        }
    if (!flag) {
        alert("没有在主函数内添加return 0");
    } else {
        console.log(strArr.join("\n"));
    }
}

$(function () {
    var proTitle = "A+B problem"
    var timeLimit = "1000";
    var memoryLimit = "256";
    var content = "输入A，B\n输出A+B";
    var inputLayout = "多行输出\n每组输入包含两个整数A,B，用一个空格分隔";
    var outputLayout = "多行输出\n每行输出一个整数，表示A+B";
    var inputExample = "5 8\n1 2";
    var outputExample = "13\n3";

    $("#problem-title").text(proTitle);
    var tle = $("<p></p>", {
        text: "Time limit:" + timeLimit + "ms",
    })
    var mle = $("<p></p>", {
        text: "Memory limit: " + memoryLimit + "Mb",
    })
    $("#problem-description").append(tle, mle);
    content = content.split("\n");
    for (var i = 0; i < content.length; i++) {
        var des = $("<p></p>", {
            text: content[i],
        })
        $("#problem-description").append(des);
    }
    inputLayout = inputLayout.split('\n');
    for (var i = 0; i<inputLayout.length; i++){
        var inputRequire = $("<p></p>", {
            text: inputLayout[i],
        })
        $("#input-description").append(inputRequire);
    }

    outputLayout = outputLayout.split('\n');
    for (var i = 0; i<outputLayout.length; i++){
        var outputRequire = $("<p></p>", {
            text: outputLayout[i],
        })
        $("#output-description").append(outputRequire);
    }

    inputExample = inputExample.split('\n');
    for (var i = 0; i<inputExample.length; i++){
        var inputRequire = $("<p></p>", {
            text: inputExample[i],
        })
        $("#input-example").append(inputRequire);
    }

    outputExample = outputExample.split('\n');
    for (var i = 0; i<outputExample.length; i++){
        var outputRequire = $("<p></p>", {
            text: outputExample[i],
        })
        $("#output-example").append(outputRequire);
    }
})