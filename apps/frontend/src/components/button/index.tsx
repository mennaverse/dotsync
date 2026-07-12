import clsx from "clsx/lite";
import type { ButtonHTMLAttributes } from "react";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {}

export function Button({ className, children, ...props }: ButtonProps) {
  return (
    <button
      {...props}
      className={clsx(
        "bg-teal-400 rounded-lg px-4 py-2 hover:bg-teal-400/50 cursor-pointer",
        "transition-colors duration-200 disabled:bg-gray-400 disabled:cursor-not-allowed",
        className,
      )}
    >
      {children}
    </button>
  );
}
