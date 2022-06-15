import React, { ReactNode } from 'react';
import { TrashIcon } from '@heroicons/react/outline';
import { Alert, Button, Checkbox, Col, Modal, notification, Popover, Row, Statistic, Table, Tooltip } from 'antd';
import { motion } from 'framer-motion';
import { autorun, computed, IReactionDisposer, makeObservable, observable } from 'mobx';
import { observer } from 'mobx-react';
import { appGlobal } from '../../../state/appGlobal';
import { api } from '../../../state/backendApi';
import { Topic, TopicAction, TopicActions, TopicConsumerGroupSummary, ConsumerGroupSummary } from '../../../state/restInterfaces';
import { uiSettings } from '../../../state/ui';
import { animProps } from '../../../utils/animationProps';
import { editQuery } from '../../../utils/queryHelper';
import { DefaultSkeleton, findPopupContainer, QuickTable } from '../../../utils/tsxUtils';
import Card from '../../misc/Card';
import { makePaginationConfig, renderLogDirSummary, sortField } from '../../misc/common';
import { KowlTable } from '../../misc/KowlTable';
import SearchBar from '../../misc/SearchBar';
import { PageComponent, PageInitHelper } from '../Page';
import { useState } from 'react';
import { CheckIcon, CircleSlashIcon, EyeClosedIcon } from '@primer/octicons-react';

@observer
class TopicList extends PageComponent {
    pageConfig = makePaginationConfig(uiSettings.topicList.pageSize);
    quickSearchReaction: IReactionDisposer;
    @observable topicToDelete: null | string = null;

    constructor(p: any) {
        super(p);
        makeObservable(this);
    }

    initPage(p: PageInitHelper): void {
        p.title = 'Topics';
        p.addBreadcrumb('Topics', '/topics');

        this.refreshData(false);
        appGlobal.onRefresh = () => this.refreshData(true);
    }

    componentDidMount() {
        // 1. use 'q' parameter for quick search (if it exists)
        editQuery((query) => {
            if (query['q']) uiSettings.topicList.quickSearch = String(query['q']);
        });

        // 2. whenever the quick search box changes, update the url
        this.quickSearchReaction = autorun(() => {
            editQuery((query) => {
                const q = String(uiSettings.topicList.quickSearch);
                if (q) query['q'] = q;
            });
        });
    }
    componentWillUnmount() {
        if (this.quickSearchReaction) this.quickSearchReaction();
    }

    refreshData(force: boolean) {
        api.refreshTopics(force);
    }

    isFilterMatch(filter: string, item: Topic): boolean {
        if (item.topicName.toLowerCase().includes(filter.toLowerCase())) return true;
        return false;
    }

