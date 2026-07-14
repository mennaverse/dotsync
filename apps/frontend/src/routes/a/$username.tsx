import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/a/$username")({
  component: RouteComponent,
});

function RouteComponent() {
  const { username } = Route.useParams();

  return (
    <div className="container">
      <h1>Profile Page</h1>
      <p>This is the profile page of {username}.</p>
    </div>
  );
}
