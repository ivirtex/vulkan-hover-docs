# VkPhysicalDeviceShaderImageFootprintFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceShaderImageFootprintFeaturesNV - Structure describing
shader image footprint features that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderImageFootprintFeaturesNV` structure is
defined as:

``` c
// Provided by VK_NV_shader_image_footprint
typedef struct VkPhysicalDeviceShaderImageFootprintFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           imageFootprint;
} VkPhysicalDeviceShaderImageFootprintFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-imageFootprint"></span> `imageFootprint` specifies
  whether the implementation supports the `ImageFootprintNV` SPIR-V
  capability.

## <a href="#_description" class="anchor"></a>Description

See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-footprint"
target="_blank" rel="noopener">Texel Footprint Evaluation</a> for more
information.

If the `VkPhysicalDeviceShaderImageFootprintFeaturesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceShaderImageFootprintFeaturesNV` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceShaderImageFootprintFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceShaderImageFootprintFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderImageFootprintFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_IMAGE_FOOTPRINT_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_shader_image_footprint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shader_image_footprint.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderImageFootprintFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
