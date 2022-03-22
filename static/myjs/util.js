/**
 * 弹框信息
 * @param message
 * @param type
 * @param localhost
 */
function showMessage(message, type = "danger", localhost = "") {
    var alert_type = "alert-"+type;
    $('#myModal').modal('show')
    $("#message_content").html(message);
    $("#message_content").addClass(alert_type);
    setTimeout(function(){
        $('#myModal').modal('hide')
        $("#message_content").removeClass(alert_type);
        if (localhost) {
            window.location.href = localhost;
        }
    },2000)
}