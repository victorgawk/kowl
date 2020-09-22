import React from 'react';
import { observer } from 'mobx-react';
import { PageComponent, PageInitHelper } from '../Page';
import { api } from '../../../state/backendApi';
import { Row, Statistic, Table } from 'antd';
import Card from '../../misc/Card';
import { appGlobal } from '../../../state/appGlobal';
import { motion } from 'framer-motion';
import { animProps } from '../../../utils/animationProps';

@observer
class SchemaList extends PageComponent<{}> {
    initPage(p: PageInitHelper): void {
        p.title = 'Schema Registry';
        p.addBreadcrumb('Schema Registry', '/schema-registry');
        this.refreshData(false);
        appGlobal.onRefresh = () => this.refreshData(true)
    }
    refreshData(force?: boolean) {
        api.refreshSchemaOverview(force);
    }
    render() {
        const { mode, compatibilityLevel, subjects} = {...api.SchemaOverview}
        return (
            <motion.div {...animProps} key={'b'} style={{ margin: '0 1rem' }}>
                <Card>
                    <Row>
                        <Statistic title="Mode" value={mode}></Statistic>
                        <Statistic title="Compatibility Level" value={compatibilityLevel}></Statistic>
                    </Row>
                </Card>
                <Card>
                    <Table
                        /* TODO: display any request errors */
                        /* TODO: styling of table rows when hovering */
                        onRow={({ name, latestVersion}) => ({
                            onClick: () => appGlobal.history.push(`/schema-registry/${name}?version=${latestVersion}`)
                        })}
                        columns={[
                            { title: 'Name', dataIndex: 'name' },
                            { title: 'Compatibility Level', dataIndex: 'compatibilityLevel' },
                            { title: 'Versions', dataIndex: 'versionsCount' },
                            { title: 'Latest Version', dataIndex: 'latestVersion' },
                        ]}
                        dataSource={subjects}
                        // TODO: Useful pagination settings
                        pagination={false}
                    ></Table>
                </Card>
            </motion.div>
        );
    }
}

export default SchemaList;
