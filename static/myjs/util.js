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

/**
 * 图标
 * @param fileType
 */
function getIcon(fileType = '', category = 0) {
    var img = '/static/icon/';
    switch (category) {
        case 0:
            img += 'folder.png';
            break;
        case 1:
            img += 'image.png';
            break;
        case 2:
            if (fileType == 'doc' || fileType == 'docx') {
                img += 'word.png';
            } else if(fileType == 'ppt' || fileType == 'pptx') {
                img += 'ppt.png';
            } else {
                img += 'txt.png';
            }
            break;
        case 3:
            img += 'video.png';
            break;
        case 4:
            img += 'mp3.png';
            break;
        case 5:
            img += 'attachment.png';
            break;
        case 6:
            img += 'unknown.png';
            break;
        default:
            break;
    }

    return img;
}