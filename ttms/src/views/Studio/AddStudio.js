import React, { useState, useEffect } from 'react'
import { Button, Drawer, Form, Input, message, Modal } from 'antd'
import { $addstudio, $revise, $getOne } from '../../api/studioapi';
import { ExclamationCircleFilled } from '@ant-design/icons'
const { confirm } = Modal;

export default function AddStudio({ open, setOpen, loadList, Id, setId }) {
    //提示框
    const [messageApi, contextHolder] = message.useMessage();
    //是否打开抽屉
    // const [open, setOpen] = useState(false);
    //定义表单事例
    let [form] = Form.useForm()
    useEffect(() => {
        if (Id !== 0) {
            $getOne({ Id }).then(data => {
                form.setFieldsValue(data)
            })
        }
    }, [Id])

    //关闭抽屉的方法
    const onClose = () => {
        clear();
        setId(0)
        setOpen(false);
    };
    //表单提交的方法
    const onFinish = async (values) => {
        if (Id) {
            $revise(values).then(data => {
                if (data.BaseResp.StatusMessage === 'success') {
                    messageApi.open({
                        type: 'success',
                        content: '修改成功',
                    });
                    loadList()//重新加载数据
                } else {
                    messageApi.open({
                        type: 'error',
                        content: '修改失败座位规模不能比原来更大',
                    });
                }
            })
        } else {
            let success = await $addstudio(values)
            // console.log(success.BaseResp.StatusMessage)
            // .then(success, message => {
            if (success.BaseResp.StatusMessage === 'success') {
                confirm({
                    title: '系统提示',
                    icon: <ExclamationCircleFilled />,
                    content: '添加成功',
                    okText: '确定',
                    cancelText: '取消',
                });
                clear()//清空表单
                loadList()//加载表单
            } else {
                confirm({
                    title: '系统提示',
                    icon: <ExclamationCircleFilled />,
                    content: '添加失败',
                    okText: '确定',
                    cancelText: '取消',
                });
            }
            // })
            // console.log('Success:', values);
        }

    };
    //清空表单方法
    const clear = () => {
        form.setFieldsValue({ Name: '', RowsCount: '', ColsCount: '' })
    }
    return (
        <>
            {contextHolder}
            <Drawer title={Id ? '修改演出厅' : '添加演出厅'} width={500} placement="right" onClose={onClose} open={open}>
                <Form
                    name="basic"
                    form={form}
                    labelCol={{
                        span: 6,
                    }}
                    wrapperCol={{
                        span: 16,
                    }}
                    style={{
                        maxWidth: 600,
                    }}

                    onFinish={onFinish}
                    autoComplete="off"
                >
                    <Form.Item
                        label="演出厅编号"
                        name="Id"
                        hidden
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        label="演出厅名称"
                        name="Name"
                        rules={[
                            {
                                required: true,
                                message: '请输入演出厅名称!',
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="座位行数"
                        name="RowsCount"
                        rules={[
                            {
                                required: true,
                                message: '请输座位行数!',
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        label="座位列数"
                        name="ColsCount"
                        rules={[
                            {
                                required: true,
                                message: '请输入座位列数!',
                            },
                        ]}
                    >
                        <Input />
                    </Form.Item>


                    <Form.Item
                        wrapperCol={{
                            offset: 9,
                            span: 16,
                        }}
                    >
                        <Button type="primary" htmlType="submit">
                            {Id ? '修改' : '添加'}
                        </Button>
                        <Button onClick={clear} style={{ marginLeft: '10px' }} >
                            取消
                        </Button>
                    </Form.Item>
                </Form>
            </Drawer>
        </>
    )
}
