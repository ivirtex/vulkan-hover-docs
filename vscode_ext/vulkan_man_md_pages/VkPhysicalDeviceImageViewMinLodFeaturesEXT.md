# VkPhysicalDeviceImageViewMinLodFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceImageViewMinLodFeaturesEXT - Structure describing
whether clamping the min LOD of a image view is supported by the
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImageViewMinLodFeaturesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_image_view_min_lod
typedef struct VkPhysicalDeviceImageViewMinLodFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           minLod;
} VkPhysicalDeviceImageViewMinLodFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-minLod"></span> `minLod` indicates whether the
  implementation supports clamping the minimum LOD value during <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-image-level-selection"
  target="_blank" rel="noopener">Image Level(s) Selection</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-gather"
  target="_blank" rel="noopener">Texel Gathering</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-integer-coordinate-operations"
  target="_blank" rel="noopener">Integer Texel Coordinate Operations</a>
  with a given [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) by
  [VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewMinLodCreateInfoEXT.html)::`minLod`.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceImageViewMinLodFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceImageViewMinLodFeaturesEXT` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceImageViewMinLodFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceImageViewMinLodFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceImageViewMinLodFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_MIN_LOD_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_view_min_lod](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_view_min_lod.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImageViewMinLodFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
