import { Response } from "@/types/response";
import qs from "qs";

type RequestMethod = "GET" | "POST" | "PUT" | "DELETE";

class Request {
  private baseUrl: string;
  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  public async request<T, D>(
    url: string,
    method: RequestMethod,
    data?: T
  ): Promise<Response<D>> {
    const options: RequestInit = {
      method,
      headers: {
        "Content-Type": "application/json",
      },
    };

    if (data) {
      options.body = JSON.stringify(data);
    }

    console.log(`${this.baseUrl}${url}`);
    const res = await fetch(`${this.baseUrl}${url}`, options);

    if (!res.ok) {
      throw new Error(`网络错误: ${res.status}`);
    }

    return res.json();
  }

  public async get<D>(url: string, params?: object): Promise<Response<D>> {
    const respData = this.request<unknown, D>(
      `${url}?${qs.stringify(params)}`,
      "GET"
    );
    return respData;
  }

  public async post<T, D>(url: string, data: T): Promise<Response<D>> {
    const respData = this.request<T, D>(url, "POST", data);
    return respData;
  }

  public async put<T, D>(url: string, data: T): Promise<Response<D>> {
    const respData = this.request<T, D>(url, "PUT", data);
    return respData;
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
