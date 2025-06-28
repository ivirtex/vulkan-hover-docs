# vkDestroyBufferCollectionFUCHSIA(3) Manual Page

## Name

vkDestroyBufferCollectionFUCHSIA - Destroy a buffer collection



## [](#_c_specification)C Specification

To release a [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html):

```c++
// Provided by VK_FUCHSIA_buffer_collection
void vkDestroyBufferCollectionFUCHSIA(
    VkDevice                                    device,
    VkBufferCollectionFUCHSIA                   collection,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the `VkBufferCollectionFUCHSIA`
- `collection` is the [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle
- `pAllocator` is a pointer to a [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure controlling host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyBufferCollectionFUCHSIA-collection-06407)VUID-vkDestroyBufferCollectionFUCHSIA-collection-06407  
  [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) and [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) objects that referenced `collection` upon creation by inclusion of a [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html) or [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html) chained to their [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) or [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) structures respectively, **may** outlive `collection`

Valid Usage (Implicit)

- [](#VUID-vkDestroyBufferCollectionFUCHSIA-device-parameter)VUID-vkDestroyBufferCollectionFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyBufferCollectionFUCHSIA-collection-parameter)VUID-vkDestroyBufferCollectionFUCHSIA-collection-parameter  
  `collection` **must** be a valid [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle
- [](#VUID-vkDestroyBufferCollectionFUCHSIA-pAllocator-parameter)VUID-vkDestroyBufferCollectionFUCHSIA-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyBufferCollectionFUCHSIA-collection-parent)VUID-vkDestroyBufferCollectionFUCHSIA-collection-parent  
  `collection` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyBufferCollectionFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0