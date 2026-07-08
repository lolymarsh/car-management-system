import { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import {
  Card, Form, Input, InputNumber, Button, message, Spin,
} from 'antd';
import api from '../api';

export default function CarForm() {
  const { car_id } = useParams();
  const navigate = useNavigate();
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const isEdit = Boolean(car_id);

  useEffect(() => {
    if (!isEdit) return;
    setLoading(true);
    api.get(`/api/cars/${car_id}`)
      .then((res) => {
        form.setFieldsValue(res.data);
      })
      .catch(() => {
        message.error('Failed to load car');
        navigate('/cars');
      })
      .finally(() => setLoading(false));
  }, [car_id, isEdit, form, navigate]);

  const handleSubmit = async (values) => {
    setSubmitting(true);
    try {
      if (isEdit) {
        await api.put(`/api/cars/${car_id}`, values);
        message.success('Car updated');
      } else {
        await api.post('/api/cars', values);
        message.success('Car created');
      }
      navigate('/cars');
    } catch {
      message.error('Failed to save car');
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) return <Spin size="large" style={{ display: 'block', margin: '100px auto' }} />;

  return (
    <Card title={isEdit ? 'Edit Car' : 'Add Car'}>
      <Form
        form={form}
        layout="vertical"
        onFinish={handleSubmit}
        style={{ maxWidth: 600 }}
      >
        <Form.Item
          name="registration_number"
          label="Registration Number"
          rules={[{ required: true, message: 'Please enter registration number' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="brand"
          label="Brand"
          rules={[{ required: true, message: 'Please enter brand' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="model"
          label="Model"
          rules={[{ required: true, message: 'Please enter model' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item name="color" label="Color">
          <Input />
        </Form.Item>

        <Form.Item name="year" label="Year">
          <InputNumber min={1886} max={2100} style={{ width: '100%' }} />
        </Form.Item>

        <Form.Item name="notes" label="Notes">
          <Input.TextArea rows={4} />
        </Form.Item>

        <Form.Item>
          <Space>
            <Button type="primary" htmlType="submit" loading={submitting}>
              {isEdit ? 'Update' : 'Create'}
            </Button>
            <Button onClick={() => navigate('/cars')}>Cancel</Button>
          </Space>
        </Form.Item>
      </Form>
    </Card>
  );
}
