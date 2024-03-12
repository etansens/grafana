import { PanelMenuItem, PanelModel } from '@grafana/data';
import { QueryBuilderLabelFilter } from '@grafana/prometheus/src/querybuilder/shared/types';
import { getDataSourceSrv } from '@grafana/runtime';
import { SceneQueryRunner, SceneTimeRange, VizPanel } from '@grafana/scenes';
import appEvents from 'app/core/app_events';
import { buildVisualQueryFromString } from 'app/plugins/datasource/prometheus/querybuilder/parsing';
import { ShowModalReactEvent } from 'app/types/events';

import { DashboardModel } from '../dashboard/state';
import { DashboardScene } from '../dashboard-scene/scene/DashboardScene';
import { getQueryRunnerFor } from '../dashboard-scene/utils/utils';

import { DataTrailDrawer } from './DataTrailDrawer';

export function addDataTrailPanelAction(
  dashboard: DashboardScene | DashboardModel,
  panel: VizPanel | PanelModel,
  items: PanelMenuItem[]
) {
  const queryRunner =
    panel instanceof VizPanel
      ? getQueryRunnerFor(panel)
      : new SceneQueryRunner({ datasource: panel.datasource || undefined, queries: panel.targets || [] });
  if (!queryRunner) {
    return;
  }

  const ds = getDataSourceSrv().getInstanceSettings(queryRunner.state.datasource);

  if (!ds || ds.meta.id !== 'prometheus' || queryRunner.state.queries.length < 1) {
    return;
  }

  if (ds === undefined) {
    return;
  }

  const queries = queryRunner.state.queries.map((q) => q.expr);

  type MetricAndLabels = {
    metric: string;
    labels: QueryBuilderLabelFilter[];
    queries: string[];
  };

  const metricItems: MetricAndLabels[] = [];

  // We only support label filters with the '=' operator
  function isEquals(labelFilter: QueryBuilderLabelFilter) {
    return labelFilter.op === '=';
  }

  queries.forEach((query) => {
    const struct = buildVisualQueryFromString(query);
    if (struct.errors.length > 0) {
      return;
    }

    const { metric, labels } = struct.query;

    metricItems.push({ metric, labels: labels.filter(isEquals), queries });
    struct.query.binaryQueries?.forEach(({ query }) => {
      const { metric, labels } = query;
      metricItems.push({ metric, labels: labels.filter(isEquals), queries });
    });
  });

  const getClickHandler = ({ metric, labels, queries }: MetricAndLabels) => {
    const timeRange =
      dashboard instanceof DashboardScene
        ? dashboard.state.$timeRange!.clone()
        : new SceneTimeRange({ ...dashboard.time });

    const drawer = new DataTrailDrawer({
      metric,
      labels,
      queries,
      dsRef: ds,
      timeRange,
    });

    if (dashboard instanceof DashboardScene) {
      return () => dashboard.showModal(drawer);
    } else {
      return () => {
        const payload = {
          component: DataTrailDrawer.Component,
          props: { model: drawer },
        };

        appEvents.publish(new ShowModalReactEvent(payload));
      };
    }
  };

  const uniqueMenuTexts = new Set<string>();
  function isUnique({ text }: { text: string }) {
    const before = uniqueMenuTexts.size;
    uniqueMenuTexts.add(text);
    const after = uniqueMenuTexts.size;
    return after > before;
  }

  const subMenu: PanelMenuItem[] = metricItems.map((item) => ({
    text: `${item.metric}${item.labels.length === 0 ? '' : `{${item.labels.map(({ label, op, value }) => `${label}${op}${value}`)}}`}`,
    onClick: getClickHandler(item),
  }));

  items.push({
    text: 'Explore metrics',
    iconClassName: 'code-branch',
    shortcut: 'p m',
    subMenu: subMenu.filter(isUnique),
  });
}
