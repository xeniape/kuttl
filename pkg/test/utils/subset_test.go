package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubset(t *testing.T) {
	assert.Nil(t, IsSubset(map[string]interface{}{
		"hello": "world",
	}, map[string]interface{}{
		"hello": "world",
		"bye":   "moon",
	}))

	assert.NotNil(t, IsSubset(map[string]interface{}{
		"hello": "moon",
	}, map[string]interface{}{
		"hello": "world",
		"bye":   "moon",
	}))

	assert.Nil(t, IsSubset(map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": "world",
		},
	}, map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": "world",
			"bye":   "moon",
		},
	}))

	assert.NotNil(t, IsSubset(map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": "moon",
		},
	}, map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": "world",
			"bye":   "moon",
		},
	}))

	assert.NotNil(t, IsSubset(map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": "moon",
		},
	}, map[string]interface{}{
		"hello": "world",
	}))

	assert.NotNil(t, IsSubset(map[string]interface{}{
		"hello": "world",
	}, map[string]interface{}{}))

	assert.Nil(t, IsSubset(map[string]interface{}{
		"hello": []int{
			1, 2, 3,
		},
	}, map[string]interface{}{
		"hello": []int{
			1, 2, 3,
		},
	}))

	assert.Nil(t, IsSubset(map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": []map[string]interface{}{
				{
					"image": "hello",
				},
			},
		},
	}, map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": []map[string]interface{}{
				{
					"image": "hello",
					"bye":   "moon",
				},
			},
		},
	}))

	assert.Nil(t, IsSubset(map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": []map[string]interface{}{
				{
					"image": "hello",
				},
			},
		},
	}, map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": []map[string]interface{}{
				{
					"image": "hello",
					"bye":   "moon",
				},
				{
					"bye": "moon",
				},
			},
		},
	}))

	assert.Nil(t, IsSubset(map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": []map[string]interface{}{
				{
					"image": "hello",
				},
			},
		},
	}, map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": []map[string]interface{}{
				{
					"bye":   "moon",
					"image": "hello",
				},
				{
					"bye": "moon",
				},
			},
		},
	}))

	assert.NotNil(t, IsSubset(map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": []map[string]interface{}{
				{
					"image": "hello",
				},
			},
		},
	}, map[string]interface{}{
		"hello": map[string]interface{}{
			"hello": []map[string]interface{}{
				{
					"image": "world",
				},
			},
		},
	}))

	assert.Nil(t, IsSubset(map[string]interface{}{
		"containers": map[string]interface{}{
			"nginx": []map[string]interface{}{
				{
					"name":  "nginx-1",
					"image": "nginx:1.7.9",
				},
			},
		},
	}, map[string]interface{}{
		"containers": map[string]interface{}{
			"nginx": []map[string]interface{}{
				{
					"name":  "nginx-1",
					"image": "nginx:1.7.9",
				},
				{
					"name":  "nginx-2",
					"image": "nginx:1.7.9",
				},
			},
		},
	}))
	assert.Nil(t, IsSubset(map[string]interface{}{
		"containers": map[string]interface{}{
			"nginx": []map[string]interface{}{
				{
					"name":  "nginx-3",
					"image": "nginx:0.7.9",
				},
			},
		},
	}, map[string]interface{}{
		"containers": map[string]interface{}{
			"nginx": []map[string]interface{}{
				{
					"name":  "nginx-1",
					"image": "nginx:1.7.9",
				},
				{
					"name":  "nginx-2",
					"image": "nginx:1.7.9",
				},
				{
					"name":  "nginx-3",
					"image": "nginx:0.7.9",
				},
			},
		},
	}))
}
