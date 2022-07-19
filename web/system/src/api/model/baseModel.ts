export interface BasePageReq {
  page: number;
  pageSize: number;
}

export interface BaseListResp<T> {
  items: T[];
  total: number;
}

export interface BaseInfoResp<T> {
  code: number;
  msg: string;
  data: T;
}

export interface BaseResp {
  code: number;
  msg: string;
}
