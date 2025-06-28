# VkCommandPool(3) Manual Page

## Name

VkCommandPool - Opaque handle to a command pool object



## [](#_c_specification)C Specification

Command pools are opaque objects that command buffer memory is allocated from, and which allow the implementation to amortize the cost of resource creation across multiple command buffers. Command pools are externally synchronized, meaning that a command pool **must** not be used concurrently in multiple threads. That includes use via recording commands on any command buffers allocated from the pool, as well as operations that allocate, free, and reset command buffers or the pool itself.

Command pools are represented by `VkCommandPool` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkCommandPool)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBufferAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferAllocateInfo.html), [vkCreateCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCommandPool.html), [vkDestroyCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyCommandPool.html), [vkFreeCommandBuffers](https://registry.khronos.org/vulkan/specs/latest/man/html/vkFreeCommandBuffers.html), [vkResetCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetCommandPool.html), [vkTrimCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkTrimCommandPool.html), [vkTrimCommandPoolKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkTrimCommandPoolKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCommandPool)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0