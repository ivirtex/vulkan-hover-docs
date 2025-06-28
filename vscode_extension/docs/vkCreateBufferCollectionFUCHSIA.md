# vkCreateBufferCollectionFUCHSIA(3) Manual Page

## Name

vkCreateBufferCollectionFUCHSIA - Create a new buffer collection



## [](#_c_specification)C Specification

To create a [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) for Vulkan to participate in the buffer collection:

```c++
// Provided by VK_FUCHSIA_buffer_collection
VkResult vkCreateBufferCollectionFUCHSIA(
    VkDevice                                    device,
    const VkBufferCollectionCreateInfoFUCHSIA*  pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkBufferCollectionFUCHSIA*                  pCollection);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the `VkBufferCollectionFUCHSIA`
- `pCreateInfo` is a pointer to a [VkBufferCollectionCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionCreateInfoFUCHSIA.html) structure containing parameters affecting creation of the buffer collection
- `pAllocator` is a pointer to a [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure controlling host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter
- `pCollection` is a pointer to a [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle in which the resulting buffer collection object is returned

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateBufferCollectionFUCHSIA-device-parameter)VUID-vkCreateBufferCollectionFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateBufferCollectionFUCHSIA-pCreateInfo-parameter)VUID-vkCreateBufferCollectionFUCHSIA-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkBufferCollectionCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionCreateInfoFUCHSIA.html) structure
- [](#VUID-vkCreateBufferCollectionFUCHSIA-pAllocator-parameter)VUID-vkCreateBufferCollectionFUCHSIA-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateBufferCollectionFUCHSIA-pCollection-parameter)VUID-vkCreateBufferCollectionFUCHSIA-pCollection-parameter  
  `pCollection` **must** be a valid pointer to a [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle
- [](#VUID-vkCreateBufferCollectionFUCHSIA-device-queuecount)VUID-vkCreateBufferCollectionFUCHSIA-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_INVALID_EXTERNAL_HANDLE`
- `VK_ERROR_INITIALIZATION_FAILED`

Host Access

All functions referencing a [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) **must** be externally synchronized with the exception of `vkCreateBufferCollectionFUCHSIA`.

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkBufferCollectionCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionCreateInfoFUCHSIA.html), [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateBufferCollectionFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0