import type { PropsWithChildren } from "react";
import { Navbar } from "./navbar";

export function MainShell({ children }: PropsWithChildren) {
  return (
    <div className="flex flex-col items-center min-h-screen">
      <Navbar />
      <main className="flex-grow container">{children}</main>
    </div>
  );
}
