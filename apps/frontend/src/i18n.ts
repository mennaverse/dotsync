import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import en from "./locales/en.json";
import ptBR from "./locales/pt-br.json";

declare module "i18next" {
  interface CustomTypeOptions {
    resources: {
      en: typeof en;
      "pt-BR": typeof ptBR;
    };
  }
}

import LanguageDetector from "i18next-browser-languagedetector";

i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    fallbackLng: "en",

    interpolation: {
      escapeValue: false, // not needed for react as it escapes by default
    },

    resources: {
      en,
      "pt-BR": ptBR,
    },
  });

export default i18n;
