// This file is generated by Umi automatically
// DO NOT CHANGE IT MANUALLY!
import type {
    ProLayoutProps
} from "/home/ryan/GolandProjects/test/go-easy-boot-web/node_modules/_@ant-design_pro-layout@7.0.1-beta.28@@ant-design/pro-layout";
import type InitialStateType from '@@/plugin-initialState/@@initialState';

type InitDataType = ReturnType<typeof InitialStateType>;


export type RunTimeLayoutConfig = (
    initData: InitDataType,
) => ProLayoutProps & {
    childrenRender?: (dom: JSX.Element, props: ProLayoutProps) => React.ReactNode,
    unAccessible?: JSX.Element,
    noFound?: JSX.Element,
};
