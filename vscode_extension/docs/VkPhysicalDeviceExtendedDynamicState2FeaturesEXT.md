# VkPhysicalDeviceExtendedDynamicState2FeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceExtendedDynamicState2FeaturesEXT - Structure describing what extended dynamic state can be used



## [](#_c_specification)C Specification

The `VkPhysicalDeviceExtendedDynamicState2FeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_extended_dynamic_state2
typedef struct VkPhysicalDeviceExtendedDynamicState2FeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           extendedDynamicState2;
    VkBool32           extendedDynamicState2LogicOp;
    VkBool32           extendedDynamicState2PatchControlPoints;
} VkPhysicalDeviceExtendedDynamicState2FeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`extendedDynamicState2` indicates that the implementation supports the following dynamic states:
  
  - `VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE`
  - `VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE`
  - `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE`
- []()`extendedDynamicState2LogicOp` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_LOGIC_OP_EXT`
- []()`extendedDynamicState2PatchControlPoints` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT`

## [](#_description)Description

If the `VkPhysicalDeviceExtendedDynamicState2FeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceExtendedDynamicState2FeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExtendedDynamicState2FeaturesEXT-sType-sType)VUID-VkPhysicalDeviceExtendedDynamicState2FeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_2_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExtendedDynamicState2FeaturesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0