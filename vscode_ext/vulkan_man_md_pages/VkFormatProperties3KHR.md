# VkFormatProperties3(3) Manual Page

## Name

VkFormatProperties3 - Structure specifying image format properties



## <a href="#_c_specification" class="anchor"></a>C Specification

To query supported format extended features which are properties of the
physical device, add [VkFormatProperties3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties3.html)
structure to the `pNext` chain of
[VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html).

The [VkFormatProperties3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties3.html) structure is defined
as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkFormatProperties3 {
    VkStructureType          sType;
    void*                    pNext;
    VkFormatFeatureFlags2    linearTilingFeatures;
    VkFormatFeatureFlags2    optimalTilingFeatures;
    VkFormatFeatureFlags2    bufferFeatures;
} VkFormatProperties3;
```

or the equivalent

``` c
// Provided by VK_KHR_format_feature_flags2
typedef VkFormatProperties3 VkFormatProperties3KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `linearTilingFeatures` is a bitmask of
  [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2.html) specifying
  features supported by images created with a `tiling` parameter of
  `VK_IMAGE_TILING_LINEAR`.

- `optimalTilingFeatures` is a bitmask of
  [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2.html) specifying
  features supported by images created with a `tiling` parameter of
  `VK_IMAGE_TILING_OPTIMAL`.

- `bufferFeatures` is a bitmask of
  [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2.html) specifying
  features supported by buffers.

## <a href="#_description" class="anchor"></a>Description

The bits reported in `linearTilingFeatures`, `optimalTilingFeatures` and
`bufferFeatures` **must** include the bits reported in the corresponding
fields of `VkFormatProperties2`::`formatProperties`.

Valid Usage (Implicit)

- <a href="#VUID-VkFormatProperties3-sType-sType"
  id="VUID-VkFormatProperties3-sType-sType"></a>
  VUID-VkFormatProperties3-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_3`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkFormatFeatureFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFormatProperties3"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
