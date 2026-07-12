import type { PropsWithChildren } from "react";
import { Navbar } from "./navbar";

export function MainShell({ children }: PropsWithChildren) {
  return (
    <div className="flex flex-col min-h-screen size-full">
      <Navbar />
      <main className="flex-grow w-full relative">{children}</main>
    </div>
  );
}
