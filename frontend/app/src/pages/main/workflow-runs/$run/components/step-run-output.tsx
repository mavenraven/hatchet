import { CodeEditor } from '@/components/ui/code-editor';
import { Loading } from '@/components/ui/loading';

export interface StepRunOutputProps {
  output: string;
  isLoading: boolean;
  errors: string[];
}

export const StepRunOutput: React.FC<StepRunOutputProps> = ({
  output,
  isLoading,
  errors,
}) => {
  if (isLoading) {
    return <Loading />;
  }

  console.log('err', errors);
  return (
    <>
      <CodeEditor
        language="json"
        className="mb-4"
        height="400px"
        code={JSON.stringify(
          errors.length > 0
            ? errors.map((error) => error.split('\\n')).flat()
            : JSON.parse(output),
          null,
          2,
        )}
      />
    </>
  );
};
