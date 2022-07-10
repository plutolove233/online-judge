$(function () {
    pro_title = "A+B problem"
    who = "shyhao"
    submitStatus = "AC";
    code = "#include <cstdio>;\nusing namespace std;\nint main(){\n\treturn 0;\n}";
    var d1 = $("<div></div>")
    var t = $("<span></span>", {
        text: pro_title,
    })
    t.addClass("problemTitle");
    var prefix = $("<span></span>",{
        text: "题目：",
    })
    d1.append(prefix, t);
    
    var d2 = $("<div></div>");
    var author = $("<span></span>", {
        text: who,
    });
    var endpoint = $("<span></span>", {
        text: "提交的代码",
    });
    author.addClass("author");
    d2.append(author, endpoint);

    var d3 = $("<div></div>")
    var prefixStatus = $("<span></span>",{
        text: "提交状态:",
    })
    var statusC = $("<span></span>", {
        text: submitStatus,
    })
    statusC.addClass("author");
    d3.append(prefixStatus, statusC);
    $("#problem-header").append(d1, d2, d3);

    
    console.log(code);
    $("#code-content").text(code);
    hljs.initHighlightingOnLoad();
    hljs.initLineNumbersOnLoad();
    $('#code-content').css("cpp");

})