import axios from "axios";
// qs is a library for converting json into x-www-form-urlencoded 
// import qs from 'qs'
const service = axios.create()
// adjust URL according to env settings
switch(process.env.NODE_ENV){
  case 'production':
    service.defaults.baseURL='http://api.xxx.com'
    break
  case 'dev':
    service.defaults.baseURL='http://127..0.0.1:8500'
    break
  default:
    service.defaults.baseURL='http://127..0.0.1:8500'
}
/**
 *  set timeout and cross domain credential
 */
service.defaults.timeout=10000          // ten seconds
service.defaults.withCredentials=true   // for example: session and cookie

service.defaults.headers['Content-Type']='application/json' 
// service.defaults.transformRequest=data=>qs.stringify(data)  //if content-type is x-www-form-urlencoded, use qs to convert
/**
 * set interceptors for requests
 * for example: add token from localStorage to the headers
 */
service.interceptors.request.use(config=>{
  let token=localStorage.getItem('token')
  token && (config.headers.Authoriztion=token)
  return config 
},error=>{
  return Promise.reject(error)
})
/**
 * set interceptors for responses
 */
service.interceptors.response.use(response=>{
    return response.data //return body data. 
},error=>{
  let { response } = error
  if(response){
    switch(response.status){
      case 404:
        // redirect
        break  
    }
  }else{
    // there is no response from server
    if(!window.navigator.onLine){
      return
    }
    return Promise.reject(error)
  }
})

export default service;