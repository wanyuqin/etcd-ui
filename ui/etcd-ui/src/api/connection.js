import request from "@/utils/request"


export function createConnection(data) {
    return request({
        url: "/v1/connections",
        method: "post",
        data: data
    })
}

export function listConnection(param) {
    return request({
        url: `/v1/connections?page=${param.page}&page_size=${param.pageSize}&name=${param.name}`,
        method: "get"
    })

}

export function getConnection(id) {
    return request({
        url: `/v1/connections/${id}`,
        method: "get"
    })
}

export function delConnection(id) {
    return request({
        url: `/v1/connections/${id}`,
        method: "delete"
    })
}

export function updateConnection(data) {
    return request({
        url: `/v1/connections/${data.id}`,
        method: "put",
        data: data
    })
}