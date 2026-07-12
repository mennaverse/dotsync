import clsx from "clsx/lite";
import type { PropsWithChildren } from "react";

interface HeadingProps extends PropsWithChildren {
  level?: 1 | 2 | 3 | 4 | 5 | 6;
  className?: string;
}

export function Heading({ level = 1, className, children }: HeadingProps) {
  if (level === 1) {
    return (
      <h1
        className={clsx(
          "text-4xl font-bold mb-4 text-teal-800 dark:text-teal-400",
          className,
        )}
      >
        {children}
      </h1>
    );
  }
  if (level === 2) {
    return (
      <h2
        className={clsx(
          "text-3xl font-bold mb-4 text-teal-800 dark:text-teal-400",
          className,
        )}
      >
        {children}
      </h2>
    );
  }
  if (level === 3) {
    return (
      <h3
        className={clsx(
          "text-2xl font-bold mb-4 text-teal-800 dark:text-teal-400",
          className,
        )}
      >
        {children}
      </h3>
    );
  }
  if (level === 4) {
    return (
      <h4
        className={clsx(
          "text-xl font-bold mb-4 text-teal-800 dark:text-teal-400",
          className,
        )}
      >
        {children}
      </h4>
    );
  }
  if (level === 5) {
    return (
      <h5
        className={clsx(
          "text-lg font-bold mb-4 text-teal-800 dark:text-teal-400",
          className,
        )}
      >
        {children}
      </h5>
    );
  }
  if (level === 6) {
    return (
      <h6
        className={clsx(
          "text-base font-bold mb-4 text-teal-800 dark:text-teal-400",
          className,
        )}
      >
        {children}
      </h6>
    );
  }
  return null;
}
