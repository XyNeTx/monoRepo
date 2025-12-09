import type { NextConfig } from "next";
import { loadEnvFile } from "process";

loadEnvFile("../.env")

const nextConfig: NextConfig = {
  /* config options here */
  reactCompiler: true,
  cacheComponents: true,
  // turbopack:  {
  //   root: path.join(__dirname,'frontend'),
  // }
};

export default nextConfig;
