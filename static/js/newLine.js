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