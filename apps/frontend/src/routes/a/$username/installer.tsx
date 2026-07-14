import { Page } from "@/components/page";
import { Heading } from "@/components/typography/heading";
import { createFileRoute } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";
import { useDocumentTitle } from "usehooks-ts";

export const Route = createFileRoute("/a/$username/installer")({
  component: RouteComponent,
});

function RouteComponent() {
  const { t } = useTranslation();

  useDocumentTitle(t("account_installer_doctitle"));

  return (
    <Page contained>
      <Heading level={1}>{t("account_installer_heading")}</Heading>
    </Page>
  );
}
