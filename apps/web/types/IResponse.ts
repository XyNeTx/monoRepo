export interface IResponse<T>{
    data? : Partial<T>,
    message : string
}