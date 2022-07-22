$(function () {
    // submitID = "20301";
    // pro_title = "A+B problem";
    // who = "shyhao";
    // submitStatus = "AC";
    // code = "#include <cstdio>;\nusing namespace std;\nint main(){\n\treturn 0;\n}";
    let submitID = $.getUrlParam("SubmitID");
    $.ajax({
        headers:{
            "token": localStorage.getItem("token"),
        },
        url: 'http://localhost:8000/api_1_0/submit/showCode?SubmitID=' + submitID,
        type: 'GET',
        dataType: 'json',
        async: true,
        success: function (res) {
            if (res.code != "2000") {
                alert("获取提交列表失败，错误信息为：\n" + res['message']);
            } else {
                var d = $("<div></div>", {});
                var preID = $("<span></span>", {
                    text: "提交ID：",
                });
                var sumit_id = $("<span></span>", {
                    text: res.data.SubmitID,
                });
                d.append(preID, sumit_id);

                var d1 = $("<div></div>")
                var t = $("<span></span>", {
                    text: res.data.Title,
                })
                t.addClass("problemTitle");
                var prefix = $("<span></span>", {
                    text: "题目：",
                })
                d1.append(prefix, t);

                var d2 = $("<div></div>");
                var author = $("<span></span>", {
                    text: res.data.UserID,
                });
                var endpoint = $("<span></span>", {
                    text: "提交的代码",
                });
                author.addClass("author");
                d2.append(author, endpoint);

                var d3 = $("<div></div>")
                var prefixStatus = $("<span></span>", {
                    text: "提交状态:",
                })
                var statusC = $("<span></span>", {
                    text: res.data.SubmitStatus,
                })
                statusC.addClass("author");
                d3.append(prefixStatus, statusC);
                $("#problem-header").append(d, d1, d2, d3);


                console.log(res.data.Code);
                $("#code-content").text(res.data.Code);
                hljs.initHighlightingOnLoad();
                hljs.initLineNumbersOnLoad();
                $('#code-content').css("cpp");
            }
        },
        error: function (param) {
            console.log(param);
        },
    })



})