$(function () {
    submitRecordList = [{
        "SubmitRecordID": "100110011001100",
        "ProblemID": "1001",
        "SubmitStatus": "AC",
        "CreateTime": "2022-7-11 22:28:00",
    }, {
        "SubmitRecordID": "100110011001100",
        "ProblemID": "1001",
        "SubmitStatus": "WA",
        "CreateTime": "2022-7-11 22:28:00",
    }, {
        "SubmitRecordID": "100110011001100",
        "ProblemID": "1001",
        "SubmitStatus": "TLE",
        "CreateTime": "2022-7-11 22:28:00",
    }]


    for (var i = 0; i < submitRecordList.length; i++) {
        var line = $("<tr></tr>");

        var link = $("<a></a>",{
            href: "../html/showCode.html?SubmitRecordID"+submitRecordList[i]['SubmitRecordID'],
            text: submitRecordList[i]['SubmitRecordID'],
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
        }
        else {
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
})