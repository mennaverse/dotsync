import { Popover } from "@ark-ui/react/popover";
import { Portal } from "@ark-ui/react/portal";
import { useNavigate } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";
import { useState } from "react";
import { UserPanelButton } from "./user-panel-button";

interface UserPanelProps {
  username: string;
}

interface NavigationParams {
  to: string;
  params?: Record<string, any>;
}

export function UserPanel({ username }: UserPanelProps) {
  const [open, setOpen] = useState<boolean>(false);
  const { t } = useTranslation();
  const navigate = useNavigate();

  const handleNavigation = async ({ to, params }: NavigationParams) => {
    setOpen(false);
    await navigate({ to, params });
  };

  const handleProfileNavigation = async ({ to, params }: NavigationParams) => {
    await handleNavigation({ to, params: { username, ...params } });
  };

  const handleLogout = async () => {
    await handleNavigation({
      to: "/logout",
      params: { returnUrl: window.location.pathname },
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
          <Popover.Content className="bg-gray-200 dark:bg-gray-800 border border-gray-500 min-w-48 p-2 mr-4 rounded-md shadow-md duration-200">
            <div className="flex flex-col gap-2 items-start text-gray-800 dark:text-gray-100 text">
              <UserPanelButton
                icon="user"
                onClick={() =>
                  handleProfileNavigation({
                    to: "/a/$username",
                  })
                }
              >
                {t("navbar_panel_my_account")}
              </UserPanelButton>
              <UserPanelButton
                icon="code-branch"
                onClick={() =>
                  handleProfileNavigation({
                    to: "/a/$username/repository",
                  })
                }
              >
                {t("navbar_panel_repositories")}
              </UserPanelButton>
              <UserPanelButton
                icon="file-download"
                onClick={() =>
                  handleProfileNavigation({
                    to: "/a/$username/installer",
                  })
                }
              >
                {t("navbar_panel_installer")}
              </UserPanelButton>
              <hr className="w-full" />
              <UserPanelButton
                icon="arrow-right-from-bracket"
                onClick={handleLogout}
              >
                {t("navbar_panel_logout")}
              </UserPanelButton>
            </div>
          </Popover.Content>
        </Popover.Positioner>
      </Portal>
    </Popover.Root>
  );
}
