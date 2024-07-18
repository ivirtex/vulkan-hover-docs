# VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT - Structure describing
advanced blending features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_blend_operation_advanced
typedef struct VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           advancedBlendCoherentOperations;
} VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-advancedBlendCoherentOperations"></span>
  `advancedBlendCoherentOperations` specifies whether blending using <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blend-advanced"
  target="_blank" rel="noopener">advanced blend operations</a> is
  guaranteed to execute atomically and in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-primitive-order"
  target="_blank" rel="noopener">primitive order</a>. If this is
  `VK_TRUE`, `VK_ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT` is
  treated the same as `VK_ACCESS_COLOR_ATTACHMENT_READ_BIT`, and
  advanced blending needs no additional synchronization over basic
  blending. If this is `VK_FALSE`, then memory dependencies are required
  to guarantee order between two advanced blending operations that occur
  on the same sample.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_blend_operation_advanced](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_blend_operation_advanced.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
