import { useState, useEffect, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
import {
  Table, Button, Input, Space, Card, Popconfirm, message, Tag, Empty,
} from 'antd';
import {
  PlusOutlined, SearchOutlined, DeleteOutlined, EditOutlined,
} from '@ant-design/icons';
import api from '../api';

export default function CarList() {
  const navigate = useNavigate();
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);
  const [search, setSearch] = useState('');

  const fetchData = useCallback(async () => {
    setLoading(true);
    try {
      const res = await api.post('/api/cars/filter', {
        search,
        page,
        page_size: pageSize,
        sort_by: 'created_at',
        sort_dir: 'desc',
      });
      setData(res?.data?.cars);
      setTotal(res?.data?.total);
    } catch {
      message.error('Failed to fetch cars');
    } finally {
      setLoading(false);
    }
  }, [search, page, pageSize]);

  useEffect(() => {
    fetchData();
  }, []);

  const handleDelete = async (id) => {
    try {
      await api.delete(`/api/cars/${id}`);
      message.success('Car deleted');
      fetchData();
    } catch {
      message.error('Failed to delete car');
    }
  };

  const columns = [
    {
      title: 'Registration No.',
      dataIndex: 'registration_number',
      key: 'registration_number',
      sorter: true,
    },
    {
      title: 'Brand',
      dataIndex: 'brand',
      key: 'brand',
      sorter: true,
    },
    {
      title: 'Model',
      dataIndex: 'model',
      key: 'model',
      sorter: true,
    },
    {
      title: 'Color',
      dataIndex: 'color',
      key: 'color',
      render: (text) => text || <Tag>N/A</Tag>,
    },
    {
      title: 'Year',
      dataIndex: 'year',
      key: 'year',
      sorter: true,
      render: (text) => text || '-',
    },
    {
      title: 'Notes',
      dataIndex: 'notes',
      key: 'notes',
      ellipsis: true,
    },
    {
      title: 'Actions',
      key: 'actions',
      render: (_, record) => (
        <Space>
          <Button
            type="link"
            icon={<EditOutlined />}
            onClick={() => navigate(`/cars/${record.car_id}/edit`)}
          >
            Edit
          </Button>
          <Popconfirm
            title="Delete this car?"
            onConfirm={() => handleDelete(record.car_id)}
            okText="Yes"
            cancelText="No"
          >
            <Button type="link" danger icon={<DeleteOutlined />}>
              Delete
            </Button>
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <Card
      title="Car List"
      extra={
        <Space>
          <Input.Search
            placeholder="Search..."
            prefix={<SearchOutlined />}
            allowClear
            onSearch={(value) => {
              setSearch(value);
              setPage(1);
            }}
            style={{ width: 250 }}
          />
          <Button
            type="primary"
            icon={<PlusOutlined />}
            onClick={() => navigate('/cars/new')}
          >
            Add Car
          </Button>
        </Space>
      }
    >
      <Table
        columns={columns}
        dataSource={data}
        rowKey="car_id"
        loading={loading}
        pagination={{
          current: page,
          pageSize,
          total,
          showSizeChanger: true,
          showTotal: (t) => `Total ${t} items`,
          onChange: (p, ps) => {
            setPage(p);
            setPageSize(ps);
          },
        }}
        locale={{
          emptyText: <Empty description="No cars found" />,
        }}
      />
    </Card>
  );
}
