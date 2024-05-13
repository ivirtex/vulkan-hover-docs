# VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM(3) Manual Page

## Name

VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM - Structure
describing fragment density map offset features that can be supported by
an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM` structure is
defined as:

``` c
// Provided by VK_QCOM_fragment_density_map_offset
typedef struct VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           fragmentDensityMapOffset;
} VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-fragmentDensityMapOffsets"></span>
  `fragmentDensityMapOffsets` specifies whether the implementation
  supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-fragmentdensitymapoffsets"
  target="_blank" rel="noopener">fragment density map offsets</a>

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM-sType-sType"
  id="VUID-VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM-sType-sType"></a>
  VUID-VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_OFFSET_FEATURES_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_fragment_density_map_offset](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_fragment_density_map_offset.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
