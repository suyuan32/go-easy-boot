import service from "@/utils/request";

export const login = (data) => {
    return service({
        url: '/api/user/login',
        method: 'post',
        data
    })
}

export const signup = (data) => {
    return service({
        url: '/api/user/register',
        method: 'post',
        data
    })
}

