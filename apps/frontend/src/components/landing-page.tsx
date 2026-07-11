import { useTranslation } from "react-i18next";
import { useDocumentTitle } from "usehooks-ts";

export function LandingPage() {
  const { t } = useTranslation();
  useDocumentTitle(t("title"));

  return (
    <div className="container">
      <h1>Welcome to the Landing Page</h1>
      <p>This is the landing page of our application.</p>
    </div>
  );
}
