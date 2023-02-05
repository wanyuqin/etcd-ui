import request from "@/utils/request"


export function clusterStatus() {
    return request({
        url: "/v1/cluster:status",
        method: "get"
    })
}
