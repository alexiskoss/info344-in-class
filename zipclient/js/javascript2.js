function insertAjaxResult(address, handler) {
    $.ajax({
        url: address,
        success: handler
    });
}

$(function() {
    $("#submit").click(showParams);
});

function helloNameHandler(text) {
    $("#getName").html(text);
};

function showParams() {
    //var query = { name: $("#nameinput").val() }; // can use either or object or serialize v
    var query = $("#name-form").serialize();
    $.ajax({ url: "http://localhost:4000/hello", data: query, success: helloNameHandler});
}

function getMemoryHandler(text) {
    $("#getMemory").html("Alloc: " + text.Alloc);
}
setInterval(function() { $.ajax({ url: "http://localhost:4000/memory", success: getMemoryHandler})}, 1000);