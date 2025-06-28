# VkQueueFlagBits(3) Manual Page

## Name

VkQueueFlagBits - Bitmask specifying capabilities of queues in a queue family



## [](#_c_specification)C Specification

Bits which **may** be set in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html)::`queueFlags`, indicating capabilities of queues in a queue family are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkQueueFlagBits {
    VK_QUEUE_GRAPHICS_BIT = 0x00000001,
    VK_QUEUE_COMPUTE_BIT = 0x00000002,
    VK_QUEUE_TRANSFER_BIT = 0x00000004,
    VK_QUEUE_SPARSE_BINDING_BIT = 0x00000008,
  // Provided by VK_VERSION_1_1
    VK_QUEUE_PROTECTED_BIT = 0x00000010,
  // Provided by VK_KHR_video_decode_queue
    VK_QUEUE_VIDEO_DECODE_BIT_KHR = 0x00000020,
  // Provided by VK_KHR_video_encode_queue
    VK_QUEUE_VIDEO_ENCODE_BIT_KHR = 0x00000040,
  // Provided by VK_NV_optical_flow
    VK_QUEUE_OPTICAL_FLOW_BIT_NV = 0x00000100,
  // Provided by VK_ARM_data_graph
    VK_QUEUE_DATA_GRAPH_BIT_ARM = 0x00000400,
} VkQueueFlagBits;
```

## [](#_description)Description

- `VK_QUEUE_GRAPHICS_BIT` specifies that queues in this queue family support graphics operations.
- `VK_QUEUE_COMPUTE_BIT` specifies that queues in this queue family support compute operations.
- `VK_QUEUE_TRANSFER_BIT` specifies that queues in this queue family support transfer operations.
- `VK_QUEUE_SPARSE_BINDING_BIT` specifies that queues in this queue family support sparse memory management operations (see [Sparse Resources](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#sparsememory)). If any of the sparse resource features are enabled, then at least one queue family **must** support this bit.
- `VK_QUEUE_VIDEO_DECODE_BIT_KHR` specifies that queues in this queue family support [video decode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-decode-operations).
- `VK_QUEUE_VIDEO_ENCODE_BIT_KHR` specifies that queues in this queue family support [video encode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-encode-operations).
- `VK_QUEUE_OPTICAL_FLOW_BIT_NV` specifies that queues in this queue family support optical flow operations.
- `VK_QUEUE_DATA_GRAPH_BIT_ARM` specifies that queues in this queue family support [data graph operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#graphs-operations).
- `VK_QUEUE_PROTECTED_BIT` specifies that queues in this queue family support the `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT` bit. (see [Protected Memory](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-protected-memory)). If the physical device supports the [`protectedMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-protectedMemory) feature, at least one of its queue families **must** support this bit.

If an implementation exposes any queue family that supports graphics operations, at least one queue family of at least one physical device exposed by the implementation **must** support both graphics and compute operations.

Furthermore, if the [`protectedMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-protectedMemory) physical device feature is supported, then at least one queue family of at least one physical device exposed by the implementation **must** support graphics operations, compute operations, and protected memory operations.

Note

All commands that are allowed on a queue that supports transfer operations are also allowed on a queue that supports either graphics or compute operations. Thus, if the capabilities of a queue family include `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`, then reporting the `VK_QUEUE_TRANSFER_BIT` capability separately for that queue family is **optional**.

For further details see [Queues](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-queues).

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkQueueFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueueFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0