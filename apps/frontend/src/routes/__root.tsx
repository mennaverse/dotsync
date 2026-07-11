import { Outlet, createRootRoute } from "@tanstack/react-router";
import { TanStackDevtools } from "@tanstack/react-devtools";
import { formDevtoolsPlugin } from "@tanstack/react-form-devtools";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { MainShell } from "@/components/main-shell";

export const Route = createRootRoute({
  component: RootComponent,
});

function RootComponent() {
  return (
    <>
      <MainShell>
        <Outlet />
      </MainShell>

      <ReactQueryDevtools />
      <TanStackRouterDevtools />
      <TanStackDevtools
        config={{ hideUntilHover: true }}
        plugins={[formDevtoolsPlugin()]}
      />
    </>
  );
}
