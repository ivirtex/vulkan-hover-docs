# vkDestroyBufferCollectionFUCHSIA(3) Manual Page

## Name

vkDestroyBufferCollectionFUCHSIA - Destroy a buffer collection



## <a href="#_c_specification" class="anchor"></a>C Specification

To release a
[VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html):

``` c
// Provided by VK_FUCHSIA_buffer_collection
void vkDestroyBufferCollectionFUCHSIA(
    VkDevice                                    device,
    VkBufferCollectionFUCHSIA                   collection,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the
  `VkBufferCollectionFUCHSIA`

- `collection` is the
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle

- `pAllocator` is a pointer to a
  [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html) structure
  controlling host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyBufferCollectionFUCHSIA-collection-06407"
  id="VUID-vkDestroyBufferCollectionFUCHSIA-collection-06407"></a>
  VUID-vkDestroyBufferCollectionFUCHSIA-collection-06407  
  [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) and [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) objects that
  referenced `collection` upon creation by inclusion of a
  [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)
  or
  [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html)
  chained to their [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) or
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) structures respectively,
  **may** outlive `collection`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyBufferCollectionFUCHSIA-device-parameter"
  id="VUID-vkDestroyBufferCollectionFUCHSIA-device-parameter"></a>
  VUID-vkDestroyBufferCollectionFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyBufferCollectionFUCHSIA-collection-parameter"
  id="VUID-vkDestroyBufferCollectionFUCHSIA-collection-parameter"></a>
  VUID-vkDestroyBufferCollectionFUCHSIA-collection-parameter  
  `collection` **must** be a valid
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle

- <a href="#VUID-vkDestroyBufferCollectionFUCHSIA-pAllocator-parameter"
  id="VUID-vkDestroyBufferCollectionFUCHSIA-pAllocator-parameter"></a>
  VUID-vkDestroyBufferCollectionFUCHSIA-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyBufferCollectionFUCHSIA-collection-parent"
  id="VUID-vkDestroyBufferCollectionFUCHSIA-collection-parent"></a>
  VUID-vkDestroyBufferCollectionFUCHSIA-collection-parent  
  `collection` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyBufferCollectionFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
