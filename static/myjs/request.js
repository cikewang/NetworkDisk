
const HTTP_REQUEST_URL = "http://pan.cikewang.com";
const TOKEN = localStorage.getItem("token")

function baseRequest(url, method, data)
{
    let Url = HTTP_REQUEST_URL;
    return new Promise((reslove, reject) => {
        $.ajax({
            url: Url + '/' + url,
            method: method || 'GET',
            beforeSend: function(request) {
                if (TOKEN) {
                    request.setRequestHeader("Authorization", TOKEN);
                }},
            data: data || {},
            success: (res) => {
                reslove(res)
            },
            fail: (msg) => {
                reject(msg)
            }
        })
    });
}

const request = {};

['options', 'get', 'post', 'put', 'head', 'delete', 'trace', 'connect'].forEach((method) => {
    request[method] = (api, data, opt) => baseRequest(api, method, data, opt || {})
});

export default request;