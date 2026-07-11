import { SearchBar } from "./search-bar";
import { NavLink } from "./nav-link";
import { Link } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";

export function Navbar() {
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
          <NavLink to="/">Home</NavLink>
          <NavLink to="/about">About</NavLink>
        </ul>
      </nav>
    </div>
  );
}
