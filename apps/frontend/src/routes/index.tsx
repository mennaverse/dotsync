import { createFileRoute } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";
import { useDocumentTitle } from "usehooks-ts";

export const Route = createFileRoute("/")({
  component: RouteComponent,
});

function RouteComponent() {
  const { t } = useTranslation();
  useDocumentTitle(t("index_doctitle"));

  return (
    <div className="container">
      <h1>Welcome to the Landing Page</h1>
      <p>This is the landing page of our application.</p>
    </div>
  );
}
