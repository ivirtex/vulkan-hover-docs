# VkImageConstraintsInfoFUCHSIA(3) Manual Page

## Name

VkImageConstraintsInfoFUCHSIA - Structure of image-based buffer collection constraints



## [](#_c_specification)C Specification

The `VkImageConstraintsInfoFUCHSIA` structure is defined as:

```c++
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkImageConstraintsInfoFUCHSIA {
    VkStructureType                               sType;
    const void*                                   pNext;
    uint32_t                                      formatConstraintsCount;
    const VkImageFormatConstraintsInfoFUCHSIA*    pFormatConstraints;
    VkBufferCollectionConstraintsInfoFUCHSIA      bufferCollectionConstraints;
    VkImageConstraintsInfoFlagsFUCHSIA            flags;
} VkImageConstraintsInfoFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `formatConstraintsCount` is the number of elements in `pFormatConstraints`.
- `pFormatConstraints` is a pointer to an array of [VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatConstraintsInfoFUCHSIA.html) structures of size `formatConstraintsCount` that is used to further constrain buffer collection format selection for image-based buffer collections.
- `bufferCollectionConstraints` is a [VkBufferCollectionConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionConstraintsInfoFUCHSIA.html) structure used to supply parameters for the negotiation and allocation for buffer-based buffer collections.
- `flags` is a [VkImageConstraintsInfoFlagBitsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFlagBitsFUCHSIA.html) value specifying hints about the type of memory Sysmem should allocate for the buffer collection.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06395)VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06395  
  All elements of `pFormatConstraints` **must** have at least one bit set in its [VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatConstraintsInfoFUCHSIA.html)::`requiredFormatFeatures`
- [](#VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06396)VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06396  
  If `pFormatConstraints->imageCreateInfo->usage` contains `VK_IMAGE_USAGE_SAMPLED_BIT`, then `pFormatConstraints->requiredFormatFeatures` **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`
- [](#VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06397)VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06397  
  If `pFormatConstraints->imageCreateInfo->usage` contains `VK_IMAGE_USAGE_STORAGE_BIT`, then `pFormatConstraints->requiredFormatFeatures` **must** contain `VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT`
- [](#VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06398)VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06398  
  If `pFormatConstraints->imageCreateInfo->usage` contains `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`, then `pFormatConstraints->requiredFormatFeatures` **must** contain `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT`
- [](#VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06399)VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06399  
  If `pFormatConstraints->imageCreateInfo->usage` contains `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, then `pFormatConstraints->requiredFormatFeatures` **must** contain `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06400)VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-06400  
  If `pFormatConstraints->imageCreateInfo->usage` contains `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`, then `pFormatConstraints->requiredFormatFeatures` **must** contain at least one of `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` or `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageConstraintsInfoFUCHSIA-attachmentFragmentShadingRate-06401)VUID-VkImageConstraintsInfoFUCHSIA-attachmentFragmentShadingRate-06401  
  If the [`attachmentFragmentShadingRate`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-attachmentFragmentShadingRate) feature is enabled, and `pFormatConstraints->imageCreateInfo->usage` contains `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`, then `pFormatConstraints->requiredFormatFeatures` **must** contain `VK_FORMAT_FEATURE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-VkImageConstraintsInfoFUCHSIA-sType-sType)VUID-VkImageConstraintsInfoFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_CONSTRAINTS_INFO_FUCHSIA`
- [](#VUID-VkImageConstraintsInfoFUCHSIA-pNext-pNext)VUID-VkImageConstraintsInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-parameter)VUID-VkImageConstraintsInfoFUCHSIA-pFormatConstraints-parameter  
  `pFormatConstraints` **must** be a valid pointer to an array of `formatConstraintsCount` valid [VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatConstraintsInfoFUCHSIA.html) structures
- [](#VUID-VkImageConstraintsInfoFUCHSIA-bufferCollectionConstraints-parameter)VUID-VkImageConstraintsInfoFUCHSIA-bufferCollectionConstraints-parameter  
  `bufferCollectionConstraints` **must** be a valid [VkBufferCollectionConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionConstraintsInfoFUCHSIA.html) structure
- [](#VUID-VkImageConstraintsInfoFUCHSIA-flags-parameter)VUID-VkImageConstraintsInfoFUCHSIA-flags-parameter  
  `flags` **must** be a valid combination of [VkImageConstraintsInfoFlagBitsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFlagBitsFUCHSIA.html) values
- [](#VUID-VkImageConstraintsInfoFUCHSIA-formatConstraintsCount-arraylength)VUID-VkImageConstraintsInfoFUCHSIA-formatConstraintsCount-arraylength  
  `formatConstraintsCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkBufferCollectionConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionConstraintsInfoFUCHSIA.html), [VkImageConstraintsInfoFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFlagsFUCHSIA.html), [VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatConstraintsInfoFUCHSIA.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageConstraintsInfoFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0