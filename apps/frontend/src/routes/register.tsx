import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/register")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div className="container">
      <h1>Register Page</h1>
      <p>This is the register page of our application.</p>
    </div>
  );
}
