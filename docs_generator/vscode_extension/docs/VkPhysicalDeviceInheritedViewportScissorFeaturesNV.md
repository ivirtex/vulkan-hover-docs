# VkPhysicalDeviceInheritedViewportScissorFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceInheritedViewportScissorFeaturesNV - Structure describing the viewport scissor inheritance behavior for an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceInheritedViewportScissorFeaturesNV` structure is defined as:

```c++
// Provided by VK_NV_inherited_viewport_scissor
typedef struct VkPhysicalDeviceInheritedViewportScissorFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           inheritedViewportScissor2D;
} VkPhysicalDeviceInheritedViewportScissorFeaturesNV;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`inheritedViewportScissor2D` indicates whether secondary command buffers can inherit most of the dynamic state affected by `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT`, `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT`, `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT`, `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT`, `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT`, `VK_DYNAMIC_STATE_VIEWPORT` or `VK_DYNAMIC_STATE_SCISSOR`, from a primary command buffer.

## [](#_description)Description

If the `VkPhysicalDeviceInheritedViewportScissorFeaturesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceInheritedViewportScissorFeaturesNV`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceInheritedViewportScissorFeaturesNV-sType-sType)VUID-VkPhysicalDeviceInheritedViewportScissorFeaturesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INHERITED_VIEWPORT_SCISSOR_FEATURES_NV`

## [](#_see_also)See Also

[VK\_NV\_inherited\_viewport\_scissor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_inherited_viewport_scissor.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceInheritedViewportScissorFeaturesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0