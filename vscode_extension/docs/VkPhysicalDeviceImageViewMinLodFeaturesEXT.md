# VkPhysicalDeviceImageViewMinLodFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceImageViewMinLodFeaturesEXT - Structure describing whether clamping the min LOD of an image view is supported by the implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceImageViewMinLodFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_image_view_min_lod
typedef struct VkPhysicalDeviceImageViewMinLodFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           minLod;
} VkPhysicalDeviceImageViewMinLodFeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`minLod` indicates whether the implementation supports clamping the minimum LOD value during [Image Level(s) Selection](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-image-level-selection), [Texel Gathering](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-gather) and [Integer Texel Coordinate Operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-integer-coordinate-operations) with a given [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) by [VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewMinLodCreateInfoEXT.html)::`minLod`.

## [](#_description)Description

If the `VkPhysicalDeviceImageViewMinLodFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceImageViewMinLodFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceImageViewMinLodFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceImageViewMinLodFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_MIN_LOD_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_image\_view\_min\_lod](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_view_min_lod.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceImageViewMinLodFeaturesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0