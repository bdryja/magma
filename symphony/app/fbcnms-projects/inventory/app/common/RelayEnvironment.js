/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {StoreUpdater as RelayStoreUpdater} from 'relay-runtime';

import axios from 'axios';
import {Environment, Network, RecordSource, Store} from 'relay-runtime';

function fetchQuery(operation, variables) {
  return axios
    .post('/graph/query', {
      query: operation.text,
      variables,
    })
    .then(response => {
      return response.data;
    });
}

const RelayEnvironment = new Environment({
  network: Network.create(fetchQuery),
  store: new Store(new RecordSource()),
});

export default RelayEnvironment;

export type StoreUpdater = RelayStoreUpdater;
