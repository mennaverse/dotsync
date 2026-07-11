import { createFileRoute } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";
import { useDocumentTitle } from "usehooks-ts";

export const Route = createFileRoute("/home")({
  component: RouteComponent,
});

function RouteComponent() {
  const { t } = useTranslation();
  useDocumentTitle(t("title.home"));

  return <div className="container">Home Page</div>;
}
