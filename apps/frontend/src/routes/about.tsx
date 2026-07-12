import { createFileRoute } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";
import { useDocumentTitle } from "usehooks-ts";

export const Route = createFileRoute("/about")({
  component: RouteComponent,
});

function RouteComponent() {
  return <AboutPage />;
}

function AboutPage() {
  const { t } = useTranslation();
  useDocumentTitle(t("about_doctitle"));

  return (
    <div className="container">
      <h1>About Page</h1>
      <p>This is the about page of our application.</p>
    </div>
  );
}
