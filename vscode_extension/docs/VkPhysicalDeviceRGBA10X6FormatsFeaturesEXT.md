# VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT - Structure describing whether rendering to VK\_FORMAT\_R10X6G10X6B10X6A10X6\_UNORM\_4PACK16 formats can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_rgba10x6_formats
typedef struct VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           formatRgba10x6WithoutYCbCrSampler;
} VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`formatRgba10x6WithoutYCbCrSampler` indicates that `VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16` **can** be used with a `VkImageView` with `subresourceRange.aspectMask` equal to `VK_IMAGE_ASPECT_COLOR_BIT` without a [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) enabled.

## [](#_description)Description

If the `VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RGBA10X6_FORMATS_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_rgba10x6\_formats](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_rgba10x6_formats.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0