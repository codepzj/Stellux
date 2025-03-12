// 请求分页携带的params
export interface Page{
    page_no:number,
    size:number,
    field?: string,
    order?: "ASC"|"DESC",
    keyword?: string
}

// 分页后返回的data的数据类型
export interface PageVO<T>{
    page_no:number,
    size:number,
    field:string,
    order: "ASC"|"DESC",
    total_count:number,
    total_page:number,
    list: T[]
}