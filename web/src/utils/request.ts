import { Response } from "@/types/response";
import qs from "qs";

type RequestMethod = "GET" | "POST" | "PUT" | "DELETE";

class Request {
  private readonly baseUrl: string;
  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  private async request<T, D>(
    url: string,
    method: RequestMethod,
    data?: T
  ): Promise<Response<D>> {
    // 请求配置项
    const options: RequestInit = {
      method,
      headers: {
        "Content-Type": "application/json",
      },
      body: data ? JSON.stringify(data) : undefined, // body携带参数
    };

    try {
      const res = await fetch(`${this.baseUrl}${url}`, options);
      const result = await res.json();
      if (!res.ok) {
        throw new Error(result.msg);
      }
      return result;
    } catch (err: any) {
      throw new Error(err.message);
    }
  }

  public get<D>(url: string, params?: object): Promise<Response<D>> {
    return this.request<any, D>(`${url}?${qs.stringify(params)}`, "GET");
  }

  public async post<T, D>(url: string, data: T): Promise<Response<D>> {
    return this.request<T, D>(url, "POST", data);
  }

  public async put<T, D>(url: string, data: T): Promise<Response<D>> {
    return this.request<T, D>(url, "PUT", data);
  }

  public async delete<T, D>(url: string, data: T): Promise<Response<D>> {
    const respData = this.request<T, D>(url, "DELETE", data);
    return respData;
  }
}

const baseUrl = process.env.API_BASE_URL;
if (!baseUrl) {
  throw new Error("baseUrl未设置，将读取默认配置");
}
const request = new Request(baseUrl);
export default request;
