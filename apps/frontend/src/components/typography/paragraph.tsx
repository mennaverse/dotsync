import clsx from "clsx";
import type { HTMLAttributes } from "react";

interface ParagraphProps extends HTMLAttributes<HTMLParagraphElement> {}

export function Paragraph({ className, children, ...props }: ParagraphProps) {
  return (
    <p
      {...props}
      className={clsx(
        `text-gray-800 dark:text-gray-200 text-base leading-relaxed`,
        className,
      )}
    >
      {children}
    </p>
  );
}
