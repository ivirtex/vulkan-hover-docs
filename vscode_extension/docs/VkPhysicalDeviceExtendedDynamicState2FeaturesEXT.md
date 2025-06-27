# VkPhysicalDeviceExtendedDynamicState2FeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceExtendedDynamicState2FeaturesEXT - Structure describing
what extended dynamic state can be used



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceExtendedDynamicState2FeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_extended_dynamic_state2
typedef struct VkPhysicalDeviceExtendedDynamicState2FeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           extendedDynamicState2;
    VkBool32           extendedDynamicState2LogicOp;
    VkBool32           extendedDynamicState2PatchControlPoints;
} VkPhysicalDeviceExtendedDynamicState2FeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-extendedDynamicState2"></span>
  `extendedDynamicState2` indicates that the implementation supports the
  following dynamic states:

  - `VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE`

  - `VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE`

  - `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE`

- <span id="features-extendedDynamicState2LogicOp"></span>
  `extendedDynamicState2LogicOp` indicates that the implementation
  supports the following dynamic state:

  - `VK_DYNAMIC_STATE_LOGIC_OP_EXT`

- <span id="features-extendedDynamicState2PatchControlPoints"></span>
  `extendedDynamicState2PatchControlPoints` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT`

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceExtendedDynamicState2FeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceExtendedDynamicState2FeaturesEXT` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceExtendedDynamicState2FeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceExtendedDynamicState2FeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceExtendedDynamicState2FeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_2_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_extended_dynamic_state2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state2.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceExtendedDynamicState2FeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
