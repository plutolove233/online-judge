function newProblem() {
    numberField = ["TimeLimit", "MemoryLimit"];
    let data = {};
    let value = $("#problemDescriptionForm").serializeArray();
    $.each(value, function (index, item) {
        if ($.inArray(item.name, numberField) != -1) {
            data[item.name] = Number(item.value);
        } else {
            data[item.name] = item.value;
        }
    });
    console.log(JSON.stringify(data));
    console.log(data);
}