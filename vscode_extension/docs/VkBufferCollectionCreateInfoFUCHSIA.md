# VkBufferCollectionCreateInfoFUCHSIA(3) Manual Page

## Name

VkBufferCollectionCreateInfoFUCHSIA - Structure specifying desired parameters to create the buffer collection



## [](#_c_specification)C Specification

The `VkBufferCollectionCreateInfoFUCHSIA` structure is defined as:

```c++
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkBufferCollectionCreateInfoFUCHSIA {
    VkStructureType    sType;
    const void*        pNext;
    zx_handle_t        collectionToken;
} VkBufferCollectionCreateInfoFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure
- `collectionToken` is a `zx_handle_t` containing the Sysmem client’s buffer collection token

## [](#_description)Description

Valid Usage

- [](#VUID-VkBufferCollectionCreateInfoFUCHSIA-collectionToken-06393)VUID-VkBufferCollectionCreateInfoFUCHSIA-collectionToken-06393  
  `collectionToken` **must** be a valid `zx_handle_t` to a Zircon channel allocated from Sysmem (`fuchsia.sysmem.Allocator`/AllocateSharedCollection) with `ZX_DEFAULT_CHANNEL_RIGHTS` rights

Valid Usage (Implicit)

- [](#VUID-VkBufferCollectionCreateInfoFUCHSIA-sType-sType)VUID-VkBufferCollectionCreateInfoFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_CREATE_INFO_FUCHSIA`
- [](#VUID-VkBufferCollectionCreateInfoFUCHSIA-pNext-pNext)VUID-VkBufferCollectionCreateInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateBufferCollectionFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferCollectionCreateInfoFUCHSIA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0