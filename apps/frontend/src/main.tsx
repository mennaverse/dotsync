import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { createRouter, RouterProvider } from "@tanstack/react-router";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { routeTree } from "./route-tree.gen";
import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faArrowRightFromBracket,
  faCodeBranch,
  faEye,
  faEyeSlash,
  faMoon,
  faSun,
  faUser,
} from "@fortawesome/free-solid-svg-icons";
import "./i18n";

import "virtual:uno.css";

library.add(
  faSun,
  faMoon,
  faEye,
  faEyeSlash,
  faCodeBranch,
  faUser,
  faArrowRightFromBracket,
);

const router = createRouter({ routeTree });

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const queryClient = new QueryClient();

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </StrictMode>,
);
