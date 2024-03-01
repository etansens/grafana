import { Observable, ReplaySubject, Subject, firstValueFrom, map, scan, startWith } from 'rxjs';

import { PluginPreloadResult } from '../pluginPreloader';

import { PluginExtensionRegistry, PluginExtensionRegistryItem } from './types';
import { deepFreeze, isPluginCapability, logWarning } from './utils';
import { isPluginExtensionConfigValid } from './validators';

export class ReactivePluginExtensionsRegistry {
  private resultSubject: Subject<PluginPreloadResult>;
  private registrySubject: ReplaySubject<PluginExtensionRegistry>;

  constructor() {
    this.resultSubject = new Subject<PluginPreloadResult>();
    // This is the subject that we expose.
    // (It will buffer the last value on the stream - the registry - and emit it to new subscribers immediately.)
    this.registrySubject = new ReplaySubject<PluginExtensionRegistry>(1);

    this.resultSubject
      .pipe(
        scan(resultsToRegistry, {}),
        // Emit an empty object to start the stream (it is only going to do it once during construction, and then just passes down the values)
        startWith({}),
        map((registry) => deepFreeze(registry))
      )
      // Emitting the new registry to `this.registrySubject`
      .subscribe(this.registrySubject);
  }

  register(result: PluginPreloadResult): void {
    this.resultSubject.next(result);
  }

  asObservable(): Observable<PluginExtensionRegistry> {
    return this.registrySubject.asObservable();
  }

  getRegistry(): Promise<PluginExtensionRegistry> {
    return firstValueFrom(this.asObservable());
  }
}

function resultsToRegistry(registry: PluginExtensionRegistry, result: PluginPreloadResult): PluginExtensionRegistry {
  const { pluginId, extensionConfigs, error } = result;

  // TODO: We should probably move this section to where we load the plugin since this is only used
  // to provide a log to the user.
  if (error) {
    logWarning(`"${pluginId}" plugin failed to load, skip registering its extensions.`);
    return registry;
  }

  for (const extensionConfig of extensionConfigs) {
    let { extensionPointId } = extensionConfig;

    // Change the extensionPointId for capabilities
    if (isPluginCapability(extensionConfig)) {
      const regex = /capabilities\/([a-zA-Z0-9_.-]+)$/;
      const match = regex.exec(extensionPointId);

      if (!match) {
        logWarning(
          `"${pluginId}" plugin has an invalid capability ID: ${extensionPointId.replace('capabilities/', '')} (It must be a string)`
        );
        continue;
      }

      const capabilityId = match[1];

      extensionPointId = `capabilities/${pluginId}/${capabilityId}`;
      extensionConfig.extensionPointId = extensionPointId;
    }
    
    // Check if the config is valid
    if (!extensionConfig || !isPluginExtensionConfigValid(pluginId, extensionConfig)) {
      return registry;
    }

    let registryItem: PluginExtensionRegistryItem = {
      config: extensionConfig,

      // Additional meta information about the extension
      pluginId,
    };

    // Capability (only a single value per identifier, can be overriden)
    if (isPluginCapability(extensionConfig)) {
      registry[extensionPointId] = [registryItem];
    }
    // Extension (multiple extensions per extension point identifier)
    else if (!Array.isArray(registry[extensionPointId])) {
      registry[extensionPointId] = [registryItem];
    } else {
      registry[extensionPointId].push(registryItem);
    }
  }

  return registry;
}
