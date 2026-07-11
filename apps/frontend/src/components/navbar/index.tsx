import { SearchBar } from "./search-bar";
import { NavLink } from "./nav-link";
import { Link } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";
import { useCurrentUser } from "@/hooks/use-current-user";

export function Navbar() {
  const { data: currentUser } = useCurrentUser();
  const { t } = useTranslation();

  return (
    <div className="w-full bg-teal-600 text-white p-2 flex justify-center items-center">
      <nav className="container flex justify-between items-center">
        <Link
          className="text-xl hover:text-teal-200 transition-colors duration-200"
          to="/"
        >
          {t("title")}
        </Link>
        <ul className="flex gap-4 items-center">
          <li className="w-full">
            <SearchBar />
          </li>
          <NavLink to="/home">{t("home")}</NavLink>
          <NavLink to="/about">{t("about")}</NavLink>
          {!currentUser && (
            <>
              <li className="bg-teal-800 rounded-md px-2 py-1 hover:bg-teal-800/50 transition-colors duration-200">
                <Link to="/login">{t("login")}</Link>
              </li>
              <li className="bg-teal-500 rounded-md px-2 py-1 hover:bg-teal-500/50 transition-colors duration-200">
                <Link to="/register">{t("register")}</Link>
              </li>
            </>
          )}
          {currentUser && (
            <li>
              <Link
                to={`/profile/$username`}
                params={{ username: currentUser.username }}
              >
                {currentUser.username}
              </Link>
            </li>
          )}
        </ul>
      </nav>
    </div>
  );
}
