# VkBufferCollectionCreateInfoFUCHSIA(3) Manual Page

## Name

VkBufferCollectionCreateInfoFUCHSIA - Structure specifying desired
parameters to create the buffer collection



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferCollectionCreateInfoFUCHSIA` structure is defined as:

``` c
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkBufferCollectionCreateInfoFUCHSIA {
    VkStructureType    sType;
    const void*        pNext;
    zx_handle_t        collectionToken;
} VkBufferCollectionCreateInfoFUCHSIA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this structure

- `collectionToken` is a `zx_handle_t` containing the Sysmem client’s
  buffer collection token

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkBufferCollectionCreateInfoFUCHSIA-collectionToken-06393"
  id="VUID-VkBufferCollectionCreateInfoFUCHSIA-collectionToken-06393"></a>
  VUID-VkBufferCollectionCreateInfoFUCHSIA-collectionToken-06393  
  `collectionToken` **must** be a valid `zx_handle_t` to a Zircon
  channel allocated from Sysmem
  (`fuchsia.sysmem.Allocator`/AllocateSharedCollection) with
  `ZX_DEFAULT_CHANNEL_RIGHTS` rights

Valid Usage (Implicit)

- <a href="#VUID-VkBufferCollectionCreateInfoFUCHSIA-sType-sType"
  id="VUID-VkBufferCollectionCreateInfoFUCHSIA-sType-sType"></a>
  VUID-VkBufferCollectionCreateInfoFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_CREATE_INFO_FUCHSIA`

- <a href="#VUID-VkBufferCollectionCreateInfoFUCHSIA-pNext-pNext"
  id="VUID-VkBufferCollectionCreateInfoFUCHSIA-pNext-pNext"></a>
  VUID-VkBufferCollectionCreateInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateBufferCollectionFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferCollectionCreateInfoFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
