import {
  FontAwesomeIcon,
  type FontAwesomeIconProps,
} from "@fortawesome/react-fontawesome";
import clsx from "clsx";
import type { ButtonHTMLAttributes } from "react";

interface UserPanelButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  icon?: FontAwesomeIconProps["icon"];
}

export function UserPanelButton({
  icon,
  children,
  className,
  ...rest
}: UserPanelButtonProps) {
  return (
    <button
      className={clsx(
        "w-full bg-transparent hover:bg-gray-300 dark:hover:bg-gray-700 py-1 px-2 rounded text-left",
        className,
      )}
      {...rest}
    >
      <div className="flex gap-2 items-center">
        {icon && <FontAwesomeIcon icon={icon} />}
        <span>{children}</span>
      </div>
    </button>
  );
}
