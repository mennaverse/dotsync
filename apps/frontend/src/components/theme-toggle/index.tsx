import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useEffect, useRef } from "react";
import { useLocalStorage } from "usehooks-ts";

type Theme = "light" | "dark";

export function ThemeToggle() {
  const [selectedTheme, setSelectedTheme] = useLocalStorage<Theme>(
    "theme",
    "light",
  );
  const bodyRef = useRef<HTMLElement>(document.body);

  useEffect(() => {
    handleThemeChange(selectedTheme);
  }, [selectedTheme]);

  const handleThemeChange = (theme: Theme) => {
    bodyRef.current.classList.remove("light", "dark");
    bodyRef.current.classList.add(theme);
  };

  return (
    <button
      className="bg-teal-800 rounded-full size-8 hover:bg-teal-800/50 transition-colors duration-200"
      onClick={() => {
        const theme = selectedTheme === "light" ? "dark" : "light";
        setSelectedTheme(theme);
        handleThemeChange(theme);
      }}
    >
      {selectedTheme === "light" ? (
        <FontAwesomeIcon icon="sun" />
      ) : (
        <FontAwesomeIcon icon="moon" />
      )}
    </button>
  );
}
