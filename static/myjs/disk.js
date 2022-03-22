import request from "./request.js";

var params = {
    timestamp : Date.parse(new Date())
}

function FileList() {
    return request.post("user/doRegister", params)
}

const disk = {
    FileList : FileList,
};

export default disk;