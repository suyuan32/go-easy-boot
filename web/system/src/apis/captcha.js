import service from "@/utils/request";

export const getCaptcha = (data) => {
    return service({
        url: '/captcha',
        method: 'get',
        data
    })
}