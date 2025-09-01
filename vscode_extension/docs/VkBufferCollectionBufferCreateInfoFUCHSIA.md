# VkBufferCollectionBufferCreateInfoFUCHSIA(3) Manual Page

## Name

VkBufferCollectionBufferCreateInfoFUCHSIA - Create a VkBufferCollectionFUCHSIA-compatible VkBuffer



## [](#_c_specification)C Specification

The `VkBufferCollectionBufferCreateInfoFUCHSIA` structure is defined as:

```c++
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkBufferCollectionBufferCreateInfoFUCHSIA {
    VkStructureType              sType;
    const void*                  pNext;
    VkBufferCollectionFUCHSIA    collection;
    uint32_t                     index;
} VkBufferCollectionBufferCreateInfoFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure
- `collection` is the [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle
- `index` is the index of the buffer in the buffer collection from which the memory will be imported

## [](#_description)Description

Valid Usage

- [](#VUID-VkBufferCollectionBufferCreateInfoFUCHSIA-index-06388)VUID-VkBufferCollectionBufferCreateInfoFUCHSIA-index-06388  
  `index` **must** be less than [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionPropertiesFUCHSIA.html)::`bufferCount`

Valid Usage (Implicit)

- [](#VUID-VkBufferCollectionBufferCreateInfoFUCHSIA-sType-sType)VUID-VkBufferCollectionBufferCreateInfoFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_BUFFER_CREATE_INFO_FUCHSIA`
- [](#VUID-VkBufferCollectionBufferCreateInfoFUCHSIA-collection-parameter)VUID-VkBufferCollectionBufferCreateInfoFUCHSIA-collection-parameter  
  `collection` **must** be a valid [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferCollectionBufferCreateInfoFUCHSIA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0