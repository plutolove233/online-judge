$(function(){
    let submitID = $.getUrlParam("SubmitID");
    $.ajax({
        headers:{
            "token": localStorage.getItem("token"),
        },
        url: 'http://localhost:8000/api_1_0/submit/judge?SubmitID=' + submitID,
        type: 'GET',
        dataType: 'json',
        async: true,
        success: function(res){
            console.log(res);
            if (res.code != "2000"){
                alert("提交评测失败，错误信息为：\n" + res.message);
            }else{
                let SubmitStatus = res['data']['SubmitStatus'];
                let Message = res['data']['Message'];
                if (SubmitStatus==="AC"){
                    submit_res = $("<p></p>", {
                        text: SubmitStatus,
                        class: "ac",
                    })
                    Message = "答案正确";
                }
                else {
                    submit_res = $("<p></p>", {
                        text: SubmitStatus,
                        class: "err",
                    })
                }
                $("#res").append(submit_res);
                $("#message").append(Message);
            }
        },
        error: function(param){
            console.log(param);
        },
    })
})