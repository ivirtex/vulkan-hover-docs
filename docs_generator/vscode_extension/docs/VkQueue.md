# VkQueue(3) Manual Page

## Name

VkQueue - Opaque handle to a queue object



## [](#_c_specification)C Specification

Creating a logical device also creates the queues associated with that device. The queues to create are described by a set of [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html) structures that are passed to [vkCreateDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDevice.html) in `pQueueCreateInfos`. Queues **cannot** be independently destroyed, and are instead destroyed with the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) that they were created from.

Queues are represented by `VkQueue` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_HANDLE(VkQueue)
```

## [](#_see_also)See Also

[VK\_DEFINE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkExportMetalCommandQueueInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalCommandQueueInfoEXT.html), [VkExternalComputeQueueCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueCreateInfoNV.html), [vkGetDeviceQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceQueue.html), [vkGetDeviceQueue2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceQueue2.html), [vkGetQueueCheckpointData2NV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetQueueCheckpointData2NV.html), [vkGetQueueCheckpointDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetQueueCheckpointDataNV.html), [vkQueueBeginDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueBeginDebugUtilsLabelEXT.html), [vkQueueBindSparse](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueBindSparse.html), [vkQueueEndDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueEndDebugUtilsLabelEXT.html), [vkQueueInsertDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueInsertDebugUtilsLabelEXT.html), [vkQueueNotifyOutOfBandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueNotifyOutOfBandNV.html), [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html), [vkQueueSetPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSetPerformanceConfigurationINTEL.html), [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit.html), [vkQueueSubmit2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit2.html), [vkQueueSubmit2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit2KHR.html), [vkQueueWaitIdle](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueWaitIdle.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueue)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0