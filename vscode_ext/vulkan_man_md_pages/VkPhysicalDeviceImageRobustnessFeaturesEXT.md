# VkPhysicalDeviceImageRobustnessFeatures(3) Manual Page

## Name

VkPhysicalDeviceImageRobustnessFeatures - Structure describing the
out-of-bounds behavior for an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImageRobustnessFeatures` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceImageRobustnessFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           robustImageAccess;
} VkPhysicalDeviceImageRobustnessFeatures;
```

or the equivalent

``` c
// Provided by VK_EXT_image_robustness
typedef VkPhysicalDeviceImageRobustnessFeatures VkPhysicalDeviceImageRobustnessFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-robustImageAccess"></span>
  `robustImageAccess` indicates whether image accesses are tightly
  bounds-checked against the dimensions of the image view. <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-input-validation"
  target="_blank" rel="noopener">Invalid texels</a> resulting from out
  of bounds image loads will be replaced as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-texel-replacement"
  target="_blank" rel="noopener">Texel Replacement</a>, with either
  (0,0,1) or (0,0,0) values inserted for missing G, B, or A components
  based on the format.

If the `VkPhysicalDeviceImageRobustnessFeatures` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceImageRobustnessFeatures` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceImageRobustnessFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceImageRobustnessFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceImageRobustnessFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ROBUSTNESS_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_robustness](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_robustness.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImageRobustnessFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
