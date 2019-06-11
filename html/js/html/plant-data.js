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
                $.each(JSON.parse(res["msg"]), function(i, item) {
                    $("#plantsList").append('<option value="{id}">{plant_id} - {name}</option>'
                        .replace('{id}', item["id"]).replace('{plant_id}', item["id"]).replace('{name}', item["name"]));

                });
                $('select').comboSelect();
            }
        } else {
            // backLogin();
        }
    });
}

getPlants();