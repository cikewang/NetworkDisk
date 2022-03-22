import request from "/static/myjs/request.js";

var params = {
    timestamp : Date.parse(new Date())
}

/**
 *  用户注册
 * @param username
 * @param password
 * @param password_verify
 * @returns {*}
 * @constructor
 */
 function UserRegister(username, password, password_verify) {
    params['username'] = username;
    params['password'] = password;
    params['password_verify'] = password_verify;
    return request.post("user/doRegister", params, {noAuth:true})
}

/**
 * 用户登录
 * @param username
 * @param password
 * @returns {*}
 * @constructor
 */
function UserLogin(username, password) {
    params['username'] = username;
    params['password'] = password;
    return request.post("user/doLogin", params, {noAuth:true})
}

// 用户退出
function UserLogout() {
    return request.get("user/logout", params, {noAuth:true})
}

// 用户信息
function UserInfo() {
    return request.get("user/info", params,  {noAuth:true})
}

const user = {
    UserRegister : UserRegister,
    UserLogin : UserLogin,
    UserLogout : UserLogout,
    UserInfo : UserInfo,
};

export default user;