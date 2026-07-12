import { createFileRoute } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";
import { useDocumentTitle } from "usehooks-ts";

export const Route = createFileRoute("/search")({
  component: RouteComponent,
});

function RouteComponent() {
  const { t } = useTranslation();
  useDocumentTitle(t("search_doctitle"));

  return (
    <div className="container">
      <h1>Search Page</h1>
      <p>This is the search page of our application.</p>
    </div>
  );
}
