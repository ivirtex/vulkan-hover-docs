# vkCreateBufferCollectionFUCHSIA(3) Manual Page

## Name

vkCreateBufferCollectionFUCHSIA - Create a new buffer collection



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html)
for Vulkan to participate in the buffer collection:

``` c
// Provided by VK_FUCHSIA_buffer_collection
VkResult vkCreateBufferCollectionFUCHSIA(
    VkDevice                                    device,
    const VkBufferCollectionCreateInfoFUCHSIA*  pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkBufferCollectionFUCHSIA*                  pCollection);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the
  `VkBufferCollectionFUCHSIA`

- `pCreateInfo` is a pointer to a
  [VkBufferCollectionCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionCreateInfoFUCHSIA.html)
  structure containing parameters affecting creation of the buffer
  collection

- `pAllocator` is a pointer to a
  [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html) structure
  controlling host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter

- `pBufferCollection` is a pointer to a
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle in
  which the resulting buffer collection object is returned

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateBufferCollectionFUCHSIA-device-parameter"
  id="VUID-vkCreateBufferCollectionFUCHSIA-device-parameter"></a>
  VUID-vkCreateBufferCollectionFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateBufferCollectionFUCHSIA-pCreateInfo-parameter"
  id="VUID-vkCreateBufferCollectionFUCHSIA-pCreateInfo-parameter"></a>
  VUID-vkCreateBufferCollectionFUCHSIA-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkBufferCollectionCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionCreateInfoFUCHSIA.html)
  structure

- <a href="#VUID-vkCreateBufferCollectionFUCHSIA-pAllocator-parameter"
  id="VUID-vkCreateBufferCollectionFUCHSIA-pAllocator-parameter"></a>
  VUID-vkCreateBufferCollectionFUCHSIA-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateBufferCollectionFUCHSIA-pCollection-parameter"
  id="VUID-vkCreateBufferCollectionFUCHSIA-pCollection-parameter"></a>
  VUID-vkCreateBufferCollectionFUCHSIA-pCollection-parameter  
  `pCollection` **must** be a valid pointer to a
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

- `VK_ERROR_INITIALIZATION_FAILED`

Host Access

All functions referencing a
[VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) **must** be
externally synchronized with the exception of
`vkCreateBufferCollectionFUCHSIA`.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkBufferCollectionCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionCreateInfoFUCHSIA.html),
[VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateBufferCollectionFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
