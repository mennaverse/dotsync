import { MainShell } from "@/components/main-shell";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/search")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <MainShell>
      <SearchPage />
    </MainShell>
  );
}

function SearchPage() {
  return (
    <div className="container">
      <h1>Search Page</h1>
      <p>This is the search page of our application.</p>
    </div>
  );
}
