# vkGetDeviceMemoryCommitment(3) Manual Page

## Name

vkGetDeviceMemoryCommitment - Query the current commitment for a VkDeviceMemory



## [](#_c_specification)C Specification

To determine the amount of lazily-allocated memory that is currently committed for a memory object, call:

```c++
// Provided by VK_VERSION_1_0
void vkGetDeviceMemoryCommitment(
    VkDevice                                    device,
    VkDeviceMemory                              memory,
    VkDeviceSize*                               pCommittedMemoryInBytes);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the memory.
- `memory` is the memory object being queried.
- `pCommittedMemoryInBytes` is a pointer to a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) value in which the number of bytes currently committed is returned, on success.

## [](#_description)Description

The implementation **may** update the commitment at any time, and the value returned by this query **may** be out of date.

The implementation guarantees to allocate any committed memory from the `heapIndex` indicated by the memory type that the memory object was created with.

Valid Usage

- [](#VUID-vkGetDeviceMemoryCommitment-memory-00690)VUID-vkGetDeviceMemoryCommitment-memory-00690  
  `memory` **must** have been created with a memory type that reports `VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT`

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceMemoryCommitment-device-parameter)VUID-vkGetDeviceMemoryCommitment-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceMemoryCommitment-memory-parameter)VUID-vkGetDeviceMemoryCommitment-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-vkGetDeviceMemoryCommitment-pCommittedMemoryInBytes-parameter)VUID-vkGetDeviceMemoryCommitment-pCommittedMemoryInBytes-parameter  
  `pCommittedMemoryInBytes` **must** be a valid pointer to a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) value
- [](#VUID-vkGetDeviceMemoryCommitment-memory-parent)VUID-vkGetDeviceMemoryCommitment-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceMemoryCommitment)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0