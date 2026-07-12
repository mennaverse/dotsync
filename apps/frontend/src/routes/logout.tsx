import { useLogoutMutation } from "@/hooks/use-logout-mutation";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { useEffect } from "react";
import z from "zod";

const searchSchema = z.object({
  returnUrl: z.string().optional(),
});

export const Route = createFileRoute("/logout")({
  component: RouteComponent,
  validateSearch: searchSchema,
});

function RouteComponent() {
  const { returnUrl } = Route.useSearch();
  const logoutMutation = useLogoutMutation();
  const navigate = useNavigate();

  useEffect(() => {
    logoutMutation.mutateAsync(undefined, {
      onSettled: async () => {
        await navigate({ to: returnUrl ?? "/" });
      },
    });
  }, []);

  return <></>;
}
