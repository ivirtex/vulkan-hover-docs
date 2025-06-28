# VkPhysicalDeviceExtendedDynamicStateFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceExtendedDynamicStateFeaturesEXT - Structure describing what extended dynamic state can be used



## [](#_c_specification)C Specification

The `VkPhysicalDeviceExtendedDynamicStateFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_extended_dynamic_state
typedef struct VkPhysicalDeviceExtendedDynamicStateFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           extendedDynamicState;
} VkPhysicalDeviceExtendedDynamicStateFeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`extendedDynamicState` indicates that the implementation supports the following dynamic states:
  
  - `VK_DYNAMIC_STATE_CULL_MODE`
  - `VK_DYNAMIC_STATE_FRONT_FACE`
  - `VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY`
  - `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT`
  - `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT`
  - `VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE`
  - `VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE`
  - `VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE`
  - `VK_DYNAMIC_STATE_DEPTH_COMPARE_OP`
  - `VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE`
  - `VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE`
  - `VK_DYNAMIC_STATE_STENCIL_OP`

## [](#_description)Description

If the `VkPhysicalDeviceExtendedDynamicStateFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceExtendedDynamicStateFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExtendedDynamicStateFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceExtendedDynamicStateFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExtendedDynamicStateFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0