import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/a/$username/repository/$repository")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/a/$username/repository/$repository"!</div>;
}
