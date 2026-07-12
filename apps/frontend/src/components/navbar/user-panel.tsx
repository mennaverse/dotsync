import { Popover } from "@ark-ui/react/popover";
import { Portal } from "@ark-ui/react/portal";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useNavigate } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";
import { useState } from "react";

interface UserPanelProps {
  username: string;
}

export function UserPanel({ username }: UserPanelProps) {
  const [open, setOpen] = useState<boolean>(false);
  const { t } = useTranslation();
  const navigate = useNavigate();

  const handleNavigation = async ({
    to,
    params,
  }: {
    to: string;
    params?: Record<string, any>;
  }) => {
    setOpen(false);
    await navigate({ to, params });
  };

  const handleLogout = async () => {
    setOpen(false);
    await navigate({
      to: "/logout",
      search: { returnUrl: window.location.pathname },
    });
  };

  return (
    <Popover.Root
      open={open}
      onOpenChange={(e) => setOpen(e.open)}
      positioning={{
        placement: "bottom-end",
        offset: { mainAxis: 4, crossAxis: 12 },
      }}
    >
      <Popover.Trigger className="text-xl w-8 h-8 font-bold flex items-center justify-center cursor-pointer">
        {username?.[0].toUpperCase()}
      </Popover.Trigger>
      <Portal>
        <Popover.Positioner>
          <Popover.Content className="bg-gray-200 dark:bg-gray-800 border border-gray-500 w-48 p-2 mr-4 rounded-md shadow-md duration-200">
            <div className="flex flex-col gap-2 items-start text-gray-800 dark:text-gray-100 text">
              <button
                className="w-full bg-transparent hover:bg-gray-300 dark:hover:bg-gray-700 py-1 px-2 rounded text-left"
                onClick={() =>
                  handleNavigation({
                    to: "/profile/$username",
                    params: { username },
                  })
                }
              >
                <div className="flex gap-2 items-center">
                  <FontAwesomeIcon icon="user" />
                  <span>{t("navbar_panel_profile")}</span>
                </div>
              </button>
              <button
                className="w-full bg-transparent hover:bg-gray-300 dark:hover:bg-gray-700 py-1 px-2 rounded text-left"
                onClick={() =>
                  handleNavigation({
                    to: "/profile/$username/repositories",
                    params: { username },
                  })
                }
              >
                <div className="flex gap-2 items-center">
                  <FontAwesomeIcon icon="code-branch" />
                  <span>{t("navbar_panel_repositories")}</span>
                </div>
              </button>
              <hr className="w-full" />
              <button
                className="w-full bg-transparent hover:bg-gray-300 dark:hover:bg-gray-700 py-1 px-2 rounded text-left"
                onClick={handleLogout}
              >
                <div className="flex gap-2 items-center">
                  <FontAwesomeIcon icon="arrow-right-from-bracket" />
                  <span>{t("navbar_panel_logout")}</span>
                </div>
              </button>
            </div>
          </Popover.Content>
        </Popover.Positioner>
      </Portal>
    </Popover.Root>
  );
}
