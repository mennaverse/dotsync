import type { AnyFieldApi } from "@tanstack/react-form";

export function ErrorText({ field }: { field: AnyFieldApi }) {
  if (!field.state.meta.isBlurred && !field.state.meta.isValid) {
    return <div className="h-1 text-xs"></div>;
  }

  return (
    <div className="h-1 text-red-500 text-xs flex flex-col gap-1">
      {field.state.meta.errors.map((error, index) => (
        <div key={index}>{error?.message}</div>
      ))}
    </div>
  );
}
