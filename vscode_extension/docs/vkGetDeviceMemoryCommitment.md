# vkGetDeviceMemoryCommitment(3) Manual Page

## Name

vkGetDeviceMemoryCommitment - Query the current commitment for a
VkDeviceMemory



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the amount of lazily-allocated memory that is currently
committed for a memory object, call:

``` c
// Provided by VK_VERSION_1_0
void vkGetDeviceMemoryCommitment(
    VkDevice                                    device,
    VkDeviceMemory                              memory,
    VkDeviceSize*                               pCommittedMemoryInBytes);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory.

- `memory` is the memory object being queried.

- `pCommittedMemoryInBytes` is a pointer to a
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) value in which the number of bytes
  currently committed is returned, on success.

## <a href="#_description" class="anchor"></a>Description

The implementation **may** update the commitment at any time, and the
value returned by this query **may** be out of date.

The implementation guarantees to allocate any committed memory from the
`heapIndex` indicated by the memory type that the memory object was
created with.

Valid Usage

- <a href="#VUID-vkGetDeviceMemoryCommitment-memory-00690"
  id="VUID-vkGetDeviceMemoryCommitment-memory-00690"></a>
  VUID-vkGetDeviceMemoryCommitment-memory-00690  
  `memory` **must** have been created with a memory type that reports
  `VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT`

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceMemoryCommitment-device-parameter"
  id="VUID-vkGetDeviceMemoryCommitment-device-parameter"></a>
  VUID-vkGetDeviceMemoryCommitment-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDeviceMemoryCommitment-memory-parameter"
  id="VUID-vkGetDeviceMemoryCommitment-memory-parameter"></a>
  VUID-vkGetDeviceMemoryCommitment-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a
  href="#VUID-vkGetDeviceMemoryCommitment-pCommittedMemoryInBytes-parameter"
  id="VUID-vkGetDeviceMemoryCommitment-pCommittedMemoryInBytes-parameter"></a>
  VUID-vkGetDeviceMemoryCommitment-pCommittedMemoryInBytes-parameter  
  `pCommittedMemoryInBytes` **must** be a valid pointer to a
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) value

- <a href="#VUID-vkGetDeviceMemoryCommitment-memory-parent"
  id="VUID-vkGetDeviceMemoryCommitment-memory-parent"></a>
  VUID-vkGetDeviceMemoryCommitment-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceMemoryCommitment"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
