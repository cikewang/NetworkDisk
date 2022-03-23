import request from "./request.js";

var params = {
    timestamp : Date.parse(new Date())
}

function FileList(category = 0, parentId = 0, path = '') {
    params['category'] = category;
    params['parentId'] = parentId;
    // params['path'] = path;
    return request.get("disk/list", params)
}

function AddDirectory(fileName, parentId = 0) {
    params['fileName'] = fileName;
    params['parentId'] = parentId;
    return request.post("disk/add", params)
}

const disk = {
    FileList : FileList,
    AddDirectory : AddDirectory
};

export default disk;