# vkDestroyCommandPool(3) Manual Page

## Name

vkDestroyCommandPool - Destroy a command pool object



## [](#_c_specification)C Specification

To destroy a command pool, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroyCommandPool(
    VkDevice                                    device,
    VkCommandPool                               commandPool,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the command pool.
- `commandPool` is the handle of the command pool to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

When a pool is destroyed, all command buffers allocated from the pool are [freed](https://registry.khronos.org/vulkan/specs/latest/man/html/vkFreeCommandBuffers.html).

Any primary command buffer allocated from another [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) that is in the [recording or executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle) and has a secondary command buffer allocated from `commandPool` recorded into it, becomes [invalid](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).

Valid Usage

- [](#VUID-vkDestroyCommandPool-commandPool-00041)VUID-vkDestroyCommandPool-commandPool-00041  
  All `VkCommandBuffer` objects allocated from `commandPool` **must** not be in the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkDestroyCommandPool-commandPool-00042)VUID-vkDestroyCommandPool-commandPool-00042  
  If `VkAllocationCallbacks` were provided when `commandPool` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyCommandPool-commandPool-00043)VUID-vkDestroyCommandPool-commandPool-00043  
  If no `VkAllocationCallbacks` were provided when `commandPool` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyCommandPool-device-parameter)VUID-vkDestroyCommandPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyCommandPool-commandPool-parameter)VUID-vkDestroyCommandPool-commandPool-parameter  
  If `commandPool` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `commandPool` **must** be a valid [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) handle
- [](#VUID-vkDestroyCommandPool-pAllocator-parameter)VUID-vkDestroyCommandPool-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyCommandPool-commandPool-parent)VUID-vkDestroyCommandPool-commandPool-parent  
  If `commandPool` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `commandPool` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyCommandPool)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0