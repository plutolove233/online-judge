$(function(){
    problem = [{
        "ProblemID": "1001",
        "ProblemTitle": "A+B",
        "ProblemStatus": true,
    },{
        "ProblemID": "1002",
        "ProblemTitle": "01序列",
        "ProblemStatus": false,
    },{
        "ProblemID": "1003",
        "ProblemTitle": "计算Sn",
        "ProblemStatus": true,
    }]

    
    for (var i = 0; i<problem.length; i++){
        var line = $("<tr></tr>");
        
        var row1 = $("<td></td>",{
            class: "one",
        });
        if (problem[i]['ProblemStatus']){
            var block = $("<div></div>")
            block.html('<svg viewBox="64 64 896 896" focusable="false" class="" data-icon="check" width="1em" height="1em" fill="green" aria-hidden="true"><path d="M912 190h-69.9c-9.8 0-19.1 4.5-25.1 12.2L404.7 724.5 207 474a32 32 0 0 0-25.1-12.2H112c-6.7 0-10.4 7.7-6.3 12.9l273.9 347c12.8 16.2 37.4 16.2 50.3 0l488.4-618.9c4.1-5.1.4-12.8-6.3-12.8z"></path></svg>');
            row1.append(block);
        }
        
        var row2 = $("<td></td>", {
            class: "two",
            text: problem[i]['ProblemID'],
        });

        var link = $("<a></a>",{
            href: "../html/submitCode.html?ProblemID="+problem[i]['ProblemID'],
            text: problem[i]['ProblemTitle'],
        })
        var row3 = $("<td></td>", {
            class:"three",
        })
        row3.append(link);

        line.append(row1,row2,row3);
        $("#problemList").append(line);
    }

    
})

