/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import type {policy_rule} from '@fbcnms/magma-api';

import AddCircleOutline from '@material-ui/icons/AddCircleOutline';
import Button from '@fbcnms/ui/components/design-system/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import FormControl from '@material-ui/core/FormControl';
import IconButton from '@material-ui/core/IconButton';
import InputLabel from '@material-ui/core/InputLabel';
import MagmaV1API from '@fbcnms/magma-api/client/WebClient';
import PolicyFlowFields from './PolicyFlowFields';
import React from 'react';
import TextField from '@material-ui/core/TextField';
import TypedSelect from '@fbcnms/ui/components/TypedSelect';
import Typography from '@material-ui/core/Typography';
import nullthrows from '@fbcnms/util/nullthrows';
import {ACTION, DIRECTION, PROTOCOL} from './PolicyTypes';
import {CWF, FEG, LTE} from '@fbcnms/types/network';
import {coalesceNetworkType} from '@fbcnms/types/network';
import {makeStyles} from '@material-ui/styles';
import {useEffect, useState} from 'react';
import {useRouter} from '@fbcnms/ui/hooks';
import type {NetworkType} from '@fbcnms/types/network';

const useStyles = makeStyles(() => ({
  input: {width: '100%'},
}));

type Props = {
  onCancel: () => void,
  onSave: string => void,
  rule?: policy_rule,
  mirrorNetwork?: string,
};

export default function PolicyRuleEditDialog(props: Props) {
  const classes = useStyles();
  const {match} = useRouter();
  const [networkType, setNetworkType] = useState<?NetworkType>(null);
  const [mirrorNetworkType, setMirrorNetworkType] = useState<?NetworkType>(
    null,
  );
  const [networkWideRuleIDs, setNetworkWideRuldIDs] = useState(null);
  const [isNetworkWide, setIsNetworkWide] = useState<boolean>(false);
  const {networkId} = match.params;
  const {mirrorNetwork} = props;

  // Grab the network type for the network, and the mirrorNetwork if it exists.
  useEffect(() => {
    getNetworkType(networkId, setNetworkType);
    getNetworkType(mirrorNetwork, setMirrorNetworkType);
  }, [networkId, mirrorNetwork]);

  // Grab the network wide rule IDs from the network's subscriber config
  // Case on the network type to determine which endpoint to call.
  useEffect(() => {
    switch (networkType) {
      case LTE:
        MagmaV1API.getLteByNetworkIdSubscriberConfigRuleNames({
          networkId: networkId,
        }).then(ruleIDs => setNetworkWideRuldIDs(ruleIDs));
        break;
      case CWF:
        MagmaV1API.getCwfByNetworkIdSubscriberConfigRuleNames({
          networkId: networkId,
        }).then(ruleIDs => {
          setNetworkWideRuldIDs(ruleIDs);
        });
        break;
      case FEG:
        MagmaV1API.getFegByNetworkIdSubscriberConfigRuleNames({
          networkId: networkId,
        }).then(ruleIDs => setNetworkWideRuldIDs(ruleIDs));
        break;
    }
  }, [networkId, networkType]);

  // The rule is network-wide if the rule's ID exists in network-wide rule IDs
  useEffect(() => {
    networkWideRuleIDs
      ? setIsNetworkWide(networkWideRuleIDs.includes(props.rule?.id))
      : false;
  }, [networkWideRuleIDs, props.rule]);

  const [rule, setRule] = useState(
    props.rule || {
      id: '',
      priority: 1,
      flow_list: [],
      rating_group: 0,
      monitoring_key: '',
    },
  );

  const handleAddFlow = () => {
    const flowList = [
      ...(rule.flow_list || []),
      {
        action: ACTION.DENY,
        match: {
          direction: DIRECTION.UPLINK,
          ip_proto: PROTOCOL.IPPROTO_IP,
        },
      },
    ];

    setRule({...rule, flow_list: flowList});
  };

  const onFlowChange = (index, flow) => {
    const flowList = [...(rule.flow_list || [])];
    flowList[index] = flow;
    setRule({...rule, flow_list: flowList});
  };

  const handleDeleteFlow = (index: number) => {
    const flowList = [...(rule.flow_list || [])];
    flowList.splice(index, 1);
    setRule({...rule, flow_list: flowList});
  };

  const onSave = async () => {
    const ruleData = [
      {
        networkId: nullthrows(networkId),
        ruleId: rule.id,
        policyRule: rule,
      },
    ];

    if (mirrorNetwork != null) {
      ruleData.push({
        networkId: mirrorNetwork,
        ruleId: rule.id,
        policyRule: rule,
      });
    }

    if (props.rule) {
      await Promise.all(
        ruleData.map(d =>
          MagmaV1API.putNetworksByNetworkIdPoliciesRulesByRuleId(d),
        ),
      );
    } else {
      await Promise.all(
        ruleData.map(d => MagmaV1API.postNetworksByNetworkIdPoliciesRules(d)),
      );
    }

    const networkWideRuleData = {
      networkId: nullthrows(networkId),
      ruleId: rule.id,
    };
    if (isNetworkWide) {
      await postNetworkWideRuleID(networkWideRuleData, networkType);
      if (mirrorNetwork != null) {
        networkWideRuleData.networkId = mirrorNetwork;
        await postNetworkWideRuleID(networkWideRuleData, mirrorNetworkType);
      }
    } else {
      await deleteNetworkWideRuleID(networkWideRuleData, networkType);
      if (mirrorNetwork != null) {
        networkWideRuleData.networkId = mirrorNetwork;
        await deleteNetworkWideRuleID(networkWideRuleData, mirrorNetworkType);
      }
    }

    props.onSave(rule.id);
  };

  return (
    <Dialog open={true} onClose={props.onCancel} scroll="body">
      <DialogTitle>{props.rule ? 'Edit' : 'Add'} Rule</DialogTitle>
      <DialogContent>
        <TextField
          required
          className={classes.input}
          label="ID"
          margin="normal"
          disabled={!!props.rule}
          value={rule.id}
          onChange={({target}) => setRule({...rule, id: target.value})}
        />
        <TextField
          required
          className={classes.input}
          label="Precendence"
          margin="normal"
          value={rule.priority}
          onChange={({target}) =>
            setRule({...rule, priority: parseInt(target.value)})
          }
        />
        <TextField
          required
          className={classes.input}
          label="Monitoring Key"
          margin="normal"
          value={rule.monitoring_key}
          onChange={({target}) =>
            setRule({...rule, monitoring_key: target.value})
          }
        />
        <TextField
          required
          className={classes.input}
          label="Rating Group"
          margin="normal"
          value={rule.rating_group}
          type="number"
          onChange={({target}) =>
            setRule({
              ...rule,
              rating_group: parseInt(target.value),
            })
          }
        />
        <FormControl className={classes.input}>
          <InputLabel htmlFor="trackingType">Tracking Type</InputLabel>
          <TypedSelect
            items={{
              ONLY_OCS: 'Only OCS',
              ONLY_PCRF: 'Only PCRF',
              OCS_AND_PCRF: 'OCS and PCRF',
              NO_TRACKING: 'No Tracking',
            }}
            inputProps={{id: 'trackingType'}}
            value={rule.tracking_type || 'NO_TRACKING'}
            onChange={trackingType =>
              setRule({...rule, tracking_type: trackingType})
            }
          />
        </FormControl>
        <FormControl className={classes.input}>
          <InputLabel htmlFor="target">Network Wide</InputLabel>
          <TypedSelect
            items={{
              true: 'true',
              false: 'false',
            }}
            inputProps={{id: 'target'}}
            value={isNetworkWide ? 'true' : 'false'}
            onChange={target => {
              setIsNetworkWide(target === 'true');
            }}
          />
        </FormControl>
        <Typography variant="h6">
          Flows
          <IconButton onClick={handleAddFlow}>
            <AddCircleOutline />
          </IconButton>
        </Typography>
        {(rule.flow_list || []).slice(0, 30).map((flow, i) => (
          <PolicyFlowFields
            key={i}
            index={i}
            flow={flow}
            handleDelete={handleDeleteFlow}
            onChange={onFlowChange}
          />
        ))}
      </DialogContent>
      <DialogActions>
        <Button onClick={props.onCancel} skin="regular">
          Cancel
        </Button>
        <Button onClick={onSave}>Save</Button>
      </DialogActions>
    </Dialog>
  );
}

