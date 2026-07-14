import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/a/$username/installer/$installer")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/a/$username/installer/$installer"!</div>;
}
