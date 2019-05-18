var $SERVER = window.localStorage.getItem("$SERVER");
var $TOKEN = window.localStorage.getItem("$TOKEN");
var $USERNAME = window.localStorage.getItem("$USERNAME");

var backLogin = function() {
    window.localStorage.clear();
    $(window).attr('location', 'login.html');
}

var checkOut = function() {
    $.ajax({
        type: 'GET',
        headers: {
            token: $TOKEN
        },
        url: $SERVER + 'app/test',
        timeout: 30000
    }).then(function(res) {
        if (res) {
            if (res["status"] != 0) {
                backLogin();
            }
        } else {
            backLogin();
        }
    });
}


var checkUser = function() {
    if ($TOKEN == null) {
        backLogin();
    }

    checkOut();
}


checkUser();