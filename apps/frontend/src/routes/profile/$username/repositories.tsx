import { createFileRoute } from "@tanstack/react-router"

export const Route = createFileRoute("/profile/$username/repositories")({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/profile/$username/repositories"!</div>
}
