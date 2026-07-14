import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/a/$username/profile/$profile")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/a/$username/profile/$profile"!</div>;
}
