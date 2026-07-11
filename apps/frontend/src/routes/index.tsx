import { HomePage } from "@/components/home-page";
import { LandingPage } from "@/components/landing-page";
import { MainShell } from "@/components/main-shell";
import { useCurrentUser } from "@/hooks/use-current-user";
import { createFileRoute } from "@tanstack/react-router";
import { useMemo } from "react";

export const Route = createFileRoute("/")({
  component: RouteComponent,
});

function RouteComponent() {
  const { data: currentUser, isLoading } = useCurrentUser();
  const component = useMemo(
    () => (currentUser ? <HomePage /> : <LandingPage />),
    [currentUser],
  );

  if (isLoading) {
    return <MainShell></MainShell>;
  }

  return <MainShell>{component}</MainShell>;
}
