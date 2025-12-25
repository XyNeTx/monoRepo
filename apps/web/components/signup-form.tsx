'use client'
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { API_LINK } from "@/constants/constants"
import { IResponse } from "@/types/IResponse"
import axios from "axios"
import { toast } from "sonner"
import { redirect } from "next/navigation"

async function SignupClick(e:React.FormEvent<HTMLFormElement>) {
  // console.log({email});
  // console.log({password});
  e.preventDefault();
  console.log({e})
  const formData = new FormData(e.currentTarget);
  console.log({formData})
  console.log({API_LINK});
  const data = Object.fromEntries(formData);
  console.log({data});
  if(data.confirmpassword !== data.password){
    return toast.error("Password and Confirm Password not matched")
  }
  try{
    const response:IResponse = await axios.post<IResponse>(API_LINK + "/api/users/signup", data
    ).then((result)=> {
      toast.success("Registration Success Redirecting to Login. . .")
      return result.data;
    })
    console.log(response);
    redirect("/login");
  }
  catch (err){
    console.error(err);
    toast.error("Sign Up not Success Please Try Again");
    return {} as IResponse
  }
}

export function SignupForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader className="text-center">
          <CardTitle className="text-xl">Create your account</CardTitle>
          <CardDescription>
            Enter your email below to create your account
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form id="formSignUp" onSubmit={SignupClick}>
            <FieldGroup>
              <Field className="grid grid-cols-2 gap-4">
                  <Field>
                    <FieldLabel htmlFor="firstname">First Name</FieldLabel>
                    <Input id="name" name="name" type="text" required />
                  </Field>
                  <Field>
                    <FieldLabel htmlFor="lastname">Last Name</FieldLabel>
                    <Input id="surname" name="surname" type="text" required />
                  </Field>
                </Field>
              <Field>
                <FieldLabel htmlFor="age">Age</FieldLabel>
                <Input
                  id="age"
                  type="number"
                  name="age"
                  min={"18"}
                  max={"50"}
                  placeholder="18-50"
                  step={"1"}
                  required
                />
              </Field>
              <Field>
                <FieldLabel htmlFor="email">Email</FieldLabel>
                <Input
                  id="email"
                  type="email"
                  name="email"
                  placeholder="m@example.com"
                  required
                />
              </Field>
              <Field>
                <Field className="grid grid-cols-2 gap-4">
                  <Field>
                    <FieldLabel htmlFor="password">Password</FieldLabel>
                    <Input id="password" name="password" type="password" required />
                  </Field>
                  <Field>
                    <FieldLabel htmlFor="confirm-password">
                      Confirm Password
                    </FieldLabel>
                    <Input id="confirm-password" name="confirmpassword" type="password" required />
                  </Field>
                </Field>
                <FieldDescription>
                  Must be at least 8 characters long.
                </FieldDescription>
              </Field>
              <Field>
                <Button type="submit">Create Account</Button>
                <FieldDescription className="text-center">
                  Already have an account? <a href="#">Sign in</a>
                </FieldDescription>
              </Field>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
      <FieldDescription className="px-6 text-center">
        By clicking continue, you agree to our <a href="#">Terms of Service</a>{" "}
        and <a href="#">Privacy Policy</a>.
      </FieldDescription>
    </div>
  )
}