    render() {
        if (!api.topics) return DefaultSkeleton;

        const topics = api.topics;

        const partitionCountReal = topics.sum((x) => x.partitionCount);
        const partitionCountOnlyReplicated = topics.sum((x) => x.partitionCount * (x.replicationFactor - 1));

        const partitionDetails = QuickTable(
            [
                { key: 'Primary:', value: partitionCountReal },
                { key: 'Replicated:', value: partitionCountOnlyReplicated },
                { key: 'All:', value: partitionCountReal + partitionCountOnlyReplicated },
            ],
            {
                keyAlign: 'right', keyStyle: { fontWeight: 500 },
                gapWidth: 4,
                valueAlign: 'right'
            }
        );

        return (
                <motion.div {...animProps} style={{ margin: '0 1rem' }}>
                    <Card>
                        <Row>
                            <Statistic title="Total Topics" value={topics.length} />
                            <Popover title="Partition Details" content={partitionDetails} placement="right" mouseEnterDelay={0} trigger="hover">
                                <div className="hoverLink" style={{ display: 'flex', verticalAlign: 'middle', cursor: 'default' }}>
                                    <Statistic title="Total Partitions" value={partitionCountReal + partitionCountOnlyReplicated} />
                                </div>
                            </Popover>
                        </Row>
                    </Card>

                    <Card>
                        <KowlTable
                            dataSource={topics}
                            rowKey={(x) => x.topicName}
                            columns={[
                                { title: 'Name', dataIndex: 'topicName', render: (t, r) => renderName(r), sorter: sortField('topicName'), className: 'whiteSpaceDefault', defaultSortOrder: 'ascend' },
                                { title: 'Partitions', dataIndex: 'partitions', render: (t, r) => r.partitionCount, sorter: (a, b) => a.partitionCount - b.partitionCount, width: 1 },
                                { title: 'Replication', dataIndex: 'replicationFactor', sorter: sortField('replicationFactor'), width: 1 },
                                {
                                    title: 'CleanupPolicy', dataIndex: 'cleanupPolicy', width: 1,
                                    filterType: {
                                        type: 'enum',
                                        optionClassName: 'capitalize'
                                    },
                                    sorter: sortField('cleanupPolicy'),
                                },
                                {
                                    title: 'Size', render: (t, r) => renderLogDirSummary(r.logDirSummary), sorter: (a, b) => a.logDirSummary.totalSizeBytes - b.logDirSummary.totalSizeBytes, width: '70px',
                                },
                                {
                                    title: 'Current Messages', render: (t, r) => r.currentMessages, sorter: (a, b) => a.currentMessages - b.currentMessages, width: '140px',
                                },
                                {
                                    title: 'Total Messages', render: (t, r) => r.totalMessages, sorter: (a, b) => a.totalMessages - b.totalMessages, width: '140px',
                                },
                                {
                                    title: 'Lag', render: (t, r) => renderConsumerGroupSummary(r.consumerGroupSummary), sorter: (a, b) => {
                                        const maxLagA = a.consumerGroupSummary.maxLag;
                                        const maxLagB = b.consumerGroupSummary.maxLag;
                                        const hasConsumersA = a.consumerGroupSummary != null && a.consumerGroupSummary.consumerGroups != null;
                                        const hasConsumersB = b.consumerGroupSummary != null && b.consumerGroupSummary.consumerGroups != null;
                                        return (hasConsumersA ? maxLagA : -1) - (hasConsumersB ? maxLagB : -1);
                                    }, width: '70px',
                                },
                                {
                                    width: 1,
                                    title: ' ',
                                    key: 'action',
                                    className: 'msgTableActionColumn',
                                    render: (text, record) => (
                                        <div style={{ display: 'flex', gap: '4px' }}>
                                            <DeleteDisabledTooltip topic={record}>
                                                <Button
                                                    type="text"
                                                    className="iconButton"
                                                    onClick={(event) => {
                                                        event.stopPropagation();
                                                        this.topicToDelete = record.topicName;
                                                    }}
                                                    danger
                                                >
                                                    <TrashIcon />
                                                </Button>
                                            </DeleteDisabledTooltip>
                                        </div>
                                    ),
                                },
                            ]}

                            search={{
                                searchColumnIndex: 0,
                                isRowMatch: (row, regex) => {
                                    if (regex.test(row.topicName)) return true;
                                    if (regex.test(row.cleanupPolicy)) return true;
                                    return false;
                                },

                            }}

                            observableSettings={uiSettings.topicList}
                            onRow={(record) => ({
                                onClick: () => appGlobal.history.push('/topics/' + record.topicName),
                            })}
                            rowClassName="hoverLink"
                        />
                    </Card>
                    <ConfirmDeletionModal
                        topicToDelete={this.topicToDelete}
                        onCancel={() => (this.topicToDelete = null)}
                        onFinish={() => {
                            this.topicToDelete = null;
                            this.refreshData(true);
                        }}
                    />
                </motion.div>
        );
    }
}

const iconAllowed = (
    <span style={{ color: 'green' }}>
        <CheckIcon size={16} />
    </span>
);
const iconForbidden = (
    <span style={{ color: '#ca000a' }}>
        <CircleSlashIcon size={15} />
    </span>
);
const iconClosedEye = (
    <span style={{ color: '#0008', paddingLeft: '4px', transform: 'translateY(-1px)', display: 'inline-block' }}>
        <EyeClosedIcon size={14} verticalAlign="middle" />
    </span>
);

