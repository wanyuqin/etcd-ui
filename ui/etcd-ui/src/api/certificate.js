import request from "@/utils/request"



export function createCertificates(data) {
    return request({
        url: "/v1/certificates",
        method: "post",
        data: data
    })
}

export function listCertificates(param) {
    return request({
        url: `/v1/certificates?page=${param.page}&page_size=${param.pageSize}&name=${param.name}`,
        method: "get"
    })

}

export function getCertificates(id) {
    return request({
        url: `/v1/certificates/${id}`,
        method: "get"
    })
}

export function delCertificates(id) {
    return request({
        url: `/v1/certificates/${id}`,
        method: "delete"
    })
}

export function updateCertificates(data) {
    return request({
        url: `/v1/certificates/${data.id}`,
        method: "put",
        data: data
    })
}