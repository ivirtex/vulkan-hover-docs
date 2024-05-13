# VkPhysicalDeviceTextureCompressionASTCHDRFeatures(3) Manual Page

## Name

VkPhysicalDeviceTextureCompressionASTCHDRFeatures - Structure describing
ASTC HDR features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceTextureCompressionASTCHDRFeatures` structure is
defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceTextureCompressionASTCHDRFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           textureCompressionASTC_HDR;
} VkPhysicalDeviceTextureCompressionASTCHDRFeatures;
```

or the equivalent

``` c
// Provided by VK_EXT_texture_compression_astc_hdr
typedef VkPhysicalDeviceTextureCompressionASTCHDRFeatures VkPhysicalDeviceTextureCompressionASTCHDRFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-textureCompressionASTC_HDR"></span>
  `textureCompressionASTC_HDR` indicates whether all of the ASTC HDR
  compressed texture formats are supported. If this feature is enabled,
  then the `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`,
  `VK_FORMAT_FEATURE_BLIT_SRC_BIT` and
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT` features **must**
  be supported in `optimalTilingFeatures` for the following formats:

  - `VK_FORMAT_ASTC_4x4_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_5x4_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_5x5_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_6x5_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_6x6_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_8x5_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_8x6_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_8x8_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_10x5_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_10x6_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_10x8_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_10x10_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_12x10_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_12x12_SFLOAT_BLOCK`

  To query for additional properties, or if the feature is not enabled,
  [vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)
  and
  [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html)
  **can** be used to check for supported properties of individual
  formats as normal.

If the `VkPhysicalDeviceTextureCompressionASTCHDRFeatures` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceTextureCompressionASTCHDRFeatures` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceTextureCompressionASTCHDRFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceTextureCompressionASTCHDRFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceTextureCompressionASTCHDRFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXTURE_COMPRESSION_ASTC_HDR_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_texture_compression_astc_hdr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_texture_compression_astc_hdr.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceTextureCompressionASTCHDRFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
