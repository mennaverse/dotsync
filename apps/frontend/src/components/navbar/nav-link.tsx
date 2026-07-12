import { Link } from "@tanstack/react-router";
import type { ReactNode } from "react";

export function NavLink({
  to,
  className,
  children,
}: {
  to: string;
  className?: string;
  children: ReactNode;
}) {
  return (
    <Link to={to} className={className}>
      <li className="hover:text-teal-200 transition-colors duration-200">
        {children}
      </li>
    </Link>
  );
}
