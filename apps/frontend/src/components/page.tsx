import clsx from "clsx/lite";
import type { PropsWithChildren } from "react";

interface PageProps extends PropsWithChildren {
  className?: string;
  contained?: boolean;
}

export function Page({ children, className, contained }: PageProps) {
  return (
    <div
      className={clsx(
        "size-full absolute inset-0",
        className,
        contained && "container mx-auto",
      )}
    >
      {children}
    </div>
  );
}
