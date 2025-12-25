import { toast } from "sonner";

export async function toastPromise <T>
(   
    promiseFunc: Promise<T>,
    loadingTxt:string,
    successTxt:string,
    failureTxt:string,
)
{
    toast.promise((promiseFunc),{
        loading : loadingTxt,
        success : successTxt,
        error : failureTxt,
    })
}