import type { HTMLInputAutoCompleteAttribute } from "react";
import { ErrorText } from "./error-text";
import { TextField } from "./text-field";
import { useFieldContext } from "@/hooks/form";

interface FormFieldProps {
  label?: string;
  type?: "text" | "email";
  autoComplete?: HTMLInputAutoCompleteAttribute;
}

export function FormField({
  label,
  type = "text",
  autoComplete = "off",
}: FormFieldProps) {
  const field = useFieldContext<string>();

  return (
    <div className="flex flex-col mb-4 gap-1">
      <TextField
        name={field.name}
        label={label}
        type={type}
        autoComplete={autoComplete}
        value={field.state.value}
        onChange={(e) => field.handleChange(e.target.value)}
        onBlur={field.handleBlur}
      />
      <ErrorText field={field} />
    </div>
  );
}
