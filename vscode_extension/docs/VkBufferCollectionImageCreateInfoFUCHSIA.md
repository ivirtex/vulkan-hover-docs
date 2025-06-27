# VkBufferCollectionImageCreateInfoFUCHSIA(3) Manual Page

## Name

VkBufferCollectionImageCreateInfoFUCHSIA - Create a
VkBufferCollectionFUCHSIA-compatible VkImage



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferCollectionImageCreateInfoFUCHSIA` structure is defined as:

``` c
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkBufferCollectionImageCreateInfoFUCHSIA {
    VkStructureType              sType;
    const void*                  pNext;
    VkBufferCollectionFUCHSIA    collection;
    uint32_t                     index;
} VkBufferCollectionImageCreateInfoFUCHSIA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this structure

- `collection` is the
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle

- `index` is the index of the buffer in the buffer collection from which
  the memory will be imported

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBufferCollectionImageCreateInfoFUCHSIA-index-06391"
  id="VUID-VkBufferCollectionImageCreateInfoFUCHSIA-index-06391"></a>
  VUID-VkBufferCollectionImageCreateInfoFUCHSIA-index-06391  
  `index` **must** be less than
  [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html)::`bufferCount`

Valid Usage (Implicit)

- <a href="#VUID-VkBufferCollectionImageCreateInfoFUCHSIA-sType-sType"
  id="VUID-VkBufferCollectionImageCreateInfoFUCHSIA-sType-sType"></a>
  VUID-VkBufferCollectionImageCreateInfoFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_IMAGE_CREATE_INFO_FUCHSIA`

- <a
  href="#VUID-VkBufferCollectionImageCreateInfoFUCHSIA-collection-parameter"
  id="VUID-VkBufferCollectionImageCreateInfoFUCHSIA-collection-parameter"></a>
  VUID-VkBufferCollectionImageCreateInfoFUCHSIA-collection-parameter  
  `collection` **must** be a valid
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferCollectionImageCreateInfoFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
