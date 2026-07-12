import { ErrorText } from "./error-text";
import { TextField } from "./text-field";
import { useFieldContext } from "@/hooks/form";

interface FormFieldProps {
  label?: string;
}

export function FormField({ label }: FormFieldProps) {
  const field = useFieldContext<string>();

  return (
    <div className="flex flex-col mb-4 gap-1">
      <TextField
        name={field.name}
        label={label}
        value={field.state.value}
        onChange={(e) => field.handleChange(e.target.value)}
        onBlur={field.handleBlur}
      />
      <ErrorText field={field} />
    </div>
  );
}
