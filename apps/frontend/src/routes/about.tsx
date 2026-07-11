import { MainShell } from "@/components/main-shell";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/about")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <MainShell>
      <AboutPage />
    </MainShell>
  );
}

function AboutPage() {
  return (
    <div className="container">
      <h1>About Page</h1>
      <p>This is the about page of our application.</p>
    </div>
  );
}
