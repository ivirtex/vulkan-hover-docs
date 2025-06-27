# VkPhysicalDeviceInheritedViewportScissorFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceInheritedViewportScissorFeaturesNV - Structure
describing the viewport scissor inheritance behavior for an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceInheritedViewportScissorFeaturesNV` structure is
defined as:

``` c
// Provided by VK_NV_inherited_viewport_scissor
typedef struct VkPhysicalDeviceInheritedViewportScissorFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           inheritedViewportScissor2D;
} VkPhysicalDeviceInheritedViewportScissorFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-inheritedViewportScissor2D"></span>
  `inheritedViewportScissor2D` indicates whether secondary command
  buffers can inherit most of the dynamic state affected by
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT`,
  `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT`,
  `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT`,
  `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT`,
  `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT`,
  `VK_DYNAMIC_STATE_VIEWPORT` or `VK_DYNAMIC_STATE_SCISSOR`, from a
  primary command buffer.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceInheritedViewportScissorFeaturesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceInheritedViewportScissorFeaturesNV` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceInheritedViewportScissorFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceInheritedViewportScissorFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceInheritedViewportScissorFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INHERITED_VIEWPORT_SCISSOR_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_inherited_viewport_scissor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_inherited_viewport_scissor.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceInheritedViewportScissorFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
