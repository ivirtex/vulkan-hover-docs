# VkPhysicalDeviceExtendedDynamicStateFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceExtendedDynamicStateFeaturesEXT - Structure describing
what extended dynamic state can be used



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceExtendedDynamicStateFeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_extended_dynamic_state
typedef struct VkPhysicalDeviceExtendedDynamicStateFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           extendedDynamicState;
} VkPhysicalDeviceExtendedDynamicStateFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-extendedDynamicState"></span>
  `extendedDynamicState` indicates that the implementation supports the
  following dynamic states:

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

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceExtendedDynamicStateFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceExtendedDynamicStateFeaturesEXT` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceExtendedDynamicStateFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceExtendedDynamicStateFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceExtendedDynamicStateFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_extended_dynamic_state](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceExtendedDynamicStateFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
