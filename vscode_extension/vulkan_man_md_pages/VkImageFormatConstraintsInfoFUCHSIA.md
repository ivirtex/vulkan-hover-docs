# VkImageFormatConstraintsInfoFUCHSIA(3) Manual Page

## Name

VkImageFormatConstraintsInfoFUCHSIA - Structure image-based buffer
collection constraints



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageFormatConstraintsInfoFUCHSIA` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this structure

- `imageCreateInfo` is the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  used to create a [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) that is to use memory from
  the [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html)

- `requiredFormatFeatures` is a bitmask of `VkFormatFeatureFlagBits`
  specifying required features of the buffers in the buffer collection

- `flags` is reserved for future use

- `sysmemPixelFormat` is a `PixelFormatType` value from the
  `fuchsia.sysmem/image_formats.fidl` FIDL interface

- `colorSpaceCount` the element count of `pColorSpaces`

- `pColorSpaces` is a pointer to an array of
  [VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSysmemColorSpaceFUCHSIA.html) structs of
  size `colorSpaceCount`

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkImageFormatConstraintsInfoFUCHSIA-sType-sType"
  id="VUID-VkImageFormatConstraintsInfoFUCHSIA-sType-sType"></a>
  VUID-VkImageFormatConstraintsInfoFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_FORMAT_CONSTRAINTS_INFO_FUCHSIA`

- <a href="#VUID-VkImageFormatConstraintsInfoFUCHSIA-pNext-pNext"
  id="VUID-VkImageFormatConstraintsInfoFUCHSIA-pNext-pNext"></a>
  VUID-VkImageFormatConstraintsInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkImageFormatConstraintsInfoFUCHSIA-imageCreateInfo-parameter"
  id="VUID-VkImageFormatConstraintsInfoFUCHSIA-imageCreateInfo-parameter"></a>
  VUID-VkImageFormatConstraintsInfoFUCHSIA-imageCreateInfo-parameter  
  `imageCreateInfo` **must** be a valid
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure

- <a
  href="#VUID-VkImageFormatConstraintsInfoFUCHSIA-requiredFormatFeatures-parameter"
  id="VUID-VkImageFormatConstraintsInfoFUCHSIA-requiredFormatFeatures-parameter"></a>
  VUID-VkImageFormatConstraintsInfoFUCHSIA-requiredFormatFeatures-parameter  
  `requiredFormatFeatures` **must** be a valid combination of
  [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html) values

- <a
  href="#VUID-VkImageFormatConstraintsInfoFUCHSIA-requiredFormatFeatures-requiredbitmask"
  id="VUID-VkImageFormatConstraintsInfoFUCHSIA-requiredFormatFeatures-requiredbitmask"></a>
  VUID-VkImageFormatConstraintsInfoFUCHSIA-requiredFormatFeatures-requiredbitmask  
  `requiredFormatFeatures` **must** not be `0`

- <a href="#VUID-VkImageFormatConstraintsInfoFUCHSIA-flags-zerobitmask"
  id="VUID-VkImageFormatConstraintsInfoFUCHSIA-flags-zerobitmask"></a>
  VUID-VkImageFormatConstraintsInfoFUCHSIA-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkImageFormatConstraintsInfoFUCHSIA-pColorSpaces-parameter"
  id="VUID-VkImageFormatConstraintsInfoFUCHSIA-pColorSpaces-parameter"></a>
  VUID-VkImageFormatConstraintsInfoFUCHSIA-pColorSpaces-parameter  
  `pColorSpaces` **must** be a valid pointer to an array of
  `colorSpaceCount` valid
  [VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSysmemColorSpaceFUCHSIA.html) structures

- <a
  href="#VUID-VkImageFormatConstraintsInfoFUCHSIA-colorSpaceCount-arraylength"
  id="VUID-VkImageFormatConstraintsInfoFUCHSIA-colorSpaceCount-arraylength"></a>
  VUID-VkImageFormatConstraintsInfoFUCHSIA-colorSpaceCount-arraylength  
  `colorSpaceCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags.html),
[VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFUCHSIA.html),
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
[VkImageFormatConstraintsFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatConstraintsFlagsFUCHSIA.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSysmemColorSpaceFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageFormatConstraintsInfoFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
