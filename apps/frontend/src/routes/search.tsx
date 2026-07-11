import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/search")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div className="container">
      <h1>Search Page</h1>
      <p>This is the search page of our application.</p>
    </div>
  );
}
