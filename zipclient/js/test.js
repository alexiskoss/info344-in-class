function insertAjaxResult(address, handler) {
    $.ajax({
        url: address,
        success: handler
    });
}

function helloName() {
    $("#submit").click(showParams);
}

function helloNameHandler(text) {
    $("#nameForm").html(text);
};

function showParams() {
    var queryString = "name=" + escape($("#nameinput").val());
    $.ajax({ url: "http://localhost:4000/hello", data: queryString, success: helloNameHandler});
}

function getMemoryHandler(text) {
    $("#getMemory").html("Alloc: " + text.Alloc);
}
setInterval(function() { $.ajax({ url: "http://localhost:4000/memory", success: getMemoryHandler})}, 1000);