const renderName = (topic: Topic) => {
    const actions = topic.allowedActions;

    if (!actions || actions[0] == 'all') return topic.topicName; // happens in non-business version

    let missing = 0;
    for (const a of TopicActions) if (!actions.includes(a)) missing++;

    if (missing == 0) return topic.topicName; // everything is allowed

    // There's at least one action the user can't do
    // Show a table of what they can't do
    const popoverContent = (
        <div>
            <div style={{ marginBottom: '1em' }}>
                You're missing permissions to view
                <br />
                one more aspects of this topic.
            </div>
            {QuickTable(
                TopicActions.map((a) => ({
                    key: a,
                    value: actions.includes(a) ? iconAllowed : iconForbidden,
                })),
                {
                    gapWidth: '6px',
                    gapHeight: '2px',
                    keyAlign: 'right',
                    keyStyle: { fontSize: '86%', fontWeight: 700, textTransform: 'capitalize' },
                    tableStyle: { margin: 'auto' },
                }
            )}
        </div>
    );

    return (
        <Popover content={popoverContent} placement="right" mouseEnterDelay={0.1} mouseLeaveDelay={0.1}>
            <span>
                {topic.topicName}
                {iconClosedEye}
            </span>
        </Popover>
    );
};

function ConfirmDeletionModal({ topicToDelete, onFinish, onCancel }: { topicToDelete: string | null; onFinish: () => void; onCancel: () => void }) {
    const [deletionPending, setDeletionPending] = useState(false);
    const [error, setError] = useState<string | Error | null>(null);

    const cleanup = () => {
        setDeletionPending(false);
        setError(null);
    };

    const finish = () => {
        onFinish();
        cleanup();
        notification['success']({
            message: `Topic \`${topicToDelete}\` deleted successfully`,
        });
    };

    const cancel = () => {
        onCancel();
        cleanup();
    };

    return (
        <Modal
            className="topicDeleteModal"
            visible={topicToDelete != null}
            centered
            closable={false}
            maskClosable={!deletionPending}
            keyboard={!deletionPending}
            okText={error ? 'Retry' : 'Yes'}
            confirmLoading={deletionPending}
            okType="danger"
            cancelText="No"
            cancelButtonProps={{ disabled: deletionPending }}
            onCancel={cancel}
            onOk={() => {
                setDeletionPending(true);
                api.deleteTopic(topicToDelete!) // modal is not shown when topic is null
                    .then(finish)
                    .catch(setError)
                    .finally(() => { setDeletionPending(false) });
            }}
        >
            <>
                {error && <Alert type="error" message={`An error occurred: ${typeof error === 'string' ? error : error.message}`} />}
                <p>
                    Are you sure you want to delete topic <strong>{topicToDelete}</strong>? This action is irrevocable.
                </p>
            </>
        </Modal>
    );
}

function DeleteDisabledTooltip(props: { topic: Topic; children: JSX.Element }): JSX.Element {
    const { topic } = props;
    const deleteButton = props.children;

    const wrap = (button: JSX.Element, message: string) => (
        <Tooltip placement="top" trigger="hover" mouseLeaveDelay={0} getPopupContainer={findPopupContainer} overlay={message}>
            {React.cloneElement(button, {
                disabled: true,
                className: (button.props.className ?? '') + ' disabled',
                onClick: undefined,
            })}
        </Tooltip>
    );

    if (uiSettings.adminOperations.readOnlyMode) return <>{wrap(deleteButton, "Read only mode.")}</>;

    return <>{hasDeletePrivilege(topic.allowedActions) ? deleteButton : wrap(deleteButton, "You don't have 'deleteTopic' permission for this topic.")}</>;
}

function hasDeletePrivilege(allowedActions?: Array<TopicAction>) {
    return Boolean(allowedActions?.includes('all') || allowedActions?.includes('deleteTopic'));
}

function renderConsumerGroupSummary(summary: TopicConsumerGroupSummary): JSX.Element {
  if (summary == null || summary.consumerGroups == null) return <></>;
  const msg = QuickTable(
      summary.consumerGroups.map((consumerGroup) => ({ key: consumerGroup.groupName + ":", value: consumerGroup.lag })),
      {
          keyAlign: 'right', keyStyle: { fontWeight: 500 },
          gapWidth: 4,
          valueAlign: 'right'
      }
  );
  return <>
    <Popover title="Consumer Groups" content={msg} placement="rightTop" overlayClassName="popoverSmall">
      {summary.maxLag}
    </Popover>
  </>;
}

export default TopicList;
