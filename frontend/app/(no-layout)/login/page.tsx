import { Login1 } from "@/components/login1";
import axios from 'axios'
import { cacheLife } from "next/cache";

export interface GoResponse{
    message :string
}

export async function fetchHelloGo(){
    'use cache'
    cacheLife('max');
    try{
        const response:string = await axios.get<GoResponse>("http://localhost:8080/").then((res)=>{
            //console.log({res});
            return res.data.message;
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
    const data = await fetchHelloGo();
    const logoProps = {
        url: "https://www.shadcnblocks.com",
        src: "https://deifkwefumgah.cloudfront.net/shadcnblocks/block/logos/shadcnblockscom-wordmark.svg",
        alt: "logo",
    }
    return (
        <div className="justify-center items-center bg-white dark:bg-black">
            <Login1 heading={data} buttonText="Login" signupText="Create New Account" logo={logoProps} signupUrl=""/>
        </div>
    );
}