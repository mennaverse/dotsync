import { SearchBar } from "./search-bar";
import { NavLink } from "./nav-link";
import { Link } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";
import { useCurrentUser } from "@/hooks/use-current-user";
import { ThemeToggle } from "../theme-toggle";
import { UserPanel } from "./user-panel";

export function Navbar() {
  const { data: currentUser } = useCurrentUser();
  const { t } = useTranslation();

  return (
    <div className="w-full bg-teal-600 dark:bg-teal-600 text-white p-2 flex justify-center items-center">
      <nav className="container flex justify-between items-center">
        <Link
          className="text-xl hover:text-teal-200 transition-colors duration-200"
          to={currentUser ? "/home" : "/"}
        >
          {t("navbar_title")}
        </Link>
        <ul className="flex gap-4 items-center">
          <li className="w-full">
            <SearchBar />
          </li>
          <li>
            <ThemeToggle />
          </li>
          <NavLink to="/home">{t("navbar_home")}</NavLink>
          <NavLink to="/about">{t("navbar_about")}</NavLink>
          {!currentUser && (
            <>
              <Link
                to="/login"
                className="bg-teal-800 rounded-md px-2 py-1 hover:bg-teal-800/50 transition-colors duration-200 whitespace-nowrap"
              >
                <li>{t("navbar_login")}</li>
              </Link>
              <Link
                to="/register"
                className="bg-teal-500 rounded-md px-2 py-1 hover:bg-teal-500/50 transition-colors duration-200"
              >
                <li>{t("navbar_register")}</li>
              </Link>
            </>
          )}
          {currentUser && (
            <li className="rounded-full bg-teal-800 hover:bg-teal-800/50 transition-colors duration-200">
              <UserPanel username={currentUser?.username} />
            </li>
          )}
        </ul>
      </nav>
    </div>
  );
}
