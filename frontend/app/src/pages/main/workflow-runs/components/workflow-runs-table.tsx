import { DataTable } from '@/components/molecules/data-table/data-table.tsx';
import { columns } from './workflow-runs-columns';
import { useMemo, useState } from 'react';
import {
  ColumnFiltersState,
  PaginationState,
  SortingState,
  VisibilityState,
} from '@tanstack/react-table';
import { useQuery } from '@tanstack/react-query';
import invariant from 'tiny-invariant';
import { WorkflowRunStatus, queries } from '@/lib/api';
import { Loading } from '@/components/ui/loading.tsx';
import { TenantContextType } from '@/lib/outlet';
import { useOutletContext } from 'react-router-dom';
import {
  FilterOption,
  ToolbarFilters,
  ToolbarType,
} from '@/components/molecules/data-table/data-table-toolbar';
import { Button } from '@/components/ui/button';
import { ArrowPathIcon } from '@heroicons/react/24/outline';
import { WorkflowRunsMetricsView } from './workflow-runs-metrics';

export interface WorkflowRunsTableProps {
  workflowId?: string;
  parentWorkflowRunId?: string;
  parentStepRunId?: string;
  initColumnVisibility?: VisibilityState;
  filterVisibility?: { [key: string]: boolean };
  refetchInterval?: number;
}

export function WorkflowRunsTable({
  workflowId,
  initColumnVisibility = {},
  filterVisibility = {},
  parentWorkflowRunId,
  parentStepRunId,
  refetchInterval = 5000,
}: WorkflowRunsTableProps) {
  const { tenant } = useOutletContext<TenantContextType>();
  invariant(tenant);

  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] =
    useState<VisibilityState>(initColumnVisibility);

  const [pagination, setPagination] = useState<PaginationState>({
    pageIndex: 0,
    pageSize: 50,
  });
  const [pageSize, setPageSize] = useState<number>(50);

  const offset = useMemo(() => {
    if (!pagination) {
      return;
    }

    return pagination.pageIndex * pagination.pageSize;
  }, [pagination]);

  const workflow = useMemo<string | undefined>(() => {
    if (workflowId) {
      return workflowId;
    }

    const filter = columnFilters.find((filter) => filter.id === 'Workflow');

    if (!filter) {
      return;
    }

    const vals = filter?.value as Array<string>;
    return vals[0];
  }, [columnFilters, workflowId]);

  const statuses = useMemo(() => {
    const filter = columnFilters.find((filter) => filter.id === 'status');

    if (!filter) {
      return;
    }

    return filter?.value as Array<WorkflowRunStatus>;
  }, [columnFilters]);

  const listWorkflowRunsQuery = useQuery({
    ...queries.workflowRuns.list(tenant.metadata.id, {
      offset,
      limit: pageSize,
      statuses,
      workflowId: workflow,
      parentWorkflowRunId,
      parentStepRunId,
    }),
    refetchInterval,
  });

  const metricsQuery = useQuery({
    ...queries.workflowRuns.metrics(tenant.metadata.id, {
      workflowId: workflow,
      parentWorkflowRunId,
      parentStepRunId,
    }),
    refetchInterval,
  });

  const {
    data: workflowKeys,
    isLoading: workflowKeysIsLoading,
    error: workflowKeysError,
  } = useQuery({
    ...queries.workflows.list(tenant.metadata.id),
  });

  const workflowKeyFilters = useMemo((): FilterOption[] => {
    return (
      workflowKeys?.rows?.map((key) => ({
        value: key.metadata.id,
        label: key.name,
      })) || []
    );
  }, [workflowKeys]);

  const workflowRunStatusFilters = useMemo((): FilterOption[] => {
    return [
      {
        value: WorkflowRunStatus.SUCCEEDED,
        label: 'Succeeded',
      },
      {
        value: WorkflowRunStatus.FAILED,
        label: 'Failed',
      },
      {
        value: WorkflowRunStatus.RUNNING,
        label: 'Running',
      },
      {
        value: WorkflowRunStatus.QUEUED,
        label: 'Queued',
      },
      {
        value: WorkflowRunStatus.PENDING,
        label: 'Pending',
      },
    ];
  }, []);

  const filters: ToolbarFilters = [
    {
      columnId: 'Workflow',
      title: 'Workflow',
      options: workflowKeyFilters,
      type: ToolbarType.Radio,
    },
    {
      columnId: 'status',
      title: 'Status',
      options: workflowRunStatusFilters,
    },
  ].filter((filter) => filterVisibility[filter.columnId] != false);

  const [rotate, setRotate] = useState(false);

  const refetch = () => {
    listWorkflowRunsQuery.refetch();
    metricsQuery.refetch();
  };

  const actions = [
    <Button
      key="refresh"
      className="h-8 px-2 lg:px-3"
      size="sm"
      onClick={() => {
        refetch();
        setRotate(!rotate);
      }}
      variant={'outline'}
      aria-label="Refresh events list"
    >
      <ArrowPathIcon
        className={`h-4 w-4 transition-transform ${rotate ? 'rotate-180' : ''}`}
      />
    </Button>,
  ];

  if (listWorkflowRunsQuery.isLoading) {
    return <Loading />;
  }

  return (
    <>
      {metricsQuery.data && (
        <div className="mb-4">
          <WorkflowRunsMetricsView metrics={metricsQuery.data} />
        </div>
      )}
      <DataTable
        error={workflowKeysError}
        isLoading={workflowKeysIsLoading}
        columns={columns}
        columnVisibility={columnVisibility}
        setColumnVisibility={setColumnVisibility}
        data={listWorkflowRunsQuery.data?.rows || []}
        filters={filters}
        actions={actions}
        sorting={sorting}
        setSorting={setSorting}
        columnFilters={columnFilters}
        setColumnFilters={setColumnFilters}
        pagination={pagination}
        setPagination={setPagination}
        onSetPageSize={setPageSize}
        pageCount={listWorkflowRunsQuery.data?.pagination?.num_pages || 0}
      />
    </>
  );
}
