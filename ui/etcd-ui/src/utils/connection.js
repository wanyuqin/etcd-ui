import { connectionType } from "@/constant/connection.js";


export function transformConnectionType(key){
    return connectionType.find((item)=>{
        return item.key === key
    })?.value
}