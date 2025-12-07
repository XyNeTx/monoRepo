import { Login1 } from "@/components/login1";

export default function Page() {
    const logoProps = {
        url: "https://www.shadcnblocks.com",
        src: "https://deifkwefumgah.cloudfront.net/shadcnblocks/block/logos/shadcnblockscom-wordmark.svg",
        alt: "logo",
    }
    return (
        <div className="justify-center items-center bg-white dark:bg-black">
            <Login1 heading="Login" buttonText="Login" signupText="Create New Account" logo={logoProps} signupUrl=""/>
        </div>
    );
}