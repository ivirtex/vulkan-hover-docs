# VkSysmemColorSpaceFUCHSIA(3) Manual Page

## Name

VkSysmemColorSpaceFUCHSIA - Structure describing the buffer collections color space



## [](#_c_specification)C Specification

The `VkSysmemColorSpaceFUCHSIA` structure is defined as:

```c++
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkSysmemColorSpaceFUCHSIA {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           colorSpace;
} VkSysmemColorSpaceFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure
- `colorSpace` value of the Sysmem `ColorSpaceType`

## [](#_description)Description

Valid Usage

- [](#VUID-VkSysmemColorSpaceFUCHSIA-colorSpace-06402)VUID-VkSysmemColorSpaceFUCHSIA-colorSpace-06402  
  `colorSpace` **must** be a `ColorSpaceType` as defined in `fuchsia.sysmem/image_formats.fidl`

Valid Usage (Implicit)

- [](#VUID-VkSysmemColorSpaceFUCHSIA-sType-sType)VUID-VkSysmemColorSpaceFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SYSMEM_COLOR_SPACE_FUCHSIA`
- [](#VUID-VkSysmemColorSpaceFUCHSIA-pNext-pNext)VUID-VkSysmemColorSpaceFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionPropertiesFUCHSIA.html), [VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatConstraintsInfoFUCHSIA.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSysmemColorSpaceFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0