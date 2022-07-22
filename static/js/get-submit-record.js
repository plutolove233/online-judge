$(function () {
    let problemID = $.getUrlParam("ProblemID");
    $.ajax({
        headers: {
            "token": localStorage.getItem("token"),
        },
        url: 'http://localhost:8000/api_1_0/submit/list?ProblemID=' + problemID,
        type: 'GET',
        dataType: 'json',
        async: true,
        success: function (res) {
            console.log(res);
            if (res['code'] != "2000") {
                alert("获取提交列表失败，错误信息为：\n" + res['message']);
            } else {
                let submitRecordList = res['data'];
                for (var i = 0; i < submitRecordList.length; i++) {
                    var line = $("<tr></tr>");

                    var link = $("<a></a>", {
                        href: "../html/showCode.html?SubmitID=" + submitRecordList[i]['SubmitID'],
                        text: submitRecordList[i]['SubmitID'],
                    })

                    var row1 = $("<td></td>", {
                        class: "one",
                    });
                    row1.append(link);


                    var row2 = $("<td></td>", {
                        class: "two",
                        text: submitRecordList[i]['ProblemID'],
                    });


                    if (submitRecordList[i]['SubmitStatus'] === "AC") {
                        var row3 = $("<td></td>", {
                            class: "three-pass",
                            text: submitRecordList[i]["SubmitStatus"],
                        })
                    } else {
                        var row3 = $("<td></td>", {
                            class: "three-failed",
                            text: submitRecordList[i]["SubmitStatus"],
                        })
                    }


                    var row4 = $("<td></td>", {
                        class: "four",
                        text: submitRecordList[i]["CreateTime"],
                    })
                    // row3.append(link);

                    line.append(row1, row2, row3, row4);
                    $("#submit-record-list").append(line);
                }
            }
        },
        error: function (param) {
            console.log(param);
        },
    })
})