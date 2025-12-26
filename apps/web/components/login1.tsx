'use client'
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { API_LINK } from "@/constants/constants";
import { toastPromise } from "@/lib/toastPromise";
import { IResponse } from "@/types/IResponse";
import axios from "axios";
import Image from "next/image";
import { FormEvent } from "react";
import { LoginDTO,SignupDTO } from "@/types/User";
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


//async function LoginClicked(e:React.FormEvent<HTMLFormElement>) {
async function LoginClicked(e:FormEvent<HTMLFormElement>) {
  e.preventDefault();
  //console.log({API_LINK});

  const formData = new FormData(e.currentTarget);
  //const formObj = Object.fromEntries(formData.entries())
  const data:LoginDTO = {
    email : formData.get("email") as string,
    password : formData.get("password") as string
  } 

  await toastPromise(Login(data),{
    loading:"Logging in Please Wait . . .",
    success:"Login Success Redirecting . . .",
    error:"Login Failed Please Try Again !!"
    }
  )

}

async function Login(data:LoginDTO) {
  await new Promise(resolve => setTimeout(resolve, 3000))
  try
  {
    await axios.post<IResponse<SignupDTO>>(API_LINK + "/api/users/login",
      data
    ).then((result)=> {
      document.cookie = `Authorization=LoggedIn;`;
      document.cookie = `name=${result.data.data?.name} ${result.data.data?.surname}`;
      //return console.log(result.data.data);
      return window.location.href = "/"
    })
  }
  catch (err){
    console.error(err);
    return {} as IResponse<SignupDTO>
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
          {/* <form onSubmit={LoginClicked}> */}
          <form onSubmit={(e)=>LoginClicked(e)}>
            <div className="min-w-sm border-muted bg-background flex w-full max-w-sm flex-col items-center gap-y-4 rounded-md border px-6 py-8 shadow-md">
              {heading && <h1 className="text-xl font-semibold">{heading}</h1>}
              <Input
                type="email"
                id="email"
                name="email"
                placeholder="Email"
                className="text-sm"
                required
              />
              <Input
                type="password"
                id="password"
                placeholder="Password"
                className="text-sm"
                name="password"
                required
              />
              <Button type="submit" className="w-full">
                {buttonText}
              </Button>
            <span className="text-end items-baseline text-xs align-end">Forgot Password ?</span>
            </div>
          </form>
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
