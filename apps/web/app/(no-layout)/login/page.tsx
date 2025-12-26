import { Login1 } from "@/components/login1";
import { API_LINK} from "@/constants/constants";
import axios from 'axios'
import { cacheLife } from "next/cache";
import { cookies } from "next/headers";
import { redirect, RedirectType } from "next/navigation";

export interface GoResponse{
    message :string
}

async function CheckCookieAuth() {
    const allCookies = (await cookies()).getAll();
    const authToken = allCookies.find(x=>x.name.toLowerCase() === "authorization")
    if(authToken){
        redirect("/",RedirectType.push);
    }
}

export async function fetchHelloGo(){
    
    'use cache'
    cacheLife('max');
    //console.log(API_LINK);
    try{
        const response:string = await axios.get<GoResponse>(API_LINK!).then((res)=>{
            //console.log({res});
            return res.data.message + " Edited";
        })
        //console.log(response);
        return response;
    }
    catch (err) {
        console.error({err});
        return "";
    }
}

export default async function Page() {
    await CheckCookieAuth();
    const data = await fetchHelloGo();
    const logoProps = {
        url: "https://www.shadcnblocks.com",
        src: "https://deifkwefumgah.cloudfront.net/shadcnblocks/block/logos/shadcnblockscom-wordmark.svg",
        alt: "logo",
    }
    return (
        <div className="justify-center items-center bg-white dark:bg-black">
            <Login1 heading={data} buttonText="Login" signupText="Create New Account" logo={logoProps} signupUrl="/signup"/>
        </div>
    );
}