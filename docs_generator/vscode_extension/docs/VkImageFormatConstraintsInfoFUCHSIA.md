# VkImageFormatConstraintsInfoFUCHSIA(3) Manual Page

## Name

VkImageFormatConstraintsInfoFUCHSIA - Structure image-based buffer collection constraints



## [](#_c_specification)C Specification

The `VkImageFormatConstraintsInfoFUCHSIA` structure is defined as:

```c++
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkImageFormatConstraintsInfoFUCHSIA {
    VkStructureType                         sType;
    const void*                             pNext;
    VkImageCreateInfo                       imageCreateInfo;
    VkFormatFeatureFlags                    requiredFormatFeatures;
    VkImageFormatConstraintsFlagsFUCHSIA    flags;
    uint64_t                                sysmemPixelFormat;
    uint32_t                                colorSpaceCount;
    const VkSysmemColorSpaceFUCHSIA*        pColorSpaces;
} VkImageFormatConstraintsInfoFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure
- `imageCreateInfo` is the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) used to create a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) that is to use memory from the [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html)
- `requiredFormatFeatures` is a bitmask of `VkFormatFeatureFlagBits` specifying required features of the buffers in the buffer collection
- `flags` is reserved for future use
- `sysmemPixelFormat` is a `PixelFormatType` value from the `fuchsia.sysmem/image_formats.fidl` FIDL interface
- `colorSpaceCount` the element count of `pColorSpaces`
- `pColorSpaces` is a pointer to an array of [VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSysmemColorSpaceFUCHSIA.html) structs of size `colorSpaceCount`

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkImageFormatConstraintsInfoFUCHSIA-sType-sType)VUID-VkImageFormatConstraintsInfoFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_FORMAT_CONSTRAINTS_INFO_FUCHSIA`
- [](#VUID-VkImageFormatConstraintsInfoFUCHSIA-pNext-pNext)VUID-VkImageFormatConstraintsInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImageFormatConstraintsInfoFUCHSIA-imageCreateInfo-parameter)VUID-VkImageFormatConstraintsInfoFUCHSIA-imageCreateInfo-parameter  
  `imageCreateInfo` **must** be a valid [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure
- [](#VUID-VkImageFormatConstraintsInfoFUCHSIA-requiredFormatFeatures-parameter)VUID-VkImageFormatConstraintsInfoFUCHSIA-requiredFormatFeatures-parameter  
  `requiredFormatFeatures` **must** be a valid combination of [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html) values
- [](#VUID-VkImageFormatConstraintsInfoFUCHSIA-requiredFormatFeatures-requiredbitmask)VUID-VkImageFormatConstraintsInfoFUCHSIA-requiredFormatFeatures-requiredbitmask  
  `requiredFormatFeatures` **must** not be `0`
- [](#VUID-VkImageFormatConstraintsInfoFUCHSIA-flags-zerobitmask)VUID-VkImageFormatConstraintsInfoFUCHSIA-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkImageFormatConstraintsInfoFUCHSIA-pColorSpaces-parameter)VUID-VkImageFormatConstraintsInfoFUCHSIA-pColorSpaces-parameter  
  `pColorSpaces` **must** be a valid pointer to an array of `colorSpaceCount` valid [VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSysmemColorSpaceFUCHSIA.html) structures
- [](#VUID-VkImageFormatConstraintsInfoFUCHSIA-colorSpaceCount-arraylength)VUID-VkImageFormatConstraintsInfoFUCHSIA-colorSpaceCount-arraylength  
  `colorSpaceCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags.html), [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFUCHSIA.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkImageFormatConstraintsFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatConstraintsFlagsFUCHSIA.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSysmemColorSpaceFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageFormatConstraintsInfoFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0