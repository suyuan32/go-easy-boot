// @ts-nocheck
// This file is generated by Umi automatically
// DO NOT CHANGE IT MANUALLY!
import React from 'react';
import accessFactory from '@/access'
import {useModel} from '@@/plugin-model';

import {AccessContext} from './context';

function Provider(props) {
    const {initialState} = useModel('@@initialState');
    const access = React.useMemo(() => accessFactory(initialState), [initialState]);

    return (
        <AccessContext.Provider value={access}>
            {props.children}
        </AccessContext.Provider>
    );
}

export function accessProvider(container) {
    return <Provider>{container}</Provider>;
}