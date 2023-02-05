import request from "@/utils/request"


export function listKv() {
    return request({
        url: "/v1/keys",
        method: "get"
    })
}

export function deleteKey(name){
    return request ({
        url: `/v1/keys?name=${name}`,
        method: "delete"
    })
}

export function getKey(name){
    return request ({
        url: `/v1/keys/value?name=${name}`,
        method: "get"
    })
}


export function putKey(data) {
    return request ({
        url: "/v1/keys",
        method: "post",
        data : data
    })
}

export function watchKey(name){
    return request ({
        url: `/v1/keys/watch?name=${name}`,
        method: "post"
    })
}