package cache

import (
	"context"
	"testing"
)

type testStruct string

func (t *testStruct) UnmarshalBinary(data []byte) error {
	*t = testStruct(data)
	return nil
}

func (t testStruct) MarshalBinary() (data []byte, err error) {
	return []byte(t), nil
}

func TestSetAndGetWithRedis(t *testing.T) {
	c := &Cache{
		Driver: DRIVER__REDIS,
		Host:   "localhost",
	}
	c.Init()

	err := c.Set(context.Background(), "foo", testStruct("bar"), 0)
	if err != nil {
		t.Fatal(err)
	}

	var result testStruct
	err = c.Get(context.Background(), "foo", &result)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)

	err = c.Del(context.Background(), "foo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetAndGetWithBuildin(t *testing.T) {
	c := &Cache{
		Driver: DRIVER__BUILDIN,
	}
	c.Init()

	var val = testStruct("bar")
	err := c.Set(context.Background(), "foo", &val, 0)
	if err != nil {
		t.Fatal(err)
	}

	var result testStruct
	err = c.Get(context.Background(), "foo", &result)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)

	err = c.Del(context.Background(), "foo")
	if err != nil {
		t.Fatal(err)
	}
}
