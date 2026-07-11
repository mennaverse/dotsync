import { MainShell } from "@/components/main-shell";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/profile/$username")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <MainShell>
      <ProfilePage />
    </MainShell>
  );
}

function ProfilePage() {
  const { username } = Route.useParams();

  return (
    <div className="container">
      <h1>Profile Page</h1>
      <p>This is the profile page of {username}.</p>
    </div>
  );
}
