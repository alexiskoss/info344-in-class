$(function() {
    $("#submit").click(showParams);
});

function cityNameHandler(text) {
    console.log(text);
    console.log(text.length);

    var states = [];
    $("#zipSelect").empty();
    $("#stateSelect").empty();
    for(i = 0; i < text.length; i++) {
        var zipElement = document.createElement("option");
        var zipNode = document.createTextNode(text[i].Code);
        zipElement.appendChild(zipNode);
        $("#zipSelect").append(zipElement);

        
        if(states.includes(text[i].State) == false) {
            var stateElement = document.createElement("option");
            var stateNode = document.createTextNode(text[i].State);
            stateElement.appendChild(stateNode);
            $("#stateSelect").append(stateElement);
            states.push(text[i].State);
            console.log(states);
        }
    }
};

function showParams(event) {
    event.preventDefault();
    //var query = { name: $("#nameinput").val() }; // can use either or object or serialize v
    var query = $("#citynameinput").val();
    console.log(query);
    $.ajax({ url: "/zips/" + query, success: cityNameHandler});
}