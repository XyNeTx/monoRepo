import { toast } from "sonner";

export function toastPromise<T>
(   
    promise: Promise<T>,
    messages: {
        loading:string,
        success:string,
        error:string,
    }
): Promise<T> {
    return toast.promise(promise,messages).unwrap()
}