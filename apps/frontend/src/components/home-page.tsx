import { useTranslation } from "react-i18next";
import { useDocumentTitle } from "usehooks-ts";
export function HomePage() {
  const { t } = useTranslation();
  useDocumentTitle(t("title.home"));

  return <div className="container">Home Page</div>;
}
