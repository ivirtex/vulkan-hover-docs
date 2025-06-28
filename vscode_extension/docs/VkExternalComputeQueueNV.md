# VkExternalComputeQueueNV(3) Manual Page

## Name

VkExternalComputeQueueNV - Opaque handle to an external compute queue



## [](#_c_specification)C Specification

External compute queues are used to join compatible external APIs to a `VkDevice`, allowing workloads submitted through these external APIs to be executed simultaneously to workloads submitted through Vulkan.

External compute queues are represented by `VkExternalComputeQueueNV` handles:

```c++
// Provided by VK_NV_external_compute_queue
VK_DEFINE_HANDLE(VkExternalComputeQueueNV)
```

## [](#_see_also)See Also

[VK\_DEFINE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_HANDLE.html), [VK\_NV\_external\_compute\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_compute_queue.html), [vkCreateExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateExternalComputeQueueNV.html), [vkDestroyExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyExternalComputeQueueNV.html), [vkGetExternalComputeQueueDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExternalComputeQueueDataNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalComputeQueueNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0