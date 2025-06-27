# VkBufferCollectionPropertiesFUCHSIA(3) Manual Page

## Name

VkBufferCollectionPropertiesFUCHSIA - Structure specifying the
negotiated format chosen by Sysmem



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferCollectionPropertiesFUCHSIA` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this structure

- `memoryTypeBits` is a bitmask containing one bit set for every memory
  type which the buffer collection can be imported as buffer collection

- `bufferCount` is the number of buffers in the collection

- `createInfoIndex` as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#sysmem-chosen-create-infos"
  target="_blank" rel="noopener">Sysmem chosen create infos</a>

- `sysmemPixelFormat` is the Sysmem `PixelFormatType` as defined in
  `fuchsia.sysmem/image_formats.fidl`

- `formatFeatures` is a bitmask of
  [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html) shared by the
  buffer collection

- `sysmemColorSpaceIndex` is a
  [VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSysmemColorSpaceFUCHSIA.html) struct
  specifying the color space

- `samplerYcbcrConversionComponents` is a
  [VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html) struct specifying the
  component mapping

- `suggestedYcbcrModel` is a
  [VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html)
  value specifying the suggested Y′C<sub>B</sub>C<sub>R</sub> model

- `suggestedYcbcrRange` is a
  [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html) value specifying the
  suggested Y′C<sub>B</sub>C<sub>R</sub> range

- `suggestedXChromaOffset` is a
  [VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html) value specifying the
  suggested X chroma offset

- `suggestedYChromaOffset` is a
  [VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html) value specifying the
  suggested Y chroma offset

## <a href="#_description" class="anchor"></a>Description

`sysmemColorSpace` is only set for image-based buffer collections where
the constraints were specified using
[VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFUCHSIA.html) in a
call to
[vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html).

For image-based buffer collections, `createInfoIndex` will identify both
the
[VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFUCHSIA.html)::`pImageCreateInfos`
element and the
[VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFUCHSIA.html)::`pFormatConstraints`
element chosen by Sysmem when
[vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html)
was called. The value of `sysmemColorSpaceIndex` will be an index to one
of the color spaces provided in the
[VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatConstraintsInfoFUCHSIA.html)::`pColorSpaces`
array.

The implementation must have `formatFeatures` with all bits set that
were set in
[VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatConstraintsInfoFUCHSIA.html)::`requiredFormatFeatures`,
by the call to
[vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html),
at `createInfoIndex` (other bits could be set as well).

Valid Usage (Implicit)

- <a href="#VUID-VkBufferCollectionPropertiesFUCHSIA-sType-sType"
  id="VUID-VkBufferCollectionPropertiesFUCHSIA-sType-sType"></a>
  VUID-VkBufferCollectionPropertiesFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_PROPERTIES_FUCHSIA`

- <a href="#VUID-VkBufferCollectionPropertiesFUCHSIA-pNext-pNext"
  id="VUID-VkBufferCollectionPropertiesFUCHSIA-pNext-pNext"></a>
  VUID-VkBufferCollectionPropertiesFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html),
[VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html),
[VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags.html),
[VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html),
[VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSysmemColorSpaceFUCHSIA.html),
[vkGetBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferCollectionPropertiesFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferCollectionPropertiesFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
