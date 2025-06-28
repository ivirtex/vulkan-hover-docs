# VkTensorUsageFlagBitsARM(3) Manual Page

## Name

VkTensorUsageFlagBitsARM - Bitmask specifying allowed usage of a tensor



## [](#_c_specification)C Specification

Bits which **can** be set in [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`usage`, specifying usage behavior of a tensor, are:

```c++
// Provided by VK_ARM_tensors
// Flag bits for VkTensorUsageFlagBitsARM
typedef VkFlags64 VkTensorUsageFlagBitsARM;
static const VkTensorUsageFlagBitsARM VK_TENSOR_USAGE_SHADER_BIT_ARM = 0x00000002ULL;
static const VkTensorUsageFlagBitsARM VK_TENSOR_USAGE_TRANSFER_SRC_BIT_ARM = 0x00000004ULL;
static const VkTensorUsageFlagBitsARM VK_TENSOR_USAGE_TRANSFER_DST_BIT_ARM = 0x00000008ULL;
static const VkTensorUsageFlagBitsARM VK_TENSOR_USAGE_IMAGE_ALIASING_BIT_ARM = 0x00000010ULL;
// Provided by VK_ARM_data_graph
static const VkTensorUsageFlagBitsARM VK_TENSOR_USAGE_DATA_GRAPH_BIT_ARM = 0x00000020ULL;
```

## [](#_description)Description

- `VK_TENSOR_USAGE_SHADER_BIT_ARM` specifies that the tensor **can** be used to create a `VkTensorViewARM` suitable for occupying a `VkDescriptorSet` slot of type `VK_DESCRIPTOR_TYPE_TENSOR_ARM` accessed by shader stages.
- `VK_TENSOR_USAGE_TRANSFER_SRC_BIT_ARM` specifies that the tensor **can** be used as the source of a *transfer command* (see the definition of [`VK_PIPELINE_STAGE_TRANSFER_BIT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-transfer)).
- `VK_TENSOR_USAGE_TRANSFER_DST_BIT_ARM` specifies that the tensor **can** be used as the destination of a transfer command.
- `VK_TENSOR_USAGE_IMAGE_ALIASING_BIT_ARM` specifies that the tensor **can** be bound to a range of memory aliased with an image created with `VK_IMAGE_TILING_OPTIMAL`. See [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-memory-aliasing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-memory-aliasing) for a complete set of rules for tensor/image aliasing.
- `VK_TENSOR_USAGE_DATA_GRAPH_BIT_ARM` specifies that the tensor **can** be used to create a `VkTensorViewARM` suitable for occupying a `VkDescriptorSet` slot of type `VK_DESCRIPTOR_TYPE_TENSOR_ARM` accessed by [data graph pipelines](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#graphs-pipelines).

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkTensorUsageFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorUsageFlagsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorUsageFlagBitsARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0