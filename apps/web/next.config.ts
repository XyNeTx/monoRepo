import type { NextConfig } from "next";
import path from "path";
import { loadEnvFile } from "process";

const nextConfig: NextConfig = {
  /* config options here */
  reactCompiler: true,
  cacheComponents: true,
  // turbopack: {
  //   root: path.join(__dirname, '..'),
  // },
};

export default nextConfig;
