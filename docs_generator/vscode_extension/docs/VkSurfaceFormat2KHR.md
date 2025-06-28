# VkSurfaceFormat2KHR(3) Manual Page

## Name

VkSurfaceFormat2KHR - Structure describing a supported swapchain format tuple



## [](#_c_specification)C Specification

The `VkSurfaceFormat2KHR` structure is defined as:

```c++
// Provided by VK_KHR_get_surface_capabilities2
typedef struct VkSurfaceFormat2KHR {
    VkStructureType       sType;
    void*                 pNext;
    VkSurfaceFormatKHR    surfaceFormat;
} VkSurfaceFormat2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `surfaceFormat` is a [VkSurfaceFormatKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFormatKHR.html) structure describing a format-color space pair that is compatible with the specified surface.

## [](#_description)Description

If the [`imageCompressionControlSwapchain`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-imageCompressionControlSwapchain) feature is supported and a [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionPropertiesEXT.html) structure is included in the `pNext` chain of this structure, then it will be filled with the compression properties that are supported for the `surfaceFormat`.

Valid Usage

- [](#VUID-VkSurfaceFormat2KHR-pNext-06750)VUID-VkSurfaceFormat2KHR-pNext-06750  
  If the `VK_EXT_image_compression_control_swapchain` extension is not supported, the `pNext` chain **must** not include an [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionPropertiesEXT.html) structure

Valid Usage (Implicit)

- [](#VUID-VkSurfaceFormat2KHR-sType-sType)VUID-VkSurfaceFormat2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR`
- [](#VUID-VkSurfaceFormat2KHR-pNext-pNext)VUID-VkSurfaceFormat2KHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionPropertiesEXT.html)
- [](#VUID-VkSurfaceFormat2KHR-sType-unique)VUID-VkSurfaceFormat2KHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSurfaceFormatKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFormatKHR.html), [vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfaceFormat2KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0