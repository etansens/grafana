import React from 'react';

import { QueryBuilderLabelFilter } from '@grafana/prometheus/src/querybuilder/shared/types';
import { getDataSourceSrv } from '@grafana/runtime';
import { SceneComponentProps, SceneObjectBase, SceneObjectState, SceneTimeRangeLike } from '@grafana/scenes';
import { DataSourceRef } from '@grafana/schema';
import { Drawer } from '@grafana/ui';
import appEvents from 'app/core/app_events';
import { ShowModalReactEvent } from 'app/types/events';

import { getDashboardSceneFor } from '../dashboard-scene/utils/utils';

import { DataTrail } from './DataTrail';
import { getDataTrailsApp } from './DataTrailsApp';
import { OpenEmbeddedTrailEvent } from './shared';

interface DataTrailDrawerState extends SceneObjectState {
  timeRange: SceneTimeRangeLike;
  metric: string;
  labels?: QueryBuilderLabelFilter[];
  queries?: string[];
  dsRef: DataSourceRef;
}

export class DataTrailDrawer extends SceneObjectBase<DataTrailDrawerState> {
  static Component = DataTrailDrawerRenderer;

  public trail: DataTrail;

  constructor(state: DataTrailDrawerState) {
    super(state);

    this.trail = buildDataTrailFromQuery(state);
    this.trail.addActivationHandler(() => {
      this.trail.subscribeToEvent(OpenEmbeddedTrailEvent, this.onOpenTrail);
    });
  }

  onOpenTrail = () => {
    getDataTrailsApp().goToUrlForTrail(this.trail.clone({ embedded: false }));
  };

  onClose = () => {
    try {
      const dashboard = getDashboardSceneFor(this);
      dashboard.closeModal();
    } catch (e: unknown) {
      const event = new ShowModalReactEvent({ component: () => null });
      appEvents.publish(event);
    }
  };
}

function DataTrailDrawerRenderer({ model }: SceneComponentProps<DataTrailDrawer>) {
  return (
    <Drawer title={'Data trail'} onClose={model.onClose} size="lg">
      <div style={{ display: 'flex', height: '100%' }}>
        <model.trail.Component model={model.trail} />
      </div>
    </Drawer>
  );
}

export function buildDataTrailFromQuery({ metric, labels, dsRef, timeRange }: DataTrailDrawerState) {
  const filters = labels?.map((label) => ({ key: label.label, value: label.value, operator: label.op }));

  const ds = getDataSourceSrv().getInstanceSettings(dsRef);

  return new DataTrail({
    $timeRange: timeRange,
    metric,
    initialDS: ds?.uid,
    initialFilters: filters,
    embedded: true,
  });
}
