# VkBufferCollectionPropertiesFUCHSIA(3) Manual Page

## Name

VkBufferCollectionPropertiesFUCHSIA - Structure specifying the negotiated format chosen by Sysmem



## [](#_c_specification)C Specification

The `VkBufferCollectionPropertiesFUCHSIA` structure is defined as:

```c++
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkBufferCollectionPropertiesFUCHSIA {
    VkStructureType                  sType;
    void*                            pNext;
    uint32_t                         memoryTypeBits;
    uint32_t                         bufferCount;
    uint32_t                         createInfoIndex;
    uint64_t                         sysmemPixelFormat;
    VkFormatFeatureFlags             formatFeatures;
    VkSysmemColorSpaceFUCHSIA        sysmemColorSpaceIndex;
    VkComponentMapping               samplerYcbcrConversionComponents;
    VkSamplerYcbcrModelConversion    suggestedYcbcrModel;
    VkSamplerYcbcrRange              suggestedYcbcrRange;
    VkChromaLocation                 suggestedXChromaOffset;
    VkChromaLocation                 suggestedYChromaOffset;
} VkBufferCollectionPropertiesFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure
- `memoryTypeBits` is a bitmask containing one bit set for every memory type which the buffer collection can be imported as buffer collection
- `bufferCount` is the number of buffers in the collection
- `createInfoIndex` as described in [Sysmem chosen create infos](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#sysmem-chosen-create-infos)
- `sysmemPixelFormat` is the Sysmem `PixelFormatType` as defined in `fuchsia.sysmem/image_formats.fidl`
- `formatFeatures` is a bitmask of [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html) shared by the buffer collection
- `sysmemColorSpaceIndex` is a [VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSysmemColorSpaceFUCHSIA.html) struct specifying the color space
- `samplerYcbcrConversionComponents` is a [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html) structure specifying the component mapping
- `suggestedYcbcrModel` is a [VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrModelConversion.html) value specifying the suggested Y′CBCR model
- `suggestedYcbcrRange` is a [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrRange.html) value specifying the suggested Y′CBCR range
- `suggestedXChromaOffset` is a [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html) value specifying the suggested X chroma offset
- `suggestedYChromaOffset` is a [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html) value specifying the suggested Y chroma offset

## [](#_description)Description

`sysmemColorSpace` is only set for image-based buffer collections where the constraints were specified using [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFUCHSIA.html) in a call to [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html).

For image-based buffer collections, `createInfoIndex` will identify both the [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFUCHSIA.html)::`pImageCreateInfos` element and the [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFUCHSIA.html)::`pFormatConstraints` element chosen by Sysmem when [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html) was called. The value of `sysmemColorSpaceIndex` will be an index to one of the color spaces provided in the [VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatConstraintsInfoFUCHSIA.html)::`pColorSpaces` array.

The implementation **must** have `formatFeatures` with all bits set that were set in [VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatConstraintsInfoFUCHSIA.html)::`requiredFormatFeatures`, by the call to [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html), at `createInfoIndex` (other bits could be set as well).

Valid Usage (Implicit)

- [](#VUID-VkBufferCollectionPropertiesFUCHSIA-sType-sType)VUID-VkBufferCollectionPropertiesFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_PROPERTIES_FUCHSIA`
- [](#VUID-VkBufferCollectionPropertiesFUCHSIA-pNext-pNext)VUID-VkBufferCollectionPropertiesFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html), [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html), [VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags.html), [VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrModelConversion.html), [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrRange.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSysmemColorSpaceFUCHSIA.html), [vkGetBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferCollectionPropertiesFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferCollectionPropertiesFUCHSIA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0