# VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT - Structure describing
whether rendering to VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16
formats can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_rgba10x6_formats
typedef struct VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           formatRgba10x6WithoutYCbCrSampler;
} VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-formatRgba10x6WithoutYCbCrSampler"></span>
  `formatRgba10x6WithoutYCbCrSampler` indicates that
  `VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16` **can** be used with a
  `VkImageView` with `subresourceRange.aspectMask` equal to
  `VK_IMAGE_ASPECT_COLOR_BIT` without a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
  target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion</a> enabled.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RGBA10X6_FORMATS_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_rgba10x6_formats](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_rgba10x6_formats.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
