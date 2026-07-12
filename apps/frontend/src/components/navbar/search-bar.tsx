import { searchBarFormSchema } from "@/schemas/search-bar";
import { Field } from "@ark-ui/react/field";
import { useForm } from "@tanstack/react-form";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";
import { useTranslation } from "react-i18next";

export function SearchBar() {
  const { t } = useTranslation();
  const form = useForm({
    defaultValues: {
      search: "",
    },
    validators: {
      onChange: searchBarFormSchema(t),
    },
    onSubmit: ({ value }) => {
      const { search } = value;
      console.log("Search submitted:", search);
    },
  });

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        e.stopPropagation();
        form.handleSubmit();
      }}
      className="flex items-center"
    >
      <form.Field
        name="search"
        children={(field) => (
          <Field.Root className="w-full max-w-md">
            <Field.Input
              name={field.name}
              onChange={(e) => field.handleChange(e.target.value)}
              onBlur={field.handleBlur}
              type="text"
              placeholder={t("search_placeholder")}
              className="w-full rounded-md bg-teal-700 border border-teal-300
                        px-2 py-1 focus:ring focus:ring-teal-200"
            />
          </Field.Root>
        )}
      />
      <button
        type="submit"
        className="ml-2 rounded-md bg-teal-500 px-2 py-1 text-white transition-colors duration-200
                hover:bg-teal-700 focus:outline-none focus:ring focus:ring-teal-200 cursor-pointer"
      >
        <FontAwesomeIcon icon={faSearch} className="size-4" />
      </button>
    </form>
  );
}
