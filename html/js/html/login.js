window.localStorage.setItem("$SERVER", "http://127.0.0.1:8080/")

var $SERVER = window.localStorage.getItem("$SERVER");


$("#login").click(function() {
    var mobile = $("#username").val();
    var password = $("#password").val();

    $.ajax({
        type: 'POST',
        url: $SERVER + 'login',
        timeout: 30000,
        data: JSON.stringify({ "mobile": mobile, "pwd": password })
    }).then(function(res) {
        if (res) {
            if (res["status"] == 0) {
                window.localStorage.setItem("$TOKEN", res["msg"]["token"]);
                window.localStorage.setItem("$USERNAME", mobile)
                $(window).attr('location', 'index.html');
            } else {
                layui.use(['layer'], function() {
                    layer.open({
                        title: '登录异常',
                        content: '用户名或密码错误，请重试！'
                    });
                });
            }
        } else {
            layui.use(['layer'], function() {
                layer.open({
                    title: '登录异常',
                    content: '用户名或密码错误，请重试！'
                });
            });
        }
    });
});

$("#register").click(function() {
    layui.use(['layer'], function() {
        layer.open({
            title: '提示',
            content: '暂未开放注册功能，敬请期待！'
        });
    });
});