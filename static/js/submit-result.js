$(function(){
    // let SubmitStatus = "AC";
    // let message = "通过";
    let SubmitStatus = "WA";
    let message = "答案错误";
    if (SubmitStatus==="AC"){
        submit_res = $("<p></p>", {
            text: SubmitStatus,
            class: "ac",
        })
    }
    else {
        submit_res = $("<p></p>", {
            text: SubmitStatus,
            class: "err",
        })
    }
    $("#res").append(submit_res);
    $("#message").append(message);
})