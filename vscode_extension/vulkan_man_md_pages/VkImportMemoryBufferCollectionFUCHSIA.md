# VkImportMemoryBufferCollectionFUCHSIA(3) Manual Page

## Name

VkImportMemoryBufferCollectionFUCHSIA - Structure to specify the Sysmem
buffer to import



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImportMemoryBufferCollectionFUCHSIA` structure is defined as:

``` c
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkImportMemoryBufferCollectionFUCHSIA {
    VkStructureType              sType;
    const void*                  pNext;
    VkBufferCollectionFUCHSIA    collection;
    uint32_t                     index;
} VkImportMemoryBufferCollectionFUCHSIA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this structure

- `collection` is the
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle

- `index` the index of the buffer to import from `collection`

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImportMemoryBufferCollectionFUCHSIA-index-06406"
  id="VUID-VkImportMemoryBufferCollectionFUCHSIA-index-06406"></a>
  VUID-VkImportMemoryBufferCollectionFUCHSIA-index-06406  
  `index` **must** be less than the value retrieved as
  [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html):bufferCount

Valid Usage (Implicit)

- <a href="#VUID-VkImportMemoryBufferCollectionFUCHSIA-sType-sType"
  id="VUID-VkImportMemoryBufferCollectionFUCHSIA-sType-sType"></a>
  VUID-VkImportMemoryBufferCollectionFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMPORT_MEMORY_BUFFER_COLLECTION_FUCHSIA`

- <a
  href="#VUID-VkImportMemoryBufferCollectionFUCHSIA-collection-parameter"
  id="VUID-VkImportMemoryBufferCollectionFUCHSIA-collection-parameter"></a>
  VUID-VkImportMemoryBufferCollectionFUCHSIA-collection-parameter  
  `collection` **must** be a valid
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportMemoryBufferCollectionFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