async function postNetworkWideRuleID(networkWideRuleData, networkType) {
  switch (networkType) {
    case LTE:
      MagmaV1API.postLteByNetworkIdSubscriberConfigRuleNamesByRuleId(
        networkWideRuleData,
      );
      break;
    case CWF:
      MagmaV1API.postCwfByNetworkIdSubscriberConfigRuleNamesByRuleId(
        networkWideRuleData,
      );
      break;
    case FEG:
      MagmaV1API.postFegByNetworkIdSubscriberConfigRuleNamesByRuleId(
        networkWideRuleData,
      );
      break;
  }
}

async function deleteNetworkWideRuleID(networkWideRuleData, networkType) {
  switch (networkType) {
    case LTE:
      MagmaV1API.deleteLteByNetworkIdSubscriberConfigRuleNamesByRuleId(
        networkWideRuleData,
      );
      break;
    case CWF:
      MagmaV1API.deleteCwfByNetworkIdSubscriberConfigRuleNamesByRuleId(
        networkWideRuleData,
      );
      break;
    case FEG:
      MagmaV1API.deleteFegByNetworkIdSubscriberConfigRuleNamesByRuleId(
        networkWideRuleData,
      );
      break;
  }
}

function getNetworkType(
  networkId: ?string,
  setNetworkType: (?NetworkType) => void,
) {
  if (networkId != null) {
    MagmaV1API.getNetworksByNetworkIdType({networkId}).then(networkType =>
      setNetworkType(coalesceNetworkType(networkId, networkType)),
    );
  }
}
