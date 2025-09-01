# VkTensorMemoryBarrierARM(3) Manual Page

## Name

VkTensorMemoryBarrierARM - Structure specifying a tensor memory barrier



## [](#_c_specification)C Specification

The `VkTensorMemoryBarrierARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkTensorMemoryBarrierARM {
    VkStructureType          sType;
    const void*              pNext;
    VkPipelineStageFlags2    srcStageMask;
    VkAccessFlags2           srcAccessMask;
    VkPipelineStageFlags2    dstStageMask;
    VkAccessFlags2           dstAccessMask;
    uint32_t                 srcQueueFamilyIndex;
    uint32_t                 dstQueueFamilyIndex;
    VkTensorARM              tensor;
} VkTensorMemoryBarrierARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcStageMask` is a [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html) mask of pipeline stages to be included in the [first synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes).
- `srcAccessMask` is a [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html) mask of access flags to be included in the [first access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes).
- `dstStageMask` is a [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html) mask of pipeline stages to be included in the [second synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes).
- `dstAccessMask` is a [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html) mask of access flags to be included in the [second access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes).
- `srcQueueFamilyIndex` is the source queue family for a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers).
- `dstQueueFamilyIndex` is the destination queue family for a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers).
- `tensor` is a handle to the tensor whose backing memory is affected by the barrier.

## [](#_description)Description

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) and [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) described by this structure include only operations and memory accesses specified by `srcStageMask` and `srcAccessMask`.

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) and [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) described by this structure include only operations and memory accesses specified by `dstStageMask` and `dstAccessMask`.

Both [access scopes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) are limited to only memory accesses to `tensor`.

If `tensor` was created with `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` is not equal to `dstQueueFamilyIndex`, this memory barrier defines a [queue family transfer operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers). When executed on a queue in the family identified by `srcQueueFamilyIndex`, this barrier defines a [queue family release operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-release) for the specified tensor, and the second synchronization and access scopes do not synchronize operations on that queue. When executed on a queue in the family identified by `dstQueueFamilyIndex`, this barrier defines a [queue family acquire operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire) for the specified tensor, and the first synchronization and access scopes do not synchronize operations on that queue.

A [queue family transfer operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers) is also defined if the values are not equal, and either is one of the special queue family values reserved for external memory ownership transfers, as described in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers). A [queue family release operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-release) is defined when `dstQueueFamilyIndex` is one of those values, and a [queue family acquire operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire) is defined when `srcQueueFamilyIndex` is one of those values.

Valid Usage

- [](#VUID-VkTensorMemoryBarrierARM-tensor-09755)VUID-VkTensorMemoryBarrierARM-tensor-09755  
  If `tensor` was created with a sharing mode of `VK_SHARING_MODE_CONCURRENT`, `srcQueueFamilyIndex` and `dstQueueFamilyIndex` **must** both be `VK_QUEUE_FAMILY_IGNORED`
- [](#VUID-VkTensorMemoryBarrierARM-tensor-09756)VUID-VkTensorMemoryBarrierARM-tensor-09756  
  If `tensor` was created with a sharing mode of `VK_SHARING_MODE_EXCLUSIVE`, `srcQueueFamilyIndex` and `dstQueueFamilyIndex` **must** both be either `VK_QUEUE_FAMILY_IGNORED`, or a valid queue family (see [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-queueprops](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-queueprops))
- [](#VUID-VkTensorMemoryBarrierARM-tensor-09757)VUID-VkTensorMemoryBarrierARM-tensor-09757  
  If `tensor` was created with a sharing mode of `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and `dstQueueFamilyIndex` are not `VK_QUEUE_FAMILY_IGNORED`, at least one of them **must** be the same as the family of the queue that will execute this barrier
- [](#VUID-VkTensorMemoryBarrierARM-tensor-09758)VUID-VkTensorMemoryBarrierARM-tensor-09758  
  If `tensor` is non-sparse then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object

Valid Usage (Implicit)

- [](#VUID-VkTensorMemoryBarrierARM-sType-sType)VUID-VkTensorMemoryBarrierARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TENSOR_MEMORY_BARRIER_ARM`
- [](#VUID-VkTensorMemoryBarrierARM-srcStageMask-parameter)VUID-VkTensorMemoryBarrierARM-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html) values
- [](#VUID-VkTensorMemoryBarrierARM-srcAccessMask-parameter)VUID-VkTensorMemoryBarrierARM-srcAccessMask-parameter  
  `srcAccessMask` **must** be a valid combination of [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html) values
- [](#VUID-VkTensorMemoryBarrierARM-dstStageMask-parameter)VUID-VkTensorMemoryBarrierARM-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html) values
- [](#VUID-VkTensorMemoryBarrierARM-dstAccessMask-parameter)VUID-VkTensorMemoryBarrierARM-dstAccessMask-parameter  
  `dstAccessMask` **must** be a valid combination of [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html) values
- [](#VUID-VkTensorMemoryBarrierARM-tensor-parameter)VUID-VkTensorMemoryBarrierARM-tensor-parameter  
  `tensor` **must** be a valid [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handle

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html), [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html), [VkTensorDependencyInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDependencyInfoARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorMemoryBarrierARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0