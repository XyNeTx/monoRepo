'use client'
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { API_LINK } from "@/constants/constants";
import { IResponse } from "@/types/IResponse";
import axios from "axios";
import Image from "next/image";
import { useState } from "react";
interface Login1Props {
  heading?: string;
  logo: {
    url: string;
    src: string;
    alt: string;
    title?: string;
  };
  buttonText?: string;
  googleText?: string;
  signupText?: string;
  signupUrl?: string;
}

async function LoginClicked(email:string,password:string) {
  console.log({email});
  console.log({password});
  console.log({API_LINK})
  console.log(process.env.NEXT_PUBLIC_API_LINK);
  try{
    const response:IResponse= await axios.post<IResponse>(API_LINK + "/users/login",{
      Email: email,
      PasswordHash: password
    }).then((result)=>{
      return result.data;
    })
    console.log(response);
  }
  catch (err){
    console.error(err);
  }
}

const Login1 = ({
  heading = "Login",
  logo = {
    url: "#",
    src: "https://deifkwefumgah.cloudfront.net/shadcnblocks/block/logos/shadcnblockscom-wordmark.svg",
    alt: "logo",
    title: "Application Logo",
  },
  buttonText = "Login",
  signupText = "Need an account?",
  signupUrl = "https://shadcnblocks.com",
}: Login1Props) => {
  const [email,setEmail] = useState<string>("");
  const [password,setPassword] = useState<string>("");
  return (
    <section className="bg-muted h-screen">
      <div className="flex h-full bg-gray-100 dark:bg-black items-center justify-center">
        {/* Logo */}
        <div className="flex flex-col items-center gap-6 lg:justify-start">
          <a href={logo.url}>
            <Image
              src={logo.src}
              alt={logo.alt}
              title={logo.title}
              className="h-10 dark:invert"
              width={200}
              height={40}
            />
          </a>
          <div className="min-w-sm border-muted bg-background flex w-full max-w-sm flex-col items-center gap-y-4 rounded-md border px-6 py-8 shadow-md">
            {heading && <h1 className="text-xl font-semibold">{heading}</h1>}
            <Input
              type="email"
              placeholder="Email"
              className="text-sm"
              required
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            <Input
              type="password"
              placeholder="Password"
              className="text-sm"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <Button type="button" onClick={() => LoginClicked(email,password)} className="w-full">
              {buttonText}
            </Button>
          </div>
          <div className="text-muted-foreground flex justify-center gap-1 text-sm">
            <p>{signupText}</p>
            <a
              href={signupUrl}
              className="text-primary font-medium hover:underline"
            >
              Sign up
            </a>
          </div>
        </div>
      </div>
    </section>
  );
};

export { Login1 };
