import service from "@/utils/request";

export const getCaptcha = (data) => {
    return service({
        url: '/api/captcha',
        method: 'get',
        data
    })
}