var $SERVER = window.localStorage.getItem("$SERVER");
var $TOKEN = window.localStorage.getItem("$TOKEN");
var $USERNAME = window.localStorage.getItem("$USERNAME");

var getPlants = function() {
    $.ajax({
        type: 'GET',
        headers: {
            token: $TOKEN
        },
        url: $SERVER + 'torch/plants',
        timeout: 30000
    }).then(function(res) {
        if (res) {
            if (res["status"] == 0) {
                // console.log(res["msg"].length);
                // console.log(JSON.parse(res["msg"]));

                $.each(JSON.parse(res["msg"]), function(i, item) {
                    // console.log(item);
                    console.log(item["name"]);
                    console.log(item["id"]);
                });
            }
        } else {
            // backLogin();
        }
    });
}

getPlants();