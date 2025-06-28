# vkDestroyQueryPool(3) Manual Page

## Name

vkDestroyQueryPool - Destroy a query pool object



## [](#_c_specification)C Specification

To destroy a query pool, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroyQueryPool(
    VkDevice                                    device,
    VkQueryPool                                 queryPool,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the query pool.
- `queryPool` is the query pool to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyQueryPool-queryPool-00793)VUID-vkDestroyQueryPool-queryPool-00793  
  All submitted commands that refer to `queryPool` **must** have completed execution
- [](#VUID-vkDestroyQueryPool-queryPool-00794)VUID-vkDestroyQueryPool-queryPool-00794  
  If `VkAllocationCallbacks` were provided when `queryPool` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyQueryPool-queryPool-00795)VUID-vkDestroyQueryPool-queryPool-00795  
  If no `VkAllocationCallbacks` were provided when `queryPool` was created, `pAllocator` **must** be `NULL`

Note

Applications **can** verify that `queryPool` **can** be destroyed by checking that `vkGetQueryPoolResults`() without the `VK_QUERY_RESULT_PARTIAL_BIT` flag returns `VK_SUCCESS` for all queries that are used in command buffers submitted for execution.

Valid Usage (Implicit)

- [](#VUID-vkDestroyQueryPool-device-parameter)VUID-vkDestroyQueryPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyQueryPool-queryPool-parameter)VUID-vkDestroyQueryPool-queryPool-parameter  
  If `queryPool` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkDestroyQueryPool-pAllocator-parameter)VUID-vkDestroyQueryPool-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyQueryPool-queryPool-parent)VUID-vkDestroyQueryPool-queryPool-parent  
  If `queryPool` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `queryPool` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyQueryPool)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0